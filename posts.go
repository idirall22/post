package post

import (
	"context"

	"github.com/idirall22/post/models"
	"github.com/idirall22/utilities"
)

// add post logic
func (s *Service) addPost(ctx context.Context, form PForm) (*models.Post, error) {

	userID, err := utilities.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	if !form.ValidateForm(userID) {
		return nil, ErrorForm
	}
	// TODO: implement get iamges url
	post, err := s.provider.New(ctx, form.Content, []string{}, userID, form.GroupID)

	if err != nil {
		return nil, err
	}

	return post, nil
}

// get post logic
func (s *Service) getPost(ctx context.Context, id, groupID int64) (*models.Post, error) {

	userID, err := utilities.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	post, err := s.provider.Get(ctx, userID, id, groupID)

	if err != nil {
		return nil, err
	}

	return post, nil
}

// list posts logic
func (s *Service) listPosts(ctx context.Context, groupID int64, offset, limit int) ([]*models.Post, error) {

	userID, err := utilities.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	if err := checkIfIDIsValid(groupID); err != nil {
		return nil, err
	}

	posts, err := s.provider.List(ctx, userID, groupID, offset, limit)

	if err != nil {
		return nil, err
	}

	return posts, nil
}

// update post logic
func (s *Service) updatePost(ctx context.Context, id int64, form PForm) (*models.Post, error) {

	userID, err := utilities.GetUserIDFromContext(ctx)

	if err != nil {
		return nil, err
	}

	if !form.ValidateForm(userID) {
		return nil, ErrorForm
	}

	post, err := s.provider.Update(ctx, form.Content, []string{}, form.UserID, id)

	if err != nil {
		return nil, err
	}

	return post, nil
}

// delete post logic
func (s *Service) deletePost(ctx context.Context, id int64) error {

	userID, err := utilities.GetUserIDFromContext(ctx)

	if err != nil {
		return err
	}

	return s.provider.Delete(ctx, userID, id)
}
