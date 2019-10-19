package post

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
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

	h := http.HandlerFunc(testService.AddPostHandler)

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
	// mux.SetURLVars(r, map[string]string{"id": "1", "groupID": "1"})

	router := mux.NewRouter()
	router.HandleFunc("/post/{groupID}/{id}", testService.GetPostHandler)
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

	router := mux.NewRouter()
	router.HandleFunc("/posts/{groupID}", testService.ListPostsHandler)
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

	router := mux.NewRouter()
	router.HandleFunc("/post/{id}", testService.UpdatePostHandler)
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

	router := mux.NewRouter()
	router.HandleFunc("/post/{id}", testService.DeletePostHandler)
	router.ServeHTTP(w, r)

	if w.Code != http.StatusNoContent {
		t.Errorf("Status code should be %d but got %d", http.StatusNoContent, w.Code)
	}
}
