package blogposts

import (
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dirEntries, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for range dirEntries {
		posts = append(posts, Post{})
	}
	return posts, nil
}
