package provider

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/idirall22/utilities"
	_ "github.com/lib/pq"
)

var (
	provider  *PostgresProvider
	database  *sql.DB
	testToken string
	postNum   = 5
	tableName = "posts"
	query     = fmt.Sprintf(`
	DROP TABLE IF EXISTS %s;

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

	provider = &PostgresProvider{DB: db, TableName: tableName}

	testToken = utilities.LoginUser(db)

	t.Run("New", testNew)
	t.Run("Get", testGet)
	t.Run("List", testList)
	t.Run("update", testUpdate)
	t.Run("delete", testDelete)
}

// Test New
func testNew(t *testing.T) {

	for i := 0; i < postNum; i++ {
		_, err := provider.New(context.Background(),
			fmt.Sprintf("post %v---", i), []string{"image_url", "image_url"}, 1, 1)

		if err != nil {
			t.Error("Error should be nil but got:", err)
		}
	}
}

// Test Get
func testGet(t *testing.T) {

	c, err := provider.Get(context.Background(), 1, 1, 1)

	if err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}
	if c.ID != 1 {
		t.Errorf("ID should be %d but got: %d", 1, c.ID)
	}
}

// Test List
func testList(t *testing.T) {
	posts, err := provider.List(context.Background(), 1, 1, 5, 0)

	if err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}

	if len(posts) != postNum {
		t.Errorf("Error comments slice length should be %d But got %d",
			postNum, len(posts))
	}
}

// Test update
func testUpdate(t *testing.T) {
	media := []string{"url_image.png"}
	_, err := provider.Update(context.Background(), "update post", media, 1, 1)

	if err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}
}

// Test delete
func testDelete(t *testing.T) {
	err := provider.Delete(context.Background(), 1, 1)
	if err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}
}
