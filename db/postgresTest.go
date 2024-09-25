package db

import (
	"authApp/models"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewTestDBConn(ctx context.Context, t *testing.T) *gorm.DB {
	t.Run("Test DB Connection", func(t *testing.T) {
		container, err := postgres.Run(
			ctx,
			"postgres:16",
			postgres.WithUsername("test"),
			postgres.WithPassword("test"),
			postgres.WithDatabase("testdb"),
			testcontainers.WithWaitStrategy(
				wait.ForLog("database system is ready to accept connections").
					WithOccurrence(2).
					WithStartupTimeout(5*time.Second),
			),
		)
		assert.NoError(t, err)

		connURI, err := container.ConnectionString(ctx, "sslmode=disable")
		assert.NoError(t, err)

		DB, err = gorm.Open(driver.Open(connURI), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		assert.NoError(t, err)

		err = DB.Migrator().CreateTable(&models.User{})
		assert.NoError(t, err)
	})

	return DB
}
