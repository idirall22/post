package post

import (
	"context"

	"github.com/idirall22/post/models"
)

// RegisterClientStream register a user to comment stream
func (s *Service) subscribeClientStream(ctx context.Context, userID, groupID int64) *models.ClientStream {

	cs := &models.ClientStream{
		Post:    make(chan *models.Post),
		UserID:  userID,
		GroupID: groupID,
	}

	s.broker.NewClient(cs)

	go func() {
		<-ctx.Done()
		s.broker.RemoveClient(cs)
	}()

	return cs
}
