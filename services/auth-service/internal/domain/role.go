package domain

import "github.com/google/uuid"

type Role struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string    `gorm:"index;not null"`
	TenantID uuid.UUID `gorm:"type:uuid;index;not null"`

	Permissions []Permission `gorm:"many2many:role_permissions"`
}

func (r *Role) BeforeCreate(tx any) error {
	r.ID = uuid.New()
	return nil
}
