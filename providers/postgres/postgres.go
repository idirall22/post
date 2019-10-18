package provider

import (
	"context"
	"database/sql"

	"github.com/idirall22/post/models"
)

// PostgresProvider structure
type PostgresProvider struct {
	DB        *sql.DB
	TableName string
}

// New create a new post
func (p *PostgresProvider) New(ctx context.Context, content string, mediaURLs []*string, userID, groupID int64) (*models.Post, error) {
	return nil, nil
}

// Get get a single post
func (p *PostgresProvider) Get(ctx context.Context, id int64) (*models.Post, error) {
	return nil, nil
}

// List return a list of posts using limit and offset
func (p *PostgresProvider) List(ctx context.Context, limit, offset int) ([]*models.Post, error) {
	return nil, nil
}

// Update update a post
func (p *PostgresProvider) Update(ctx context.Context, content string, mediaURLs []*string, userID, id int64) (*models.Post, error) {
	return nil, nil
}

// Delete delete a post
func (p *PostgresProvider) Delete(ctx context.Context, userID, id int64) error {
	return nil
}
