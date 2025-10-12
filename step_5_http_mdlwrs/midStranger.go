package step5httpmdlwrs

import (
	"encoding/json"
	"net/http"
	"regexp"
)

func Sanitize(next http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			name := r.URL.Query().Get("name")
			if !containEnglish(name){
				query := r.URL.Query()
				query.Set("name", "dirty hacker")
				r.URL.RawQuery = query.Encode()
			}
			next.ServeHTTP(w, r)
	})
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc{
		return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			name := r.URL.Query().Get("name")
			if name == ""{
				query := r.URL.Query()
				query.Set("name", "stranger")
				r.URL.RawQuery = query.Encode()
			}
			next.ServeHTTP(w, r)
	})
}

func helloHandler1(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	msg := "hello " + name
	json.NewEncoder(w).Encode(msg)
}

func containEnglish(name string) bool{
	match, err := regexp.MatchString("^[a-zA-z]+$", name)
	if err != nil{
		return false
	}
	return match
}

func main2(){
	mux := http.NewServeMux()
	mux.Handle("/hello", SetDefaultName(Sanitize(http.HandlerFunc(helloHandler1))))
	if err := http.ListenAndServe("", mux); err != nil{
		return
	}
}

