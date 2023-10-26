package entity

import (
	"time"
)

type Todo struct {
	ID           string
	Title        string
	Description  string
	Status       string
	DateFinished time.Time
}

// func (t *Todo) Done() {
// 	t.Status = "done"
// 	t.DateFinished = time.Now()
// }

// func (t *Todo) Undone() {
// 	t.Status = "undone"
// 	t.DateFinished = time.Now()
// }
