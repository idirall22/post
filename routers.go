package post

import (
	"github.com/gorilla/mux"
	u "github.com/idirall22/user"
)

// Router comment endpoints
func (s *Service) Router() *mux.Router {
	r := &mux.Router{}

	r.HandleFunc("/posts", u.AuthnticateUser(s.ListPostsHandler)).Methods("GET")
	r.HandleFunc("/posts/{id}", u.AuthnticateUser(s.GetPostHandler)).Methods("GET")
	r.HandleFunc("/posts", u.AuthnticateUser(s.AddPostHandler)).Methods("POST")
	r.HandleFunc("/posts/{id}", u.AuthnticateUser(s.UpdatePostHandler)).Methods("PUT")
	r.HandleFunc("/posts/{id}", u.AuthnticateUser(s.DeletePostHandler)).Methods("DELETE")
	r.HandleFunc("/posts/stream", u.AuthnticateUser(s.SubscribeClientStream))

	return r
}
