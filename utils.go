package post

import (
	"net/http"
	"strconv"
)

// check if id is valid
func checkIfIDIsValid(id int64) error {
	if id < 1 {
		return ErrorID
	}
	return nil
}

// parse id string to int64
func parseID(idStr string) (int64, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		return 0, ErrorID
	}
	return id, nil
}

// get offset and limit from url query
func getOffsetAndLimit(r *http.Request) (int, int) {
	offset := 10
	limit := 0

	off, err := strconv.Atoi(r.URL.Query().Get("offset"))

	if err != nil {
		offset = DefaultOffset
	} else {
		if off <= 0 {
			offset = DefaultOffset
		} else if off > MaxOffset {
			offset = MaxOffset
		}
	}

	lim, err := strconv.Atoi(r.URL.Query().Get("limit"))

	if err != nil {
		limit = 0
	} else {
		if lim <= 0 {
			limit = 0
		}
	}
	return offset, limit
}
