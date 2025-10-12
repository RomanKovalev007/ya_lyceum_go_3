package step5httpmdlwrs

import (
	"encoding/json"
	"net/http"
)

var count = 0

func middlewareMetrics(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count ++
		next.ServeHTTP(w, r)
	}
}

func main3(){
	n1 := 0
	n2 := 1


	mux := http.NewServeMux()
	mux.HandleFunc("/", middlewareMetrics(
		func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(n1)
			n3 := n1
			n1 = n2
			n2 += n3
		}))

	http.HandleFunc("/metrics", 
		func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(count)
		})

	if err := http.ListenAndServe("", mux); err != nil {
		return
	}
}
			


