package handler

import (
	"backend/internal/models"
	"backend/internal/service"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type BlockHandler struct {
	service *service.BlockService
}

func NewBlockHandler(service *service.BlockService) *BlockHandler {
	return &BlockHandler{service: service}
}

func (h *BlockHandler) CreateBlock(w http.ResponseWriter, r *http.Request) {
	var req models.CreateBlockRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	block, err := h.service.CreateBlock(r.Context(), req)
	if err != nil {
		 http.Error(w, err.Error(), http.StatusInternalServerError)
		 return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(block)
}

func (h *BlockHandler) GetBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]

    block, err := h.service.GetBlock(r.Context(), id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(block)
}

func (h *BlockHandler) GetBlocksByPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["pageId"]

	blocks, err := h.service.GetBlocksByPage(r.Context(), pageID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blocks)
}

func (h *BlockHandler) UpdateBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req models.UpdateBlockRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	block, err := h.service.UpdateBlock(r.Context(), id, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}	

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(block)
}

func (h *BlockHandler) DeleteBlock(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    err := h.service.DeleteBlock(r.Context(), id)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}