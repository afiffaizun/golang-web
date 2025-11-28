// internal/http/handler/materials.go
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/afiffaizun/golang-web/internal/material"
	"github.com/afiffaizun/golang-web/internal/storage/memory"
)

type materialsPageData struct {
	Materials []material.Material
}

func ListMaterials(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/materials" && r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	materials := memory.GetAllMaterials()
	if prefersHTML(r) || r.URL.Path == "/" {
		if err := renderTemplate(w, "index.html", materialsPageData{Materials: materials}); err != nil {
			http.Error(w, "failed to render page", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(materials)
}

func MaterialsEntryPoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ListMaterials(w, r)
	case http.MethodPost:
		CreateMaterial(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
