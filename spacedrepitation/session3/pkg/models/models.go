package models

import (
	"errors"
	"time"
)

var errNoRecord = errors.New("error terjadi saja")

type Snippet struct {
	ID      int
	title   string
	content string
	created time.Time
	expires time.Time
}
