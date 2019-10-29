package memory

import (
	"testing"

	"github.com/idirall22/post/models"
)

var testMemory = &Memory{}

func TestBroker(t *testing.T) {

	r := 0
	done := make(chan bool)
	for i := 0; i < 10; i++ {

		c := &models.ClientStream{
			Post:    make(chan *models.Post, 1),
			UserID:  int64(i),
			GroupID: 1,
		}

		testMemory.NewClient(c)

		go func(cc *models.ClientStream) {
			<-c.Post
			r++
			testMemory.RemoveClient(c)

			if r >= 9 && testMemory.GetClientsLength() == 0 {
				done <- true
			}
		}(c)
	}

	post := &models.Post{ID: 1, Content: "Comment test 1", GroupID: 1}
	testMemory.Brodcast(post)
	<-done
}
