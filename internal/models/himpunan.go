package models

import "github.com/google/uuid"

type Himpunan struct {
    ID                uuid.UUID `gorm:"type:char(36);primary_key;" json:"id"`
    JumlahPengurus    int64     `gorm:"not null" validate:"required" json:"jumlahPengurus"`
    JumlahMahasiswa   int64     `gorm:"not null" validate:"required" json:"jumlahMahasiswa"`
    JumlahDepartemen  int64     `gorm:"not null" validate:"required" json:"jumlahDepartemen"`
    NamaProker        string    `gorm:"size:255;not null" json:"namaProker"`
    GaleriMahasiswa   *string   `gorm:"size:255" json:"galeriMahasiswa"`
    CreatedAt         int64     `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt         int64     `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}