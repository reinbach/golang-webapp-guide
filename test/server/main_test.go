package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"
)

func TestHome(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	c := web.C{}
	Home(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("200 expected, got %v instead", w.Code)
	}
}

func TestNotFound(t *testing.T) {
	r, _ := http.NewRequest("GET", "/404", nil)
	w := httptest.NewRecorder()
	c := web.C{}
	NotFound(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("200 expected, got %v instead", w.Code)
	}
}
