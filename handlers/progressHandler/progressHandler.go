package progressHandler

import (
	"github.com/go-chi/chi/v5"
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

	h.courseService.WordsByStatus(ctx, status, page, limit, search)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *ProgressHandler) ClearProgress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	h.courseService.ClearProgress(ctx)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *ProgressHandler) ClearWordProgress(w http.ResponseWriter, r *http.Request) {
	courseId := chi.URLParam(r, "id")
	if courseId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(courseId)
	h.courseService.ClearWordProgress(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
