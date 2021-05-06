package db

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
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

	t.Logf("createdPost: %+v", createdPost)

	// then create a comment
	comments, err := client.Comment.CreateOne(
		Comment.Content.Set("my description"),
		// link the post we created before
		Comment.Post.Link(
			Post.ID.Equals(createdPost.ID),
		),
	).Exec(ctx)
	if err != nil {
		t.Error(err)
	}

	t.Logf("comments: %+v", comments)

	post, err := client.Post.FindUnique(
		Post.ID.Equals(createdPost.ID),
	).With(
		Post.Comments.Fetch(),
	).Exec(ctx)
	if err != nil {
		t.Error(err)
	}

	result, _ := json.MarshalIndent(post, "", "  ")
	t.Logf("post with comment: %s\n", result)

	assert.Equal(t, len(post.Comments()), 1, "Comment Added")
}