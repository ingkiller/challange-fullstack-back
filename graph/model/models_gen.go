// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Album struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	UserID         int    `json:"userId"`
	NumberOfPhotos int    `json:"numberOfPhotos"`
}

type Comment struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
	PostID int    `json:"postId"`
}

type Photo struct {
	ID           int    `json:"id"`
	AlbumID      int    `json:"albumId"`
	Title        string `json:"title"`
	URL          string `json:"Url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type Post struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Body            string `json:"body"`
	User            *User  `json:"user"`
	NumberOfComment int    `json:"numberOfComment"`
	CreatedDate     string `json:"createdDate"`
}

type PostID struct {
	PostID int `json:"postId"`
}

type Story struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

type Task struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type User struct {
	ID       *int   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Website  string `json:"website"`
}

type UserData struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}
