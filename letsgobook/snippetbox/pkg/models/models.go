package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("Model: is not matching found")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
