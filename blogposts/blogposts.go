package blogposts

import (
	"io/fs"
)

type Post struct {
}

func NewPostsFromFS(fileSystem fs.FS) []Post {
	dirEntries, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for range dirEntries {
		posts = append(posts, Post{})
	}
	return posts
}
