package middleware

import (
	"log/slog"
	"net/http"
	"strconv"
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


func LoggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler{
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