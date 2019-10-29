package post

import "github.com/idirall22/post/models"

// Broker interface
type Broker interface {
	NewClient(c *models.ClientStream)
	RemoveClient(c *models.ClientStream)
	Brodcast(comment *models.Post)
}
