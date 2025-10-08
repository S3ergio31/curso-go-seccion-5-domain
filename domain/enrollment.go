package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enrollment struct {
	ID        string           `json:"id" gorm:"type:char(36);not null;primary_key:unique_index"`
	UserID    string           `json:"user_id,omitempty" gorm:"type:char(36);not null"`
	CourseID  string           `json:"course_id" gorm:"type:char(36);not null"`
	Status    EnrollmentStatus `json:"status" gorm:"type:char(2)"`
	CreatedAt *time.Time       `json:"-"`
	UpdatedAt *time.Time       `json:"-"`
}

func (u *Enrollment) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	return nil
}

type EnrollmentStatus string

const (
	Pending  EnrollmentStatus = "P"
	Active   EnrollmentStatus = "A"
	Studying EnrollmentStatus = "S"
	Inactive EnrollmentStatus = "I"
)
