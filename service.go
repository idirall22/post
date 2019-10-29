package post

import (
	"database/sql"

	"github.com/idirall22/post/brokers/memory"
	pr "github.com/idirall22/post/providers/postgres"
)

// Service structure
type Service struct {
	provider Provider
	broker   Broker
}

// NewService create a new  post service
func StartService(db *sql.DB, tableName string) *Service {
	return &Service{
		provider: &pr.PostgresProvider{DB: db, TableName: tableName},
		broker:   &memory.Memory{},
	}
}
