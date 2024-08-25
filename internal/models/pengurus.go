package models

import (
	"github.com/google/uuid"
)

type Pengurus struct {
	ID          uuid.UUID   `gorm:"type:char(36);primary_key;" json:"id"`
	Name        string      `gorm:"size:100;not null" json:"name" validate:"required,max=100"`
    Departemen  string      `gorm:"size:20;not null" json:"departemen" validate:"required,oneof=kahim_wakahim sekretaris bendahara departemen_iptek departemen_kominfo departemen_kaderisasi departemen_prhp departemen_pengmas"`
    Jabatan     string      `gorm:"size:20;not null" json:"jabatan" validate:"required,oneof=ketua_himpunan wakil_ketua_himpunan sekretaris_1 sekretaris_2 bendahara_1 bendahara_2 kepala_departemen staff_departemen"`
    Foto        *string     `gorm:"size:255" json:"foto"`
    CreatedAt   int64       `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt   int64       `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}