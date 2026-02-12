package handler

import (
	"html/template"
	"net/http"
)

type HomeHandler struct {
	tmpl *template.Template
}

func NewHomeHandler() *HomeHandler {
	tmpl := template.Must(
		template.ParseFiles("web/templates/index.html"),
	)

	return &HomeHandler{
		tmpl: tmpl,
	}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	h.tmpl.Execute(w, nil)
}
