package postHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"lingcard-go/dto/post"
	"lingcard-go/dto/request"
	"lingcard-go/helpers/response"
	"lingcard-go/services/postService"
	"net/http"
	"strconv"
)

type PostHandler struct {
	postService *postService.PostService
}

func New(postService *postService.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) GetPostsByCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	if code == "" {
		code = "en"
	}
	posts, err := h.postService.GetPostsByCode(code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"success": true,
		"data":    posts,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func (h *PostHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")
	if postId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(postId)
	ctx := r.Context()
	pst, err := h.postService.GetOne(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"success": true,
		"data":    pst,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func (h *PostHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var requestDTO request.CommentRequestDTO
	commentId := chi.URLParam(r, "id")
	if commentId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(commentId)
	body, _ := io.ReadAll(r.Body)
	errJSON := json.Unmarshal(body, &requestDTO)
	if errJSON != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	ctx := r.Context()
	err := h.postService.CreateComment(ctx, id, requestDTO.Text)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var requestDTO post.CreatePostDTO
	body, _ := io.ReadAll(r.Body)
	errJSON := json.Unmarshal(body, &requestDTO)
	if errJSON != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	ctx := r.Context()
	err := h.postService.Create(ctx, requestDTO)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}

func (h *PostHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentId := chi.URLParam(r, "id")
	if commentId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(commentId)
	ctx := r.Context()
	err := h.postService.DeleteComment(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}
