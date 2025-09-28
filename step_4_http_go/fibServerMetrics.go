package step4httpgo

import (
	"encoding/json"
	"net/http"
)

func main3(){
	n1 := 0
	n2 := 1

	count := 0

	mux := http.NewServeMux()
	http.HandleFunc("/", 
		func(w http.ResponseWriter, r *http.Request) {
			count += 1
			json.NewEncoder(w).Encode(n1)
			n3 := n1
			n1 = n2
			n2 += n3
			
		}	)

	http.HandleFunc("/metrics", 
		func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(count)
		}	)
		
	if err := http.ListenAndServe("", mux); err != nil{
		return
	}
}