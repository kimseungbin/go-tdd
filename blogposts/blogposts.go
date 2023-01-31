package blogposts

import (
	"bufio"
	"io"
	"io/fs"
)

type Post struct {
	Title       string
	Description string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dirEntries, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, file := range dirEntries {
		post, err := getPost(fileSystem, file.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readline := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	title := readline()[7:]
	description := readline()[13:]

	return Post{Title: title, Description: description}, nil
}
