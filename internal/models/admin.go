package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key;" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name" validate:"required,max=100"`
	Email       string    `gorm:"size:100;unique;not null" json:"email" validate:"required,email"`
	Password    string    `gorm:"size:255;not null" json:"-" validate:"required,min=6"`
	FotoProfile *string   `gorm:"size:255" json:"fotoProfile"`
	Role        string    `gorm:"size:20" json:"role" validate:"required,oneof=super_admin kadep_kominfo staff_kominfo kadep_prhp staff_prhp"`
	CreatedAt   int64     `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt   int64     `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

func (admin *Admin) BeforeCreate(tx *gorm.DB) (err error) {	
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin.Password = string(hashedPassword)
    
	return nil
}
