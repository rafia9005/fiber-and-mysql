package entity

import "time"

type Users struct {
	ID        uint
	Name      string    `json:"name"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password"`
	UpdateAt  time.Time `json:"created_at"`
	CreatedAt time.Time `json:"created_at"`
}
