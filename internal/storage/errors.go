package storage

import (
	"errors"
)

var (
	ErrorUrlsNotFound = errors.New("urls not found")
	ErrorURLExists    = errors.New("url exists")
	ErrorDeleteURL    = errors.New("error delete url")
)
