package post

import (
	"context"
	"golang-proxy/db_mock"
	"golang-proxy/domain"
)

type PostModel domain.Post

func (p *PostModel) GetAll(ctx context.Context) ([]domain.Post, error) {
	return db_mock.Posts, nil
}

func (p *PostModel) GetById(ctx context.Context, id int) (*domain.Post, error) {
	for _, val := range db_mock.Posts {
		if val.Id == id {
			return &val, nil
		}
	}

	return nil, nil
}

func (p *PostModel) Delete(ctx context.Context, id int) (bool, error) {
	return false, nil
}



