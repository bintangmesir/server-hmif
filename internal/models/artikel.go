package models

import (
	"github.com/google/uuid"
)

// Artikel represents the article model
type Artikel struct {
	ID              uuid.UUID			`gorm:"type:char(36);primary_key;" json:"id"`
	Title           string          	`gorm:"size:255;not null" json:"title"`
	SubTitle        string          	`gorm:"size:255" json:"subTitle"`
	Thumbnail       *string         	`gorm:"size:255" json:"thumbnail"`
	CommentEnabled  bool            	`gorm:"default:true" json:"commentEnabled"`
	View  			int64       		`gorm:"default:0" json:"view" validate:"required"`
	CreatedAt       int64           	`gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt       int64           	`gorm:"autoUpdateTime:milli" json:"updated_at"`

	Admins          []Admin         	`gorm:"many2many:artikel_admins;" json:"admins"`
	ArtikelContents []ArtikelContent	`gorm:"foreignKey:ArtikelID" json:"artikelContents"`
	ArtikelMetas   	[]ArtikelMeta 		`gorm:"many2many:artikel_artikel_meta;" json:"artikel_metas"`
	Comment         *Comment         	`gorm:"foreignKey:ArtikelID" json:"comment"`    
}

type ArtikelAdmin struct {
	ArtikelID uuid.UUID `gorm:"type:char(36);not null" json:"artikel_id"`
	AdminID   uuid.UUID `gorm:"type:char(36);not null" json:"admin_id"`
}