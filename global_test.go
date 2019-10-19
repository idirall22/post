package post

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	pr "github.com/idirall22/post/providers/postgres"
	u "github.com/idirall22/user"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "diskshar_test"
)

var testService *Service

var testToken string
var userUsernameTest = "alice"
var userPasswordTest = "fdpjfd654/*sMLdf"

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

// Get user auth allows to login and get a test token
func getUserAuth(db *sql.DB) {

	m := make(map[string]string)
	m["username"] = userUsernameTest
	m["password"] = userPasswordTest

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	body := bytes.NewReader(b)

	serviceUser := u.StartService(db, "users")

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/login", body)

	if err != nil {
		fmt.Println(err)
		return
	}

	h := http.HandlerFunc(serviceUser.Login)
	h.ServeHTTP(w, r)

	testToken = w.Header().Get("Autherization")
}

// Connect to db test
func connectDB() error {

	dbInfos := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbInfos)
	if err != nil {
		return err
	}

	provider := &pr.PostgresProvider{DB: db, TableName: "posts"}
	testService = &Service{
		provider: provider,
	}
	err = cleanDB(db)
	if err != nil {
		return err
	}

	getUserAuth(db)

	testService = NewService(db, "posts")
	return nil
}

func TestGlobal(t *testing.T) {
	if err := connectDB(); err != nil {
		log.Fatal("Error connect database test, ", err)
	}

	defer closeDB(testService.provider.(*pr.PostgresProvider).DB)

	t.Run("add a post handler", testAddPostHandler)
	t.Run("get a post handler", testGetPostHandler)
	t.Run("lists posts handler", testListPostsHandler)
	t.Run("update posts handler", testUpdatePostHandler)
	t.Run("delete posts handler", testDeletePostHandler)

}
