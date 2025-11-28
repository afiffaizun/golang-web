package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/afiffaizun/golang-web/internal/storage/memory"
)

func GetMaterialDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/materials/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid material id", http.StatusBadRequest)
		return
	}

	material, ok := memory.GetMaterialByID(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(material)
}
