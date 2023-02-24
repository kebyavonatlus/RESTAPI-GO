package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kebyavonatlus/gorestapicourse/internal/comment"
)

type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
	UpdateComment(ctx context.Context, ID string, newComment comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
}

type Response struct {
	Message string
}

func (handler *Handler) PostComment(response http.ResponseWriter, request *http.Request) {
	var cmt comment.Comment
	if err := json.NewDecoder(request.Body).Decode(&cmt); err != nil {
		return
	}

	cmt, err := handler.Service.PostComment(request.Context(), cmt)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(response).Encode(cmt); err != nil {
		panic(err)
	}
}

func (handler *Handler) GetComment(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := handler.Service.GetComment(request.Context(), id)
	if err != nil {
		log.Print(err)
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(response).Encode(cmt); err != nil {
		panic(err)
	}
}

func (handler *Handler) UpdateComment(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if id == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	var cmt comment.Comment
	if err := json.NewDecoder(request.Body).Decode(&cmt); err != nil {
		return
	}

	cmt, err := handler.Service.UpdateComment(request.Context(), id, cmt)
	if err != nil {
		log.Print(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(response).Encode(cmt); err != nil {
		panic(err)
	}
}

func (handler *Handler) DeleteComment(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	commentId := vars["id"]
	if commentId == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err := handler.Service.DeleteComment(request.Context(), commentId)

	if err != nil {
		log.Print(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(response).Encode(Response{Message: "Successfully deleted"}); err != nil {
		panic(err)
	}
}
