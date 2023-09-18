package blogrenderer_test

import (
	"bytes"
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

		got := buffer.String()
		want := `<h1>hello world</h1>`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
