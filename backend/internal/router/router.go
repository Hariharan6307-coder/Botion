package router

import (
	"backend/internal/handler"
	"backend/internal/router/middleware"

	"github.com/gorilla/mux"
)

type Router struct {
	mux *mux.Router
}

func New() *Router {
	return &Router{
		mux: mux.NewRouter(),
	}
}

func (r *Router) SetupRoutes(blockHandler *handler.BlockHandler) {
	r.mux.Use(middleware.Logger)
	r.mux.Use(middleware.CORS)
	r.mux.Use(middleware.Recovery)

	api := r.mux.PathPrefix("/api").Subrouter()

	r.setupBlockRoutes(api, blockHandler)
}

func (r *Router) setupBlockRoutes(api *mux.Router, blockHandler *handler.BlockHandler) {
	blocks := api.PathPrefix("/blocks").Subrouter()

	blocks.HandleFunc("", blockHandler.CreateBlock).Methods("POST")
	blocks.HandleFunc("/{id}", blockHandler.GetBlock).Methods("GET")
	blocks.HandleFunc("/{id}", blockHandler.UpdateBlock).Methods("PUT")
	blocks.HandleFunc("/{id}", blockHandler.DeleteBlock).Methods("DELETE")

	api.HandleFunc("/pages/{pageId}/blocks", blockHandler.GetBlocksByPage).Methods("GET")

}

func (r *Router) Handler() *mux.Router {
	return r.mux
}
