package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - a representation of the comment
// structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all of the methods
// that out service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - is the struct on which all our
// logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new
// service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (service *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retriving a comment")
	comment, err := service.Store.GetComment(ctx, id)

	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return comment, nil
}

func (service *Service) UpdateComment(ctx context.Context, comment Comment) error {
	return ErrNotImplemented
}

func (service *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (service *Service) CreateComment(ctx context.Context, comment Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
