package handlers

import (
	"encoding/json"
	"final_project/include/models"
	"final_project/include/storage"
	"net/http"
	"strconv"
)

func CreateUserHandler(s *storage.Store) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var newUser models.User
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

func GetUserHandler(s *storage.Store) http.HandlerFunc{
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