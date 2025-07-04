package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cabralfbenja/segurointeligente/internal/dtos"
	"github.com/cabralfbenja/segurointeligente/internal/service"
)

type InsuranceHandler struct {
	insuranceService service.InsuranceService
}

func NewInsuranceHandler(insuranceService service.InsuranceService) *InsuranceHandler {
	return &InsuranceHandler{
		insuranceService: insuranceService,
	}
}

func (h *InsuranceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dtos.InsuranceDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := h.insuranceService.Create(req)

	if err != nil {
		http.Error(w, "Error creating insurance: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Regla creada exitosamente",
	})

}

func (h *InsuranceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// the query parameter "userId" should be passed in the URL
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "Missing userId query parameter", http.StatusBadRequest)
		return
	}
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid userId query parameter", http.StatusBadRequest)
		return
	}
	insurances, err := h.insuranceService.GetAll(userIdInt)
	if err != nil {
		http.Error(w, "Error fetching insurances: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(insurances); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *InsuranceHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id query parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id query parameter", http.StatusBadRequest)
		return
	}

	insurance, err := h.insuranceService.Update(id)
	if err != nil {
		http.Error(w, "Error updating insurance: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(insurance)
}
