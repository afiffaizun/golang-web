package handler

import (
	"html/template"
	"net/http"
	"strings"
	"sync"
)

var (
	tplOnce sync.Once
	tpl     *template.Template
)

func loadTemplates() {
	tpl = template.Must(template.ParseFiles(
		"web/templates/layouts.html",
		"web/templates/index.html",
		"web/templates/material_detail.html",
	))
}

func renderTemplate(w http.ResponseWriter, name string, data any) error {
	tplOnce.Do(loadTemplates)
	return tpl.ExecuteTemplate(w, name, data)
}

func prefersHTML(r *http.Request) bool {
	if r.URL.Path == "/" {
		return true
	}

	accept := r.Header.Get("Accept")
	if strings.Contains(accept, "text/html") {
		return true
	}
	// Browsers often send */*; treat as HTML when coming from GET without explicit JSON preference.
	return accept == "" && r.Method == http.MethodGet
}

func isFormRequest(r *http.Request) bool {
	ct := r.Header.Get("Content-Type")
	return strings.Contains(ct, "application/x-www-form-urlencoded") ||
		strings.Contains(ct, "multipart/form-data")
}
