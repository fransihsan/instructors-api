package utils

import (
	"fmt"
	"instructors-api/configs"
	C "instructors-api/entities/course"
	I "instructors-api/entities/instructor"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=%v",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
		config.DB_LOC,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	InitMigrate(db)
	return db
}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&I.Instructors{})
	db.AutoMigrate(&C.Courses{})
}
