package errs

//go:generate easyjson

import (
	"github.com/mailru/easyjson"
)

const (
	Internal         = 0
	BadFormat        = 1
	InvalidFormat    = 2
	Conflict         = 3
	NotFound         = 4
	NotAuthorized    = 5
	PermissionDenied = 6
)

//easyjson:json
type Error struct {
	Type    int    `json:"-"`
	Message string `json:"message"`
	json    []byte `json:"-"`
}

func NewError(t int, message string) *Error {
	err := &Error{
		Type:    t,
		Message: message,
	}
	err.json, _ = easyjson.Marshal(err)
	return err
}

func New(err error) *Error {
	return NewError(Internal, err.Error())
}

func NewInternalError(message string) *Error {
	return NewError(Internal, message)
}

func NewBadFormatError(message string) *Error {
	return NewError(BadFormat, message)
}

func NewInvalidFormatError(message string) *Error {
	return NewError(InvalidFormat, message)
}

func NewConflictError(message string) *Error {
	return NewError(Conflict, message)
}

func NewNotFoundError(message string) *Error {
	return NewError(NotFound, message)
}

func NewNotAuthorizedError(message string) *Error {
	return NewError(NotAuthorized, message)
}

func NewPermissionDeniedError(message string) *Error {
	return NewError(PermissionDenied, message)
}

func (v *Error) Error() string {
	return v.Message
}

func (v *Error) JSON() []byte {
	return v.json
}
