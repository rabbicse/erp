package domain

import "github.com/google/uuid"

type Permission struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Code string    `gorm:"uniqueIndex;not null"`
}

func (p *Permission) BeforeCreate(tx any) error {
	p.ID = uuid.New()
	return nil
}
