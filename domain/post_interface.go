package domain

import "context"

type Post struct {
	Id int
	AuthorId int
	Title string
	Content string
}

type PostRepository interface {
	GetAll(ctx context.Context) ([]Post, error)
	GetById(ctx context.Context, id int) (Post, error)
}
