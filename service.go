package post

import (
	"database/sql"
	"time"

	pr "github.com/idirall22/post/providers/postgres"
)

// TimeoutRequest time to stop a request
var TimeoutRequest = time.Second * 5

var (
	// DefaultOffset offset used in db
	DefaultOffset = 5
	// MaxOffset offset used in db
	MaxOffset = 10
)

// Service structure
type Service struct {
	provider Provider
}

// NewService create a new  post service
func NewService(db *sql.DB, tableName string) *Service {
	return &Service{provider: &pr.PostgresProvider{DB: db, TableName: tableName}}
}
