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
	_, err := fmt.Fprintf(writer, "<h1>%s</h1>\n<p>%s</p>\n", post.Title, post.Description)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(writer, "Tags: <ul>")
	if err != nil {
		return err
	}

	for _, tag := range post.Tags {
		_, err = fmt.Fprintf(writer, "<li>%s</li>", tag)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(writer, "</ul>")
	if err != nil {
		return err
	}

	return nil
}
