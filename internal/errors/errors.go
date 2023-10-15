package errors

import "errors"

type AppError error

var (
	ErrorInternal   AppError = errors.New("internal server error")
	ErrorBadRequest AppError = errors.New("bad request error")
	ErrorNotFound   AppError = errors.New("not found error")
)
