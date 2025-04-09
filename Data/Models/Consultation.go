package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Consultation struct {
	ID           uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Observations string    `gorm:"type:text" json:"observations"`

	TreatmentId *uuid.UUID `gorm:"type:char(36)" json:"treatmentId,omitempty"` // opcional
	Treatment   *Treatment `gorm:"foreignKey:TreatmentId" json:"treatment,omitempty"`

	Exams     []Exam    `gorm:"foreignKey:ConsultationId" json:"exams,omitempty"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u *Consultation) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}
