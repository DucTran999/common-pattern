package workerpool

import (
	"fmt"
	"time"
)

type Job struct {
	ID     int64
	record []string
}

func NewJob(id int64, record []string) *Job {
	return &Job{
		ID:     id,
		record: record,
	}
}

func (j *Job) Run() error {
	// Simulate process a record need 200ms
	time.Sleep(time.Millisecond * 200)

	if len(j.record) < 6 {
		err := fmt.Errorf(
			"record %d: insufficient columns (expected 6, got %d)",
			j.ID, len(j.record),
		)
		return err
	}

	return nil
}
