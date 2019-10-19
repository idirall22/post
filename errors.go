package post

import (
	"errors"
	"net/http"

	pr "github.com/idirall22/post/providers/postgres"
)

var (
	// ErrorForm when post form is not valid
	ErrorForm = errors.New("post form not valid")

	// ErrorID when id is < 1
	ErrorID = errors.New("Post ID not exists")
)

func parseError(err error) (string, int) {

	message := ""
	code := http.StatusBadRequest

	switch err {

	case ErrorForm:
		message = ErrorForm.Error()
		break

	case ErrorForm:
		message = ErrorForm.Error()
		break

	case pr.ErrorForeignKey:
		message = "Data not valid"

		code = http.StatusConflict
		break

	case pr.ErrorNoRow:
		message = pr.ErrorNoRow.Error()
		code = http.StatusNotFound
		break

	case pr.ErrorUnique:
		message = pr.ErrorUnique.Error()
		code = http.StatusConflict
		break

	case pr.ErrorServer:
		message = pr.ErrorServer.Error()
		code = http.StatusInternalServerError
		break
	}

	return message, code
}
