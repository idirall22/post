package post

import "time"

var (
	// TimeoutRequest time to stop a request
	TimeoutRequest = time.Second * 5

	// DefaultLimit limit used in db
	DefaultLimit = 5

	// MaxLimit limit used in db
	MaxLimit = 10
)
