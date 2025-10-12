/*
Пример успешного запроса:

curl -u john:secret http://localhost:8080/answer/
# Welcome, john!
Пример неуспешного запроса:

curl http://localhost:8080/answer/
# HTTP/1.1 401 Unauthorized
# WWW-Authenticate: Basic realm="Restricted"
*/
package step5httpmdlwrs


import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/answer/", Authorization(answerHandler))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	userval := r.Context().Value("username")
	if userval == nil {
		setUnauthorized(w)
		return
	}
	user, ok := userval.(string)
	if !ok {
		setUnauthorized(w)
		return
	}
	w.Write([]byte("Welcome, " + user + "!"))
}	

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Println(authHeader)
		if authHeader == "" {
			setUnauthorized(w)
			return
		}

		if !strings.HasPrefix(authHeader, "Basic ") {
			setUnauthorized(w)
			return
		}

		encodeInfo := strings.TrimPrefix(authHeader, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(encodeInfo)
		if err != nil {
			setUnauthorized(w)
			return
		}


		parts := strings.SplitN(string(decoded), ":", 2)
		if len(parts) != 2 {
			setUnauthorized(w)
			return
		}

		username := parts[0]
		password := parts[1]

		if username == "" || password == "" {
			setUnauthorized(w)
			return
		}

		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func setUnauthorized(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}