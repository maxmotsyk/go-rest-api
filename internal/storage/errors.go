package storage

import (
	"errors"
)

var (
	ErrorUrlsNotFound = errors.New("urls not found")
	ErrorURLExists    = errors.New("url exists")
)
