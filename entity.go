package main

type Post struct {
	ID        string `bson:"_id"`
	Title     string `bson:"title"`
	Content   string `bson:"content"`
	Author    string `bson:"author"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}
