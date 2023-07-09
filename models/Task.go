package models

type Task struct {
	Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status" enum:"pending,done,deleted,cancelled"`
	UserID      uint   `json:"user_id" binding:"required" gorm:"not null"`
	User        *User  `json:"user,omitempty"`
}
