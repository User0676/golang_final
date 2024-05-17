package models

import "github.com/jinzhu/gorm"

type Instructor struct {
	gorm.Model
	Name           string     `gorm:"size:255;not null" json:"name"`
	ProfileSport   string     `gorm:"size:255;not null" json:"profile_sport"`
	Qualification  *string    `gorm:"size:255" json:"qualification"`
	WorkExperience *int       `json:"work_experience"`
	Trainings      []Training `gorm:"foreignkey:InstructorID" json:"trainings"` // One-to-many relationship
}
