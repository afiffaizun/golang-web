package handler

import (
	"net/http"
	"strings"
)

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.URL.Path, "/notes") && r.Method == http.MethodPost {
		CreateNote(w, r)
	} else {
		GetMaterialDetail(w, r)
	}
}
