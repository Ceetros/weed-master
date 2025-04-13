package Models

import (
	"Api/Data/Enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Clinical struct {
	ID          uuid.UUID        `gorm:"type:char(36);primaryKey" json:"id"`
	Name        string           `gorm:"type:varchar(100);not null" json:"name"`
	Document    string           `gorm:"type:varchar(20);unique;not null" json:"document"`
	Role        Enum.RoleEnum    `gorm:"type:varchar(20);not null" json:"role"`
	NextPayment time.Time        `gorm:"not null" json:"nextPayment"`
	PixKey      string           `gorm:"type:varchar(100)" json:"pixKey"`
	PixType     Enum.PixTypeEnum `gorm:"type:varchar(20)" json:"pixType"`
	Address     Address          `gorm:"foreignKey:ClinicalID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	ClinicalUsers []ClinicalUser `gorm:"foreignKey:ClinicalId"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *Clinical) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}
