package step4httpgo

import (
	"encoding/json"
	"net/http"
)

func main4(){
	mux := http.NewServeMux()
	http.HandleFunc("/echo", 
		func(w http.ResponseWriter, r *http.Request) {
			echo := r.URL.Query().Get("echo")
			if echo == ""{
				echo = "empty"
			}
			json.NewEncoder(w).Encode(echo)
		}	)
	if err := http.ListenAndServe("", mux); err != nil{
		return
	}
}