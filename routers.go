package post

import (
	"github.com/gorilla/mux"
	u "github.com/idirall22/user"
)

// Router comment endpoints
func (s *Service) Router(r *mux.Router) {

	sr := r.PathPrefix("/posts").Subrouter()

	sr.HandleFunc("/", u.AuthnticateUser(s.ListPostsHandler)).Methods("GET")
	sr.HandleFunc("/{id}", u.AuthnticateUser(s.GetPostHandler)).Methods("GET")
	sr.HandleFunc("/", u.AuthnticateUser(s.AddPostHandler)).Methods("POST")
	sr.HandleFunc("/{id}", u.AuthnticateUser(s.UpdatePostHandler)).Methods("PUT")
	sr.HandleFunc("/{id}", u.AuthnticateUser(s.DeletePostHandler)).Methods("DELETE")
	sr.HandleFunc("/stream", u.AuthnticateUser(s.SubscribeClientStream))
}
