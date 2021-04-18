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
	fmt.Println("[All posts - USER 1]", posts, err)

	post, err := p.GetById(ctx, 1)
	fmt.Println("[Post ID 1 - USER 1]", post, err)

	post3, err := p.GetById(ctx, 3)
	fmt.Println("[Post ID 3 - USER 1]", post3, err)

	ctx2 := context.WithValue(bg, "auth_id", 2)

	post3, err = p.GetById(ctx2, 3)
	fmt.Println("[Post ID 3 - USER 2]", post3, err)
}
