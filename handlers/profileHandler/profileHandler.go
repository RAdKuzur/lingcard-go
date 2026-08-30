package profileHandler

import (
	"encoding/json"
	"io"
	"lingcard-go/dto/request"
	"lingcard-go/helpers/response"
	"lingcard-go/services/courseService"
	"lingcard-go/services/userService"
	"net/http"
)

type ProfileHandler struct {
	userService   *userService.UserService
	courseService *courseService.CourseService
}

func New(userService *userService.UserService, courseService *courseService.CourseService) *ProfileHandler {
	return &ProfileHandler{userService: userService, courseService: courseService}
}

func (h *ProfileHandler) Profile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	profile, err := h.userService.Profile(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"success": true,
		"data":    profile,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *ProfileHandler) Update(w http.ResponseWriter, r *http.Request) {
	var requestDTO request.ProfileUpdateDTO
	body, _ := io.ReadAll(r.Body)
	errJSON := json.Unmarshal(body, &requestDTO)
	if errJSON != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	ctx := r.Context()
	err1 := h.courseService.ClearProgress(ctx)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err2 := h.userService.Update(ctx, requestDTO)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}
