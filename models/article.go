package models

import "time"


type Article struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Category  string    `json:"category"`
	Summary   string    `json:"summary"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"image_url"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
