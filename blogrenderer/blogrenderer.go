package blogrenderer

import "io"

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(writer io.Writer, post Post) error {
	return nil
}
