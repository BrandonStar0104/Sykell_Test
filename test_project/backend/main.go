package main

import (
	"log"
	"net/http"

	"sykell/test/handlers"
	"sykell/test/middlewares"
)

func main() {
	http.HandleFunc("/process", middlewares.CORSMiddleware(middlewares.AuthMiddleware(handlers.ProcessURL)))
	http.HandleFunc("/results", middlewares.CORSMiddleware(middlewares.AuthMiddleware(handlers.GetResults)))
	http.HandleFunc("/result/", middlewares.CORSMiddleware(middlewares.AuthMiddleware(handlers.GetResult)))
	http.HandleFunc("/stop", middlewares.CORSMiddleware(middlewares.AuthMiddleware(handlers.StopProcessing)))

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
