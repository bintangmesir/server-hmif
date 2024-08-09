package models

import (
	"github.com/google/uuid"
)

type ArtikelContent struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;" json:"id"`
	Index     int64       `gorm:"default:0" json:"index"`
	Tipe      string    `gorm:"size:20;not null" json:"tipe" validate:"required,oneof=sub_title description image code blockquote"`
	SubTipe   string    `gorm:"size:20" json:"subTipe"`
    Content   string    `gorm:"type:text" json:"content"`
	
	ArtikelID uuid.UUID `gorm:"type:char(36);not null" json:"artikel_id"`
	Artikel   Artikel   `gorm:"foreignKey:ArtikelID" json:"artikel"`
}