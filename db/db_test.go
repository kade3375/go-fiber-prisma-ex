package db

import (
	"context"
	"encoding/json"
	"testing"
)

func TestORM(t *testing.T) {
	client := NewClient()
	if err := client.Prisma.Connect(); err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		if err := client.Prisma.Disconnect(); err != nil {
			t.Error(err)
		}
	})

	ctx := context.Background()

	// create a post
	createdPost, err := client.Post.CreateOne(
		Post.Title.Set("Hi from Prisma!"),
		Post.Published.Set(true),
		Post.Desc.Set("Prisma is a database toolkit and makes databases easy."),
	).Exec(ctx)
	if err != nil {
		t.Error(err)
	}

	result, _ := json.MarshalIndent(createdPost, "", "  ")
	t.Logf("created post: %s\n", result)

	// find a single post
	post, err := client.Post.FindUnique(
		Post.ID.Equals(createdPost.ID),
	).Exec(ctx)
	if err != nil {
		t.Error(err)
	}

	result, _ = json.MarshalIndent(post, "", "  ")
	t.Logf("post: %s\n", result)

	desc, ok := post.Desc()
	if !ok {
		t.Error("post's description is null")
	}

	t.Logf("The posts's description is: %s\n", desc)
}