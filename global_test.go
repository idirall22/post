package post

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/idirall22/post/brokers/memory"
	pr "github.com/idirall22/post/providers/postgres"
	"github.com/idirall22/utilities"
	_ "github.com/lib/pq"
)

var (
	testService *Service
	database    *sql.DB
	testToken   string
	tableName   = "posts"
	query       = fmt.Sprintf(`
	DROP TABLE IF EXISTS %s CASCADE;

	CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		content VARCHAR NOT NULL,
		media_urls text[],
		user_id INTEGER REFERENCES users(id),
		group_id INTEGER REFERENCES groups(id),
		created_at TIMESTAMP with TIME ZONE DEFAULT now(),
		deleted_at TIMESTAMP DEFAULT NULL
	);
	`, tableName, tableName)
)

// TestGlobal run tests
func TestGlobal(t *testing.T) {

	db, err := utilities.ConnectDataBaseTest()

	if err != nil {
		t.Error(err)
		return
	}

	err = utilities.BuildDataBase(db, query)

	if err != nil {
		t.Error(err)
		return
	}

	defer utilities.CloseDataBaseTest(db)

	provider := &pr.PostgresProvider{DB: db, TableName: tableName}
	broker := &memory.Memory{}
	testService = &Service{provider: provider, broker: broker}

	testToken = utilities.LoginUser(db)

	t.Run("add a post handler", testAddPostHandler)
	t.Run("get a post handler", testGetPostHandler)
	t.Run("lists posts handler", testListPostsHandler)
	t.Run("update posts handler", testUpdatePostHandler)
	t.Run("delete posts handler", testDeletePostHandler)

	t.Run("post stream client", testSubscribeClientStream)
}
