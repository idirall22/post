package post

import (
	"database/sql"

	pr "github.com/idirall22/post/providers/postgres"
)

// Service structure
type Service struct {
	provider Provider
}

// NewService create a new  post service
func NewService(db *sql.DB, tableName string) *Service {
	return &Service{provider: &pr.PostgresProvider{DB: db, TableName: tableName}}
}
