package db

import (
	"log"

	//"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open("host=localhost user=postgres dbname=goclean password=admin port=5433 sslmode=disable"), &gorm.Config{},
	)
	if err != nil {
		log.Fatalf("There was error connecting to the database: %v", err)
	}
	return db
}

// func DbConn() *gorm.DB {
// 	dsn := "root:admin@tcp(127.0.0.1:3306)/gomovies?charset=utf8mb4&parseTime=True&loc=Local"
// 	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Error Connecting to the database: %v", err)
// 	}
// 	return d
// }
