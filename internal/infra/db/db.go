package db

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	conn *gorm.DB
	once sync.Once
)

// Connect inicializa la conexión a Postgres leyendo variables de entorno.
func Connect() (*gorm.DB, error) {
	var err error

	once.Do(func() {
		// Allow full DATABASE_URL (postgres URI) to be used if provided.
		databaseURL := os.Getenv("DATABASE_URL")
		fmt.Println("DATABASE_URL:", databaseURL)
		var dsn string
		if databaseURL != "" {
			dsn = databaseURL
		} else {
			host := os.Getenv("DB_HOST")
			port := os.Getenv("DB_PORT")
			user := os.Getenv("DB_USERNAME")
			pass := os.Getenv("DB_PASSWORD")
			dbname := os.Getenv("DB_DATABASE")
			ssl := os.Getenv("DB_SSLMODE")

			if ssl == "" {
				ssl = "disable"
			}
			if port == "" {
				port = "5432"
			}

			// Validate required env vars to avoid malformed DSN
			missing := []string{}
			if host == "" {
				missing = append(missing, "DB_HOST")
			}
			if user == "" {
				missing = append(missing, "DB_USERNAME")
			}
			if dbname == "" {
				missing = append(missing, "DB_DATABASE")
			}
			if len(missing) > 0 {
				err = fmt.Errorf("missing required DB env vars: %v", missing)
				conn = nil
				return
			}

			dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, pass, dbname, port, ssl)
		}

		conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			conn = nil
		}
	})
	return conn, err
}

// Get retorna la conexión (la inicializa si no existe).
func Get() (*gorm.DB, error) {
	if conn != nil {
		return conn, nil
	}
	return Connect()
}
