package models

import (
	"errors"
	"time"
)

var ErrSnippetNoRow = errors.New("Data tidak ada di database")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
