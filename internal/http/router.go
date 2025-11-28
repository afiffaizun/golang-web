package http

import (
	stdhttp "net/http"

	"github.com/afiffaizun/golang-web/internal/http/handler"
)

func NewRouter() *stdhttp.ServeMux {
	mux := stdhttp.NewServeMux()
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/materials", handler.MaterialsEntryPoint)
	mux.HandleFunc("/materials/", handler.NotesHandler)
	fs := stdhttp.FileServer(stdhttp.Dir("web/assets"))
	mux.Handle("/static/", stdhttp.StripPrefix("/static/", fs))
	mux.HandleFunc("/", handler.ListMaterials)
	return mux
}
