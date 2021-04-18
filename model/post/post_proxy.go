package post

import (
	"context"
	"golang-proxy/model"
	"net/http"
)

type PostProxy Post

var r PostModel

func (p *PostProxy) GetAll(ctx context.Context) ([]Post, error) {
	posts, err := r.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	var filtered []Post

	for _, val := range posts {
		if val.AuthorId == ctx.Value("auth_id") {
			filtered = append(filtered, val)
		}
	}

	return filtered, err
}

func (p *PostProxy) GetById(ctx context.Context, id int) (*Post, error) {
	post, err := r.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, model.ModelError{
			Status:  http.StatusNotFound,
			Message: "Post not found",
		}
	}

	if post.AuthorId == ctx.Value("auth_id") {
		return post, err
	}

	return nil, model.ModelError{
		Status:  http.StatusForbidden,
		Message: "Unauthorized access",
	}
}

