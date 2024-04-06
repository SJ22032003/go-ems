package model

// import "time"

type Event struct {
	ID          int64  `json:"id"`
	Title       string `binding:"required" json:"title"`
	Description string `binding:"required" json:"description"`
	Location    string `binding:"required" json:"location"`
	Date        string `json:"date"`
	UserId      int64  `json:"user_id"`
}
