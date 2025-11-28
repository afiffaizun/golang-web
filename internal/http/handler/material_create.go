package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/afiffaizun/golang-web/internal/material"
	"github.com/afiffaizun/golang-web/internal/storage/memory"
)

type createMaterialRequest struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

func CreateMaterial(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req createMaterialRequest
	if isFormRequest(r) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "invalid form data", http.StatusBadRequest)
			return
		}
		req.Title = r.FormValue("title")
		req.Summary = r.FormValue("summary")
	} else {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}
	}

	req.Title = strings.TrimSpace(req.Title)
	if len(req.Title) < 3 {
		http.Error(w, "title must be at least 3 characters", http.StatusBadRequest)
		return
	}

	newMat := material.Material{
		Title:   req.Title,
		Summary: strings.TrimSpace(req.Summary),
	}

	created := memory.AddMaterial(newMat)

	if isFormRequest(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(created)
}
