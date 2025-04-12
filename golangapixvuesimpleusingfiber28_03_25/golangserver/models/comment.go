package models

type Comment struct {
	PostID int    `json:"postId"` // ID postingan yang memiliki komentar
	ID     int    `json:"id"`     // ID unik komentar
	Name   string `json:"name"`   // Nama pemberi komentar
	Email  string `json:"email"`  // Email pemberi komentar
	Body   string `json:"body"`   // Isi komentar
}
