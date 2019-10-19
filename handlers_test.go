package post

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	u "github.com/idirall22/user"
)

// Test AddPostHandler
func testAddPostHandler(t *testing.T) {

	form := PForm{
		Content: "post content",
		UserID:  1,
		GroupID: 1,
	}
	b, err := json.Marshal(form)

	if err != nil {
		t.Error(err)
		return
	}
	body := bytes.NewReader(b)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/post", body)

	if err != nil {
		t.Error(err)
		return
	}

	r.Header.Add("Authorization", testToken)
	h := http.HandlerFunc(u.AuthnticateUser(testService.AddPostHandler))

	h.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Status code should be %d but got %d", http.StatusCreated, w.Code)
	}
}

// Test GetPostHandler
func testGetPostHandler(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/post/1/1", nil)

	if err != nil {
		t.Error(err)
		return
	}

	r.Header.Add("Authorization", testToken)

	router := mux.NewRouter()
	router.HandleFunc("/post/{groupID}/{id}", u.AuthnticateUser(testService.GetPostHandler))
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %d but got %d", http.StatusOK, w.Code)
	}
}

// Test ListPostsHandler
func testListPostsHandler(t *testing.T) {

	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/posts/1?offset=2&limit=0", nil)

	if err != nil {
		t.Error(err)
		return
	}

	r.Header.Add("Authorization", testToken)

	router := mux.NewRouter()
	router.HandleFunc("/posts/{groupID}", u.AuthnticateUser(testService.ListPostsHandler))
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %d but got %d", http.StatusOK, w.Code)
	}
}

// Test UpdatePostHandler
func testUpdatePostHandler(t *testing.T) {

	form := PForm{
		Content: "post content updated",
		UserID:  1,
		GroupID: 1,
	}
	b, err := json.Marshal(form)

	if err != nil {
		t.Error(err)
		return
	}
	body := bytes.NewReader(b)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("PUT", "/post/1", body)

	if err != nil {
		t.Error(err)
		return
	}

	r.Header.Add("Authorization", testToken)

	router := mux.NewRouter()
	router.HandleFunc("/post/{id}", u.AuthnticateUser(testService.UpdatePostHandler))
	router.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Status code should be %d but got %d", http.StatusOK, w.Code)
	}
}

// Test DeletePostHandler
func testDeletePostHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("DELETE", "/post/1", nil)

	if err != nil {
		t.Error(err)
		return
	}

	r.Header.Add("Authorization", testToken)

	router := mux.NewRouter()
	router.HandleFunc("/post/{id}", u.AuthnticateUser(testService.DeletePostHandler))
	router.ServeHTTP(w, r)

	if w.Code != http.StatusNoContent {
		t.Errorf("Status code should be %d but got %d", http.StatusNoContent, w.Code)
	}
}
