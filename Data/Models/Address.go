package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Address struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey"`
	Street     string    `gorm:"type:varchar(100)"`
	Number     string    `gorm:"type:varchar(10)"`
	Complement string    `gorm:"type:varchar(100)"`
	City       string    `gorm:"type:varchar(50)"`
	State      string    `gorm:"type:varchar(2)"`
	ZipCode    string    `gorm:"type:varchar(10)"`

	ClinicalID uuid.UUID `gorm:"type:char(36);uniqueIndex"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *Address) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}
