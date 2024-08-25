package models

import (
	"github.com/google/uuid"
)

type Alumni struct {
	ID              uuid.UUID `gorm:"type:char(36);primary_key;" json:"id"`	
	Angkatan        string    `gorm:"size:4" json:"angkatan" validate:"required,max=4"`
    Nama            string    `gorm:"size:100" json:"nama" validate:"required,max=100"`
    NoTelephone     string    `gorm:"size:15" json:"noTelephone" validate:"required,max=15"`
    CreatedAt       int64     `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt       int64     `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}