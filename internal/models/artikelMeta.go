package models

import "github.com/google/uuid"

type ArtikelMeta struct {
	ID        uuid.UUID  `gorm:"type:char(36);primary_key;" json:"id"`
	Like      int64      `gorm:"default:0" json:"like"`
	Email     string     `gorm:"size:255;not null" json:"email"`
	CreatedAt int64      `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64      `gorm:"autoUpdateTime:milli" json:"updated_at"`

	Artikels  []Artikel  `gorm:"many2many:artikel_artikel_meta;" json:"artikels"`
}