package initialize

import (
	"fmt"
	"server/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func DBConnection(){
 dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", ENV_DB_USER, ENV_DB_PASSWORD, ENV_DB_URI, ENV_DB_PORT, ENV_DB_NAME)
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
    NamingStrategy:  schema.NamingStrategy{
        SingularTable: true,
    },
  })

  if err != nil {
    panic("Cannot connect to database...")
  }

  db.AutoMigrate(
    &models.Admin{},
		&models.Artikel{},
		&models.ArtikelContent{},
		&models.Comment{},
    &models.Alumni{},
    &models.Barang{},
    &models.Buku{},
    &models.Youtube{},
    &models.Pengurus{},
    &models.ArtikelMeta{},
    &models.Himpunan{},
	)

  DB = db
}