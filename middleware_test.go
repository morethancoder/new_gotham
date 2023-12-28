package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/go-chi/chi/v5"
)

func TestDatabaseMiddleware(t *testing.T) {
    dbClient := &sql.DB{}
    testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        dbClient, ok := r.Context().Value(DbClientKey).(*sql.DB)
        if !ok {
            t.Fatal("Couldnt find dbClient from Context!")
        }
        _ = dbClient

        w.WriteHeader(http.StatusOK)
    })
   
    r := chi.NewRouter()
    r.Use(DbClientMiddleware(dbClient))
    r.Get("/test", testHandler)

    req, err := http.NewRequest("GET","/test", nil)
    if err != nil {t.Fatal(err)}

    recorder := httptest.NewRecorder()
    r.ServeHTTP(recorder, req)

    if status := recorder.Code; status != http.StatusOK {
        t.Errorf("Test Handler returned wrong status code: got:%d want:%d", recorder.Code, http.StatusOK )
    }
}

