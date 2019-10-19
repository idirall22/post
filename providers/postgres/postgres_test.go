package provider

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "diskshar_test"
)

var postNum = 5
var provider *PostgresProvider

func cleanDB(db *sql.DB) error {
	query := fmt.Sprintf(`
		DROP TABLE IF EXISTS posts;

		CREATE TABLE IF NOT EXISTS posts(
		    id SERIAL PRIMARY KEY,
            content VARCHAR NOT NULL,
			media_urls text[],
			user_id INTEGER NOT NULL,
			group_id INTEGER NOT NULL,
		    created_at TIMESTAMP with TIME ZONE DEFAULT now(),
		    deleted_at TIMESTAMP DEFAULT NULL
		);
		`)

	_, err := db.Exec(query)

	if err != nil {
		return err
	}
	return nil
}

func closeDB(db *sql.DB) {
	db.Close()
}

func connectDB() error {

	dbInfos := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbInfos)
	if err != nil {
		return err
	}

	provider = &PostgresProvider{DB: db, TableName: "posts"}

	err = cleanDB(db)
	if err != nil {
		return err
	}

	return nil
}

func TestGlobal(t *testing.T) {
	err := connectDB()
	if err != nil {
		t.Error(err)
		return
	}
	defer closeDB(provider.DB)

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
