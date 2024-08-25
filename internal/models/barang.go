package models

import (
	"github.com/google/uuid"
)

type Barang struct {
	ID          uuid.UUID   `gorm:"type:char(36);primary_key;" json:"id"`	
    Nama        string      `gorm:"size:100" json:"nama" validate:"required,max=100"`
    Jumlah      int64       `gorm:"default:0" json:"jumlah" validate:"required"`
    Baik        int64       `gorm:"default:0" json:"baik" validate:"required"`
    RusakRingan int64       `gorm:"default:0" json:"rusakRingan" validate:"required"`
    RusakBerat  int64       `gorm:"default:0" json:"rusakBerat" validate:"required"`
    Keterangan  string       `gorm:"type:text" json:"keterangan" validate:"required"`    
    CreatedAt   int64       `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt   int64       `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}