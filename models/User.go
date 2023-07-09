package models

type User struct {
	Model
	Username string `json:"username" binding:"required" gorm:"unique"`
	Password string `json:"password,omitempty" binding:"required"`
}
