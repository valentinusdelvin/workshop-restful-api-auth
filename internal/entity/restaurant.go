package entity

import (
	"time"

	"github.com/google/uuid"
)

type Restaurant struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Location  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	Items     []Item    `gorm:"not null;foreignKey:RestaurantId"`
}
