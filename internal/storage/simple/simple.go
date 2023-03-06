package simple

import (
	"github.com/dollarkillerx/common/pkg/client"
	"github.com/dollarkillerx/common/pkg/conf"
	"github.com/dollarkillerx/eim/internal/pkg/models"
	"gorm.io/gorm"

	"sync"
)

type Simple struct {
	db *gorm.DB

	inventoryMu sync.Mutex
}

func NewSimple(config conf.PostgresConfiguration) (*Simple, error) {
	postgresClient, err := client.PostgresClient(config, nil)
	if err != nil {
		return nil, err
	}

	postgresClient.AutoMigrate(
		&models.User{},
		&models.Friendship{},
	)

	return &Simple{
		db: postgresClient,
	}, nil
}

func (s *Simple) DB() *gorm.DB {
	return s.db
}
