package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewErrorResponse(root error, message string, log string, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

// WrapError wraps an existing error into an AppError with a new message and log.

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func NewFullErrorResponse(statusCode int, root error, message string, log string, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewUnAuthorizedErrorResponse(root error, message string, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    message,
		Key:        key,
	}
}

func NewCustomError(root error, message string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, message, root.Error(), key)
	}

	return NewErrorResponse(errors.New(message), message, message, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong", err.Error(), "INTERNAL_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "INVALID_REQUEST")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity))
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity))
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity))
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity))
}

func ErrEntityDeleted(entity string) *AppError {
	return NewCustomError(
		nil,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity))
}

func ErrEntityExisted(entity string) *AppError {
	return NewCustomError(
		nil,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("Err%sExists", entity))
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity))
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity))
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		nil,
		fmt.Sprintf("You have no permission"),
		fmt.Sprintf("ErrNoPermission"))
}

var RecordNotFound = errors.New("record not found")
