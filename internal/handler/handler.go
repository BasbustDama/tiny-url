package handler

import (
	"net/http"
	"regexp"

	"github.com/BasbustDama/tiny-url/internal/errors"
)

const (
	methodNotAllowedMsg  = "method not allowed"
	contentTypeHeaderMsg = "Content-Type header is not \"application/json\""
	jsonParseMsg         = "parse body error"
	emptyUrlMsg          = "empty url"
	invalidUrlMsg        = "Invalid URL param"
)

const (
	contentType     = "Content-Type"
	jsonContentType = "application/json"
)

type shortenerUsecase interface {
	Create(url string) (string, error)
	Get(url string) (string, error)
}

func NewHandler(usecase shortenerUsecase) http.Handler {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/", redirectHandler(usecase))
	serverMux.HandleFunc("/add", registerHandler(usecase))

	return serverMux
}

func registerGlobalError(w http.ResponseWriter, err error) {
	if appError, ok := err.(errors.AppError); ok {
		switch appError {
		case errors.ErrorBadRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.ErrorNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		case errors.ErrorInternal:
			fallthrough
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
}

const urlRegexpPattern = `^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$`

func validateUrl(url string) bool {
	re := regexp.MustCompile(urlRegexpPattern)
	return re.MatchString(url)
}
