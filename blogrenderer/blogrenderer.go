package blogrenderer

import (
	"fmt"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(writer io.Writer, post Post) error {
	_, err := fmt.Fprintf(writer, "<h1>%s</h1>", post.Title)
	if err != nil {
		return err
	}
	return nil
}
