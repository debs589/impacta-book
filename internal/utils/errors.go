package utils

import "errors"

var (
	ErrNotFound         = errors.New("entity not found")            // 404
	ErrInvalidArguments = errors.New("empty or invalid arguments ") // 400
	ErrForbidden        = errors.New("action forbidden")            // 403
)
