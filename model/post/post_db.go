package post

import (
	"context"
)

type PostModel Post

func (p *PostModel) GetAll(ctx context.Context) ([]Post, error) {
	return Posts, nil
}

func (p *PostModel) GetById(ctx context.Context, id int) (*Post, error) {
	for _, val := range Posts {
		if val.Id == id {
			return &val, nil
		}
	}

	return nil, nil
}
