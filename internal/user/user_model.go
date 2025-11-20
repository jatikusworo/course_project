package user

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255" json:"name"`
	Email     string    `gorm:"uniqueIndex;size:255" json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
