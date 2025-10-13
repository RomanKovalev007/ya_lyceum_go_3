package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

type my_WR struct{
	http.ResponseWriter
	Status int
}

func NewWR (w http.ResponseWriter) *my_WR{
	return &my_WR{
		ResponseWriter: w,
		Status: 200}
}

func (wr *my_WR) WriteHeader(statuscode int){
	wr.Status = statuscode
	wr.ResponseWriter.WriteHeader(statuscode)
}

func (wr *my_WR) Write(b []byte) (int, error){
	if wr.Status == 200 {
        wr.ResponseWriter.WriteHeader(200)
    }
	return wr.ResponseWriter.Write(b)
}

func (wr *my_WR) Header() http.Header{
	return wr.ResponseWriter.Header()
}



type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

type Store struct{
	s map[int]User
	nextID int64
}

func NewStore() *Store{
	return &Store{
		s: make(map[int]User),
		nextID: 1,
	}
}


func(s *Store) CreateUser(name string, age int) User {
	id := atomic.AddInt64(&s.nextID, 1) - 1
	createUser := User{ID: int(id), Name: name, Age: age}
	s.s[int(id)] = createUser
	return createUser
}

func(s *Store) GetUser(id int) (User, bool) {
	user, ok := s.s[id]
	return user, ok
}

func CreateUserHandler(s *Store) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var newUser User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil{
			w.WriteHeader(500)
			return
		}
		createdUser := s.CreateUser(newUser.Name, newUser.Age)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(createdUser)
		
	}
}

func GetUserHandler(s *Store) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil{
			w.WriteHeader(500)
			return 
		}
		findUser, ok := s.GetUser(id)
		if !ok{
			w.WriteHeader(404)
			return 
		}

		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(findUser)
		if err != nil{
			return 
		}
	}
}

func loggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler{
	return func(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			wr :=  NewWR(w)

			next.ServeHTTP(wr, r)

			logger.Info("",
				"method", r.Method,
				"path", r.URL.Path,
				"status", strconv.Itoa(wr.Status),
				"time", time.Since(now),
			)
		})
	}
}




func main() {
	mux := http.NewServeMux()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	
	store := NewStore()
	mux.HandleFunc("POST /users", CreateUserHandler(store))
	mux.HandleFunc("GET /users/{id}", GetUserHandler(store))

	mdlwrlog := loggingMiddleware(logger)(mux)

	if err := http.ListenAndServe("localhost:8080", mdlwrlog); err != nil {
		panic(err)
	}
}

