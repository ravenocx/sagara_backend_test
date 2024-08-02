package db

import (
	"fmt"
	"log"
	"os"

	"github.com/ravenocx/clothes-store/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

var config = Config{}

func OpenConnection() (*gorm.DB, error) {
	config.Load()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode,
	)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("dsn : %s", dsn)
		log.Fatalf("Failed to open connection to database : %+v", err)
	}

	var exists bool
    conn.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'clothes_size')").Scan(&exists)

	if !exists{
		conn.Exec("CREATE TYPE clothes_size AS ENUM ('S', 'M', 'L', 'XL', 'XXL' )")
	}


	err = conn.AutoMigrate(
		&entities.Clothes{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate db : %+v", err)
	}

	return conn , err
}

func (c *Config) Load() {
	c.Host = os.Getenv("DB_HOST")
	c.User = os.Getenv("DB_USER")
	c.Password = os.Getenv("DB_PASSWORD")
	c.DBName = os.Getenv("DB_NAME")
	c.Port = os.Getenv("DB_PORT")
	c.SSLMode = os.Getenv("DB_SSLMODE")
}
