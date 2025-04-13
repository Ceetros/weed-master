package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ClinicalUser struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserId     uuid.UUID `gorm:"type:char(36);not null"`
	ClinicalId uuid.UUID `gorm:"type:char(36);not null"`

	User     User     `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Clinical Clinical `gorm:"foreignKey:ClinicalId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *ClinicalUser) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}
