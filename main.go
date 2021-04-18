package main

import (
	"context"
	"fmt"
	"golang-proxy/model"
)

func main()  {
	bg := context.Background()
	ctx := context.WithValue(bg, "auth_id", 1)

	p := model.PostProxy{}

	posts, err := p.GetAll(ctx)
	fmt.Println("[All posts]", posts, err)

	post, err := p.GetById(ctx, 1)
	fmt.Println("[Post of id 1]", post, err)

	post3, err := p.GetById(ctx, 3)
	fmt.Println("[Post of id 3]", post3, err)
}