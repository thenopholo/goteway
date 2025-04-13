package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thenopholo/goteway/internal/repository"
	"github.com/thenopholo/goteway/internal/service"
	"github.com/thenopholo/goteway/internal/web/server"
)

func getEnv(key, defautlValue string) string {
  if value := os.Getenv(key); value != "" {
    return value
  }
  return defautlValue
}

func main() {
  if err := godotenv.Load(); err != nil{
    log.Fatal("Error loading .env file")
  }

  connStr := fmt.Sprintf(
    "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
    getEnv("DB_HOST", "db"),
    getEnv("DB_PORT", "5432"),
    getEnv("DB_USER", "postgres"),
    getEnv("DB_PASSWORD", "postgres"),
    getEnv("DB_NAME", "goteway"),
    getEnv("DB_SSL_MODE", "disable"),
  )

  db, err := sql.Open("potsgres", connStr)
  if err != nil {
    log.Fatal("Error connecting to database", err)
  }
  defer db.Close()

  accountRepository := repository.NewAccountRepository(db)
  accountService := service.NewAccountService(accountRepository)

  port := getEnv("HTTP_PORT", "8080")
  srv := server.NewServer(accountService, port)
  srv.CofigureRoutes()

  if err := srv.Start(); err != nil {
    log.Fatal("Error starting server: ", err)
  }
}