package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

type PostRenderer struct {
	templ *template.Template
}

func NewTestRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(writer io.Writer, post Post) error {
	if err := r.templ.ExecuteTemplate(writer, "blog.gohtml", post); err != nil {
		return err
	}

	return nil
}

func (r *PostRenderer) RenderIndex(writer io.Writer, posts []Post) error {
	return nil
}
