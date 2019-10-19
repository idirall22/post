package post

import (
	"context"

	"github.com/idirall22/post/models"
)

// Provider interface
type Provider interface {
	New(context.Context, string, []string, int64, int64) (*models.Post, error)
	Get(context.Context, int64, int64, int64) (*models.Post, error)
	List(context.Context, int64, int, int) ([]*models.Post, error)
	Update(context.Context, string, []string, int64, int64) (*models.Post, error)
	Delete(context.Context, int64, int64) error
}
