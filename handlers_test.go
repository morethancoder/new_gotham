package main 

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRenderViewHandlers(t *testing.T) {
    routePath := "/"
    view := Home("Hello", "message", 10)
    reqHx, err := http.NewRequest("GET", routePath, nil)
    if err != nil {t.Fatal(err)}
    reqHx.Header.Set("Hx-Request", "true")

    req, err := http.NewRequest("GET", routePath, nil)
    if err != nil {t.Fatal(err)}

    rrHx := httptest.NewRecorder()
    rr := httptest.NewRecorder()

    RenderView(rrHx, reqHx, view, routePath)
    if status := rrHx.Code; status != http.StatusOK {
        t.Errorf("Render with hx returned wrong status: got %v want %v", status, http.StatusOK)
    }

    RenderView(rr, req, view, routePath)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Render without hx returned wrong status: got %v want %v", status, http.StatusOK)
    }
}



