package postgres

import (
	"auth-service/postgres/models"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() (*gorm.DB, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *gorm.DB
	godotenv.Load()

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	for {
		DB, err = gorm.Open(postgres.Open(dns))
		if err != nil {
			counts++
		} else {
			log.Println("Conectado com postgres")
			usersModel := &models.User{}
			if !DB.Migrator().HasTable(usersModel) {
				DB.Migrator().CreateTable(usersModel)
			}

			verificationCodeModel := &models.VerificationCode{}
			if !DB.Migrator().HasTable(verificationCodeModel) {
				DB.Migrator().CreateTable(verificationCodeModel)
			}

			connection = DB
			break
		}

		if counts > 5 {
			fmt.Println("Erro ao conectar com database")
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Erro ao conectar com database postgres. Tentando novamente em ", backOff, "segundos.")
		time.Sleep(backOff)
		continue

	}

	return connection, nil
}
