package step5httpmdlwrs

import (
	"log/slog"
	"net/http"
)

func Logger(next http.Handler) http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			slog.Info("incoming request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
			next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, middleware!"))
}

func main1(){
	mux := http.NewServeMux()
	http.Handle("/hello", Logger(http.HandlerFunc(helloHandler)))

	if err := http.ListenAndServe("", mux); err != nil{
		return
	}
}