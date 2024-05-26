// Post model
package models

type Post struct {
	ID    int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title string `gorm:"column:title" json:"title"`
	Body  string `gorm:"column:body" json:"body"`
}

func GetPost() Post {
	var post Post
	return post
}

func GetPosts() []Post {
	var posts []Post
	return posts
}
