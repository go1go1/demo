package session

import "errors"

var (
	ErrSessionNotExist      = errors.New("session not exists")
	ErrKeyNotExistInSession = errors.New("session key not exists")
)
