package model

import (
	"context"
	"golang-proxy/domain"
	"golang-proxy/model/post"
	"net/http"
)

type PostProxy domain.Post

var r post.PostModel

func (p *PostProxy) GetAll(ctx context.Context) ([]domain.Post, error) {
	posts, err := r.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	var filtered []domain.Post

	for _, val := range posts {
		if val.AuthorId == ctx.Value("auth_id") {
			filtered = append(filtered, val)
		}
	}

	return filtered, err
}

func (p *PostProxy) GetById(ctx context.Context, id int) (*domain.Post, error) {
	post, err := r.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, ModelError{
			Status:  http.StatusNotFound,
			Message: "Post not found",
		}
	}

	if post.AuthorId == ctx.Value("auth_id") {
		return post, err
	}

	return nil, ModelError{
		Status:  http.StatusForbidden,
		Message: "Unauthorized access",
	}
}

func (p *PostProxy) Delete(ctx context.Context, id int) (bool, error) {
	return r.Delete(ctx, id)
}



