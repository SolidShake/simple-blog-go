package db

import (
	"log"
	"time"
)

type Post struct {
	PostID    int        `db:"post_id"`
	Title     string     `db:"title"`
	Content   string     `db:"content"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func CreatePost(postTitle, postContent string) (err error) {
	_, err = DB.Exec("INSERT INTO post (title, content, created_at) VALUES (?, ?, NOW())", postTitle, postContent)
	if err != nil {
		log.Println(err)
	}

	return
}

func GetPostByID(id string) (post Post, err error) {
	err = DB.Get(&post, "SELECT * FROM post where post_id="+id)
	if err != nil {
		log.Println(err)
	}

	return post, err
}

func GetAllPosts() ([]Post, error) {
	posts := []Post{}
	err := DB.Select(&posts, "SELECT * FROM post order by post_id desc")
	if err != nil {
		log.Println(err)
	}

	return posts, err
}

func GetAllPostsByTag(tag string) ([]Post, error) {
	posts := []Post{}
	err := DB.Select(&posts, "SELECT * FROM post where title=\""+tag+"\" order by post_id desc")
	if err != nil {
		log.Println(err)
	}

	return posts, err
}
