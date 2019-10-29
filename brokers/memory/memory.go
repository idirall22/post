package memory

import (
	"github.com/idirall22/post/models"
)

// Memory structure
type Memory struct {
	clients map[*models.ClientStream]interface{}
}

// NewClient add new client
func (m *Memory) NewClient(c *models.ClientStream) {

	if m.clients == nil {
		m.clients = make(map[*models.ClientStream]interface{})
	}

	if _, ok := m.clients[c]; !ok {
		m.clients[c] = nil
	}
}

// RemoveClient remove client
func (m *Memory) RemoveClient(c *models.ClientStream) {

	if _, ok := m.clients[c]; ok {

		close(c.Post)
		delete(m.clients, c)
	}
}

// Brodcast send comment stream to clients
func (m *Memory) Brodcast(post *models.Post) {

	for client := range m.clients {

		if client.GroupID == post.GroupID {
			client.Post <- post
		}
	}
}

// GetClientsLength get length of clients
func (m *Memory) GetClientsLength() int {
	return len(m.clients)
}
