package models

import (
	"github.com/google/uuid"
)

type Youtube struct {
	ID          uuid.UUID   `gorm:"type:char(36);primary_key;" json:"id"`	
    Judul       string      `gorm:"size:100" json:"judul" validate:"required,max=100"`
    Link        string      `gorm:"size:255" json:"link" validate:"required,url,max=100"`
    CreatedAt   int64       `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt   int64       `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}