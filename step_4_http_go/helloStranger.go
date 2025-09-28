package step4httpgo

import (
	"encoding/json"
	"net/http"
	"regexp"
)

func main1(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == ""{
			name = "stranger"
		} else if containEnglish(name){
			name = "dirty hacker"
		}
		
		msg := "hello " + name
		json.NewEncoder(w).Encode(msg)

	})
	if err := http.ListenAndServe("", mux); err != nil{
		return
	}
}

func containEnglish(name string) bool{
	match, err := regexp.MatchString("^[a-zA-z]+$", name)
	if err != nil{
		return false
	}
	return match
}