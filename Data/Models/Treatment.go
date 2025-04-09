package Models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Treatment struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	PatientId uuid.UUID `gorm:"type:char(36);not null" json:"patientId"`
	Patient   Patient   `gorm:"foreignKey:PatientId"`

	Consultations []Consultation `gorm:"foreignKey:TreatmentId" json:"consultations,omitempty"` // 1:N opcional

	Exams []Exam          `gorm:"foreignKey:TreatmentId" json:"exams"`
	Drugs []Drug          `gorm:"foreignKey:TreatmentId" json:"drugs"`
	Notes []TreatmentNote `gorm:"foreignKey:TreatmentID"`

	StartDate time.Time  `json:"startDate"`
	EndDate   *time.Time `json:"endDate,omitempty"`
	Active    bool       `gorm:"default:false" json:"active"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *Treatment) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}
