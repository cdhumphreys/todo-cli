package csv

import (
	"strconv"
	"time"
)

const TimeFormat = time.RFC3339

type Todo struct {
	ID          int
	Description string
	CreatedAt   string
	Completed   bool
}

func (t Todo) String() []string {
	return []string{strconv.Itoa(t.ID), t.Description, t.CreatedAt, strconv.FormatBool(t.Completed)}
}
