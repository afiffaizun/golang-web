package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/afiffaizun/golang-web/internal/note"
	"github.com/afiffaizun/golang-web/internal/storage/memory"
)

type createNoteRequest struct {
	Content string `json:"content"`
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 || pathParts[3] != "notes" {
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}

	materialID, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "invalid material id", http.StatusBadRequest)
		return
	}

	// Verify material exists
	_, ok := memory.GetMaterialByID(materialID)
	if !ok {
		http.NotFound(w, r)
		return
	}

	var req createNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	req.Content = strings.TrimSpace(req.Content)
	if req.Content == "" {
		http.Error(w, "content cannot be empty", http.StatusBadRequest)
		return
	}

	newNote := note.Note{
		MaterialID: materialID,
		Content:    req.Content,
	}

	created := memory.AddNote(newNote)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(created)
}
