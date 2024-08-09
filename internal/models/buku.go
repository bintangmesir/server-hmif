package models

import (
	"github.com/google/uuid"
)

type Buku struct {
	ID          uuid.UUID   `gorm:"type:char(36);primary_key;" json:"id"`	
    Judul       string      `gorm:"size:100" json:"judul" validate:"required,max=100"`
    Kode        string      `gorm:"size:10" json:"kode" validate:"required,max=10"`
    Penulis     string      `gorm:"size:100" json:"penulis" validate:"required,max=100"`
    TahunTerbit string      `gorm:"size:4" json:"tahunTerbit" validate:"required,max=4"`
    Penerbit    string      `gorm:"size:100" json:"penerbit" validate:"required,max=100"`
    Abstrak     string      `gorm:"type:text" json:"abstrak" validate:"required"`
    Jumlah      int64       `gorm:"default:0" json:"jumlah" validate:"required"`    
    Cover       *string     `gorm:"size:255" json:"cover"`
    CreatedAt   int64       `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt   int64       `gorm:"autoUpdateTime:milli" json:"updated_at"`
}