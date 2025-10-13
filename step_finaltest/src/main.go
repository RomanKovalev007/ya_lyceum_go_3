package main

import (
	"final_project/include/config"
	"final_project/include/handlers"
	"final_project/include/logger"
	"final_project/include/middleware"
	"final_project/include/storage"
	"net/http"
)


func main() {
	mux := http.NewServeMux()
	logger := logger.New()
	store := storage.New()

	mux.HandleFunc("POST /users", handlers.CreateUserHandler(store))
	mux.HandleFunc("GET /users/{id}", handlers.GetUserHandler(store))

	mdlwrlog := middleware.LoggingMiddleware(logger)(mux)

	cfg := config.SetupConfig()
	if err := http.ListenAndServe("localhost"+cfg.Port, mdlwrlog); err != nil {
		panic(err)
	}
}
