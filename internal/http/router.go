package http

import (
	stdhttp "net/http"

	"github.com/afiffaizun/golang-web/internal/http/handler"
)

func NewRouter() *stdhttp.ServeMux {
	mux := stdhttp.NewServeMux()
	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/materials", handler.ListMaterials)
	mux.HandleFunc("/materials", handler.GetMaterialDetail)
	return mux
}
