package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username string    `gorm:"uniqueIndex;not null"`
	Email    string    `gorm:"uniqueIndex"`
	Password string    `gorm:"not null"`

	TenantID uuid.UUID `gorm:"type:uuid;index;not null"`
	IsActive bool      `gorm:"default:true"`

	Roles     []Role `gorm:"many2many:user_roles"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
