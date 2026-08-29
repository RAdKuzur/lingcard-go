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

	profile := h.userService.Profile(ctx)
	Response := map[string]interface{}{
		"success": true,
		"data":    profile,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response)
}

func (h *ProfileHandler) Update(w http.ResponseWriter, r *http.Request) {
	var requestDTO request.ProfileUpdateDTO
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &requestDTO)
	ctx := r.Context()

	h.courseService.ClearProgress(ctx)
	h.userService.Update(ctx, requestDTO)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}
