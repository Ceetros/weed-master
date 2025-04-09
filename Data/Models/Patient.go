package Models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Patient struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	Species     string    `gorm:"type:varchar(50)"`
	Name        string    `gorm:"type:varchar(100)"`
	Description string    `gorm:"type:text"`
	BirthDate   time.Time
	GuardianID  uuid.UUID `gorm:"type:char(36);not null"`

	Guardian Guardian `gorm:"foreignKey:GuardianID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *Patient) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}
