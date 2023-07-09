package models

type Task struct {
	Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" gorm:"default:'pending'" enum:"pending,done,deleted,cancelled"`
	UserID      uint   `json:"user_id" binding:"required" gorm:"foreignKey:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE; references:ID;"`
	User        User   `json:"user"`
}
