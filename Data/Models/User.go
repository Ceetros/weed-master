package Models

import (
	"Api/Data/Enum"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID     `gorm:"type:char(36);primaryKey"`
	Name      string        `gorm:"type:varchar(100);not null" json:"name"`
	CPF       string        `gorm:"type:varchar(14);unique;not null" json:"cpf"`
	Role      Enum.RoleEnum `gorm:"type:varchar(20);not null" json:"role"`
	Password  string        `gorm:"type:varchar(255);not null" json:"-"`
	Email     string        `gorm:"type:varchar(100);unique;not null" json:"email"`
	BirthDate time.Time     `gorm:"not null" json:"nascimento"`

	ClinicalUser *ClinicalUser `gorm:"foreignKey:UserId" json:"-"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}
