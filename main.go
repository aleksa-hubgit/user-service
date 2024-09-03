package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aleksa-hubgit/user-service/data"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	hostname := os.Getenv("DATABASE_HOSTNAME")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, hostname, port, name)
	// connStr := "postgresql://token:token@localhost:5432/token"
	fmt.Println(connStr)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	database := data.New(conn)
	defer conn.Close(context.Background())
	service := NewUserService(database)
	handler := NewUserHandler(service)
	r := gin.Default()
	r.GET("/:username", handler.GetUserByUsername)
	r.GET("/", handler.ListUsers)
	r.PUT("/", handler.UpdateUser)
	r.POST("/", handler.CreateUser)
	r.DELETE("/", handler.DeleteUser)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
