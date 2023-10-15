package main

import (
	"fmt"
	interfaces "goclean/pkg/v1"
	repo "goclean/pkg/v1/repository"
	usecase "goclean/pkg/v1/usecase"
	"log"
	"net"

	"gorm.io/gorm"

	dbConfig "goclean/internal/db"
	"goclean/internal/models"
	handler "goclean/pkg/v1/handler/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db := dbConfig.DbConn()
	migrations(db)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error Starting the Server: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	userUseCase := initUserServer(db)
	handler.NewServer(grpcServer, userUseCase)

	log.Fatal(grpcServer.Serve(lis))
}

func initUserServer(db *gorm.DB) interfaces.UseCaseInterface {
	userRepo := repo.New(db)
	return usecase.New(userRepo)
}
func migrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Migrated")
	}
}
