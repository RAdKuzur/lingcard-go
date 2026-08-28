package postHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"lingcard-go/dto/post"
	"lingcard-go/dto/request"
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
	posts := h.postService.GetPostsByCode(code)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}

func (h *PostHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "id")
	if postId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(postId)
	ctx := r.Context()
	post := h.postService.GetOne(ctx, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)

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
	json.Unmarshal(body, &requestDTO)
	ctx := r.Context()
	h.postService.CreateComment(ctx, id, requestDTO.Text)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var requestDTO post.CreatePostDTO
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &requestDTO)
	ctx := r.Context()
	h.postService.Create(ctx, requestDTO)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *PostHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentId := chi.URLParam(r, "id")
	if commentId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(commentId)
	ctx := r.Context()
	h.postService.DeleteComment(ctx, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
