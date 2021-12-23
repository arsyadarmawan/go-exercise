package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `gorm:"primarykey;autoIncrement" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"createdAt"`
	Name      string         `json:"name"`
	Email     string         `gorm:"not null;unique" json:"email"`
	Password  string         `json:"password"`
}
