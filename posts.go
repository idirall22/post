package post

import (
	"context"

	"github.com/idirall22/post/models"
)

// add post logic
func (s *Service) addPost(ctx context.Context, form PForm) (*models.Post, error) {
	return nil, nil
}

// get post logic
func (s *Service) getPost(ctx context.Context, id, groupID int64) (*models.Post, error) {
	return nil, nil
}

// list posts logic
func (s *Service) listPosts(ctx context.Context, groupID int64, offset, limit int) ([]*models.Post, error) {
	return nil, nil
}

// update post logic
func (s *Service) updatePost(ctx context.Context, form PForm) (*models.Post, error) {
	return nil, nil
}

// delete post logic
func (s *Service) deletePost(ctx context.Context, id int64) error {
	return nil
}
