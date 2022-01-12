package entity

type BlogEntry struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	CreatedDate int    `json:"createddate"`
}
