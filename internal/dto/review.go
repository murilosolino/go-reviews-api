package dto

type Review struct {
	Id         int64  `json:"id"`
	Review     string `json:"review"`
	AuthorName string `json:"author"`
	Url_photo  string `json:"url_photo"`
}
