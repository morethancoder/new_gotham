package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func onError(w http.ResponseWriter, err error, msg string, code int) {
	if err != nil {
		http.Error(w, msg, code)
		log.Println(msg, err)
	}
}

func RenderView(w http.ResponseWriter, r *http.Request, view templ.Component, layoutPath string) {
	if r.Header.Get("Hx-Request") == "true" {
		err := view.Render(r.Context(), w)
		onError(w, err, "Internal server error", http.StatusInternalServerError)
		return
	}
	err := Layout(layoutPath).Render(r.Context(), w)
	onError(w, err, "Internal server error", http.StatusInternalServerError)
}

func HomeGetHandler(w http.ResponseWriter, r *http.Request) {
    title := "Hello World!"
    msg := `start by editing this text, find it in ./handlers.go as var called "msg".`

    dbClient, ok := r.Context().Value(DbClientKey).(*sql.DB)
    if !ok {
        onError(w, errors.New("Couldnt retrieve dbclient from context"),
        "Internal server error", http.StatusInternalServerError)
        return
    }
    globalValues := GlobalValuesInstance{ID: 1}
    err := globalValues.Create(dbClient)
    onError(w, err, "Internal server error", http.StatusInternalServerError)

    err = globalValues.Read(dbClient)
    onError(w, err, "Internal server error", http.StatusInternalServerError)

    RenderView(w, r, Home(title, msg, globalValues.Count), "/")
} 


func CountPostHandler(w http.ResponseWriter, r *http.Request) {
    dbClient, ok := r.Context().Value(DbClientKey).(*sql.DB)
    if !ok {
        onError(w, errors.New("Couldnt retrieve dbclient from context"),
        "Internal server error", http.StatusInternalServerError)
        return
    }
    
    globalValues := GlobalValuesInstance{ID: 1}
    err := globalValues.Read(dbClient)
    onError(w, err, "Internal server error", http.StatusInternalServerError)

    err = globalValues.Update(dbClient)
    onError(w, err, "Internal server error", http.StatusInternalServerError)

    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprint(w, globalValues.Count)
}
