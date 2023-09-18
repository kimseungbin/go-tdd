package blogrenderer_test

import (
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"tdd/blogrenderer"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Description: "This is a description",
			Body:        "This is a post",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buffer := bytes.Buffer{}
		err := blogrenderer.Render(&buffer, aPost)
		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buffer.String())
	})
}
