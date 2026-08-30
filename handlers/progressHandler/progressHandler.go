package progressHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"lingcard-go/helpers/response"
	"lingcard-go/services/courseService"
	"net/http"
	"strconv"
)

type ProgressHandler struct {
	courseService *courseService.CourseService
}

func New(courseService *courseService.CourseService) *ProgressHandler {
	return &ProgressHandler{
		courseService: courseService,
	}
}

func (h *ProgressHandler) Progress(w http.ResponseWriter, r *http.Request) {
	statusParam := chi.URLParam(r, "status")
	if statusParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	status, _ := strconv.Atoi(statusParam)
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}
	search := r.URL.Query().Get("search")

	ctx := r.Context()

	words, count, err1 := h.courseService.WordsByStatus(ctx, status, page, limit, search)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"success":     true,
		"data":        words,
		"amountWords": count,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *ProgressHandler) ClearProgress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := h.courseService.ClearProgress(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}

func (h *ProgressHandler) ClearWordProgress(w http.ResponseWriter, r *http.Request) {
	courseId := chi.URLParam(r, "id")
	if courseId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(courseId)

	err := h.courseService.ClearWordProgress(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}
