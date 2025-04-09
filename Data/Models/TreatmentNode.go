package Models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type TreatmentNote struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	TreatmentID uuid.UUID `gorm:"type:char(36);not null"`
	Note        string    `gorm:"type:text;not null"`

	Treatment Treatment `gorm:"foreignKey:TreatmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func (u *TreatmentNote) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}
