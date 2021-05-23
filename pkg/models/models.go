package models

type User struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Comments []Comments `json:"comments"`
	Likes    []Likes    `json:"likes"`
}
type Comments struct {
	ID      int    `json:"id"`
	VideoID int    `json:"video_id"`
	Text    string `json:"text"`
	UserID  int    `json:"user_id"`
	Date    string `json:"date"`
}
type Likes struct {
	ID      int    `json:"id"`
	VideoID int    `json:"video_id"`
	UserID  int    `json:"user_id"`
	Date    string `json:"date"`
}
