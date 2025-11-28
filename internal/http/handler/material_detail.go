package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/afiffaizun/golang-web/internal/material"
	"github.com/afiffaizun/golang-web/internal/note"
	"github.com/afiffaizun/golang-web/internal/storage/memory"
)

type materialDetailResponse struct {
	Material material.Material `json:"material"`
	Notes    []note.Note       `json:"notes"`
}

func GetMaterialDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := extractMaterialID(r.URL.Path)
	if err != nil {
		http.Error(w, "invalid material id", http.StatusBadRequest)
		return
	}

	material, ok := memory.GetMaterialByID(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	notes := memory.GetNotesByMaterialID(id)
	if prefersHTML(r) {
		data := materialDetailResponse{
			Material: material,
			Notes:    notes,
		}
		if err := renderTemplate(w, "material_detail.html", data); err != nil {
			http.Error(w, "failed to render page", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(materialDetailResponse{
		Material: material,
		Notes:    notes,
	})
}

func extractMaterialID(path string) (int, error) {
	trimmed := strings.TrimPrefix(path, "/materials/")
	if idx := strings.Index(trimmed, "/"); idx >= 0 {
		trimmed = trimmed[:idx]
	}
	return strconv.Atoi(trimmed)
}
