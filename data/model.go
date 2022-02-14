package data

import "time"

const (
	Admin = "933268035685449789"
)

type (
	BlackList struct {
		Id        int        `gorm:"primary_key" json:"id"`
		Word      string     `gorm:"type:varchar(255);not null" json:"word"`
		CreatedAt *time.Time `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
		DeletedAt *time.Time `json:"-"`
	}
)
