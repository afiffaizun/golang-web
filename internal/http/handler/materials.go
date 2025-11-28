// internal/http/handler/materials.go
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/afiffaizun/golang-web/internal/storage/memory"
)

func ListMaterials(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(memory.GetAllMaterials())
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