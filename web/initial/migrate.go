package initial

import (
	"gorm.io/gorm"
	"iris-test/models"
	"log"
)

func ModelMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Movie{},
		&models.MovieHistory{},
	)
	if err != nil {
		log.Fatal("数据库迁移失败", err)
	}
}
