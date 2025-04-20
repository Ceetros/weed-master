package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Analytics struct {
	ID      uuid.UUID `gorm:"type:char(36);primaryKey"`
	WeedId  string    `gorm:"type:varchar(100);not null" json:"weedId"`
	Umidity int       `gorm:"not null;type:int" json:"umidity"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *Analytics) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}
