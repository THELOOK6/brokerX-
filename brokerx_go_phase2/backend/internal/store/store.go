
package store

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"brokerx/backend/internal/models"
)

var DB *gorm.DB
var RDB *redis.Client
var Ctx = context.Background()

func InitStore() error {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil { return fmt.Errorf("db open: %w", err) }
	DB = db
	if err := DB.AutoMigrate(&models.Account{}, &models.Order{}, &models.LedgerEntry{}, &models.AuditLog{}); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	RDB = redis.NewClient(&redis.Options{ Addr: os.Getenv("REDIS_ADDR"), Password: os.Getenv("REDIS_PASSWORD"), DB: 0 })
	if err := RDB.Ping(Ctx).Err(); err != nil { return fmt.Errorf("redis: %w", err) }
	return nil
}
