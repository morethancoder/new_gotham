package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
)


func main() {
    dbClient, err := InitDatabase(".env")
    if err != nil {log.Fatal(err)}
    err = InitGlobalValuesTable(dbClient)
    if err != nil {log.Fatal(err)}
    r := chi.NewRouter()
    r.Use(DbClientMiddleware(dbClient))
    fs := http.FileServer(http.Dir("static"))
    r.Handle("/static/*", http.StripPrefix("/static/", fs))
    r.Get("/", HomeGetHandler)
    r.Post("/count", CountPostHandler)
    log.Println("running on >> http://127.0.0.1:8000")
    err = http.ListenAndServe(":8000", r)
    if err != nil {
        log.Fatal(err)
    } 
}
