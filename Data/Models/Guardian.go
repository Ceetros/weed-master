package Models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Guardian struct {
	ID    uuid.UUID `gorm:"type:char(36);primaryKey"`
	CPF   string    `gorm:"type:varchar(20);uniqueIndex"`
	Name  string    `gorm:"type:varchar(100)"`
	Phone string    `gorm:"type:varchar(20)"`
	Email string    `gorm:"type:varchar(100)"`

	// Address fields
	Street     string `gorm:"type:varchar(100)"`
	Number     string `gorm:"type:varchar(10)"`
	Complement string `gorm:"type:varchar(100)"`
	City       string `gorm:"type:varchar(50)"`
	State      string `gorm:"type:varchar(2)"`
	ZipCode    string `gorm:"type:varchar(10)"`

	Patients []Patient `gorm:"foreignKey:GuardianID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *Guardian) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}
