package post

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AddPostHandler add a post
func (s *Service) AddPostHandler(w http.ResponseWriter, r *http.Request) {

	form := PForm{}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	post, err := s.addPost(ctx, form)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")

	if json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// ListPostsHandler list posts
func (s *Service) ListPostsHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	groupID, err := parseID(params["groupID"])
	if err != nil {
		return
	}

	offset, limit := getOffsetAndLimit(r)

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	post, err := s.listPosts(ctx, groupID, offset, limit)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	if json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// GetPostHandler get a post
func (s *Service) GetPostHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := parseID(params["id"])
	if err != nil {
		return
	}

	groupID, err := parseID(params["groupID"])
	if err != nil {
		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	post, err := s.getPost(ctx, id, groupID)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	if json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// UpdatePostHandler update a post
func (s *Service) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := parseID(params["id"])
	if err != nil {
		return
	}

	form := PForm{}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	post, err := s.updatePost(ctx, id, form)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	if json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// DeletePostHandler delete a post
func (s *Service) DeletePostHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := parseID(params["id"])
	if err != nil {
		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	err = s.deletePost(ctx, id)

	if err != nil {
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Header().Add("Content-Type", "application/json")
}
