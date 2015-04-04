package template

import (
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/zenazn/goji/web"
)

func TestUpdateTemplateList(t *testing.T) {
	l := UpdateTemplateList([]string{"test.html"})
	if len(l) != 1 {
		t.Errorf("Expected list of 1 templates, got %v", len(l))
	}
}

func TestStaticHandlerValid(t *testing.T) {
	p := path.Join(STATIC_URL, "css/main.css")
	r, _ := http.NewRequest("GET", p, nil)
	w := httptest.NewRecorder()
	StaticHandler(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("200 expected, got %v instead", w.Code)
	}
}

func TestStaticHandlerInValid(t *testing.T) {
	p := path.Join(STATIC_URL, "css/something.css")
	r, _ := http.NewRequest("GET", p, nil)
	w := httptest.NewRecorder()
	StaticHandler(w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("404 expected, got %v instead", w.Code)
	}
}

func TestRender(t *testing.T) {
	c := web.C{}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	Render(c, w, r, []string{"home.html"}, &Context{})
}
