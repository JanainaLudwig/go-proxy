package main

import (
	"context"
	"fmt"
	"golang-proxy/model/post"
)

func main()  {
	bg := context.Background()
	p := post.NewPostRepository()

	// For user 1
	ctx := context.WithValue(bg, "auth_id", 1)

	posts, err := p.GetAll(ctx)
	fmt.Println("[All posts - USER 1]", posts, err)

	post1, err := p.GetById(ctx, 1)
	fmt.Println("[Post ID 1 - USER 1]", post1, err)

	post3, err := p.GetById(ctx, 3)
	fmt.Println("[Post ID 3 - USER 1]", post3, err)


	// For user 2
	ctx2 := context.WithValue(bg, "auth_id", 2)

	posts, err = p.GetAll(ctx2)
	fmt.Println("[All posts - USER 1]", posts, err)

	post3, err = p.GetById(ctx2, 3)
	fmt.Println("[Post ID 3 - USER 2]", post3, err)
}
