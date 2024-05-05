package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"QueueOptimization/models"
	"QueueOptimization/service"
)

type QueueHandler struct {
	service *service.CRUDQueueService
}

func NewQueueHandler(service *service.CRUDQueueService) *QueueHandler {
	return &QueueHandler{service: service}
}

func (h *QueueHandler) CreateQueue(w http.ResponseWriter, r *http.Request) {
	var queue models.Queue
	if err := json.NewDecoder(r.Body).Decode(&queue); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.service.CreateQueue(queue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (h *QueueHandler) GetAllQueues(w http.ResponseWriter, r *http.Request) {
	queues, err := h.service.GetAllQueues()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(queues)
}

func (h *QueueHandler) GetQueueById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid queue ID", http.StatusBadRequest)
		return
	}
	queue, err := h.service.GetQueueById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(queue)
}

func (h *QueueHandler) UpdateQueue(w http.ResponseWriter, r *http.Request) {
	var queue models.Queue
	if err := json.NewDecoder(r.Body).Decode(&queue); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.UpdateQueue(queue); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *QueueHandler) DeleteQueue(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid queue ID", http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteQueue(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

