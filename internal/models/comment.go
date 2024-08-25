package models

import (
	"github.com/google/uuid"
)

type Comment struct {
	ID        	uuid.UUID	`gorm:"type:char(36);primary_key;" json:"id"`
	Text      	string     	`gorm:"type:text;not null" json:"text" validate:"required"`
	Image     	*string    	`gorm:"size:255" json:"image"`	
	Email     	string     	`gorm:"size:100;not null" json:"email" validate:"required,email"`
    CreatedAt   int64       `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt   int64       `gorm:"autoUpdateTime:milli" json:"updatedAt"`

	ArtikelID 	uuid.UUID  	`gorm:"type:char(36);not null" json:"artikel_id"`
	Artikel   	Artikel    	`gorm:"foreignKey:ArtikelID" json:"artikel"`
	ParentID  	*uuid.UUID 	`gorm:"type:char(36);index" json:"parent_id"`
	Parent    	*Comment   	`gorm:"foreignKey:ParentID" json:"parent"`
	Replies   	[]Comment  	`gorm:"foreignKey:ParentID" json:"replies"`
}