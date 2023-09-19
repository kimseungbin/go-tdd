package blogrenderer_test

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"io"
	"tdd/blogrenderer"
	"testing"
)

func TestPostRenderer_Render(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Description: "This is a post",
			Body:        "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buffer := bytes.Buffer{}

		if err := postRenderer.Render(&buffer, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buffer.String())
	})

}

func BenchmarkPostRenderer_Render(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Description: "This is a post",
			Body:        "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}

func TestPostRenderer_RenderIndex(t *testing.T) {
	var posts = []blogrenderer.Post{
		{Title: "Hello World"},
		{Title: "Hello World 2"},
	}

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it renders an index of posts", func(t *testing.T) {
		buffer := bytes.Buffer{}

		if err := postRenderer.RenderIndex(&buffer, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buffer.String())
	})
}
