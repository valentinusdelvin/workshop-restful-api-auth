package entity

import "github.com/google/uuid"

type User struct {
	UserId   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email    string    `gorm:"type:varchar(100);unique"`
	Password string    `gorm:"type:varchar(255)"`
	Role     string    `gorm:"type:varchar(20);not null"`
}
