package post

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/idirall22/utilities"
)

// AddPostHandler add a post
func (s *Service) AddPostHandler(w http.ResponseWriter, r *http.Request) {

	form := PForm{}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	post, err := s.addPost(ctx, form)

	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

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

	groupID, err := utilities.GetURLID(r, "groupID")
	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	limit, offset := utilities.GetParamsURLLimitAndOffset(r, DefaultLimit, "", "")

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	post, err := s.listPosts(ctx, groupID, offset, limit)

	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

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

	id, err := utilities.GetURLID(r, "")
	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	groupID, err := utilities.GetURLID(r, "groupID")
	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	post, err := s.getPost(ctx, id, groupID)

	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

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

	id, err := utilities.GetURLID(r, "")
	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	form := PForm{}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	post, err := s.updatePost(ctx, id, form)

	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

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

	id, err := utilities.GetURLID(r, "")
	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	err = s.deletePost(ctx, id)

	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Header().Add("Content-Type", "application/json")
}
