package post

import "errors"

var (
	// ErrorForm when post form is not valid
	ErrorForm = errors.New("post form not valid")

	// ErrorID when id is < 1
	ErrorID = errors.New("Post ID not exists")
)
