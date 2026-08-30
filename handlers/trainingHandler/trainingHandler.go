package trainingHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"lingcard-go/dto/request"
	"lingcard-go/helpers/response"
	"lingcard-go/services/courseService"
	"net/http"
	"strconv"
)

type TrainingHandler struct {
	courseService *courseService.CourseService
}

func New(courseService *courseService.CourseService) *TrainingHandler {
	return &TrainingHandler{
		courseService: courseService,
	}
}

func (h *TrainingHandler) NewWord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	word, err := h.courseService.NewWord(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"success": true,
		"data":    word,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *TrainingHandler) RepeatWord(w http.ResponseWriter, r *http.Request) {
	var status request.TrainingRequestDTO
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &status)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	ctx := r.Context()

	err1 := h.courseService.RepeatWord(ctx, id, status.Status)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}

func (h *TrainingHandler) Teachable(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code, status, err := h.courseService.Teachable(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"language": code,
			"training": status,
		},
	}
	json.NewEncoder(w).Encode(resp)
}
