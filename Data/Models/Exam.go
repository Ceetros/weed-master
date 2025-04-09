package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Exam struct {
	ID           uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name         string    `json:"name"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	Observations string    `gorm:"type:text" json:"observations"`

	TreatmentId    uuid.UUID  `gorm:"type:char(36)" json:"treatmentId"`
	ConsultationId *uuid.UUID `gorm:"type:char(36)" json:"consultationId,omitempty"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *Exam) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}
