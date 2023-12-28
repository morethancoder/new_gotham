package main

import (
	"context"
	"database/sql"
	"net/http"
)

type key int

const (
    DbClientKey key = iota
)

func DbClientMiddleware(dbClient *sql.DB) func(http.Handler) http.Handler {
    return func(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx := context.WithValue(r.Context(), DbClientKey, dbClient)
            h.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

