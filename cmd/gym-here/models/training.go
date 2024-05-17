package models

import (
	"github.com/jinzhu/gorm"
)

type Training struct {
	gorm.Model
	Name         string     `gorm:"size:255;not null;unique" json:"name"`
	WorkoutTime  string     `gorm:"size:255" json:"workout_time"` // format: HH:MM
	WorkoutDays  string     `gorm:"size:255" json:"workout_days"` // e.g., "Monday,Wednesday,Friday"
	InstructorID uint       `json:"instructor_id"`                // Foreign key
	Instructor   Instructor `gorm:"foreignkey:InstructorID" json:"instructor"`
	Clients      []Client   `gorm:"foreignkey:TrainingID" json:"clients"`
}
