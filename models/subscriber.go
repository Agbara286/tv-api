package models

import "gorm.io/gorm"

type Subscriber struct {
	gorm.Model
	Email string `gorm:"uniqueIndex;not null" json:"email"`
}