package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

type statusType int

const (
	NotStarted statusType = iota
	InProgress
	Completed
)

var toString = map[statusType]string{
	NotStarted: "NOT_STARTED",
	InProgress: "IN_PROGRESS",
	Completed:  "COMPLETED",
}

var toEnum = map[string]statusType{
	"NOT_STARTED": NotStarted,
	"IN_PROGRESS": InProgress,
	"COMPLETED":   Completed,
}

func (s *statusType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	val, ok := toString[*s]
	if !ok {
		return nil, errors.New("unsupported status enum, supported values: 0, 1, 2")
	}
	buffer.WriteString(val)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (g *statusType) UnmarshalJSON(b []byte) error {
	var j string
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	r, ok := toEnum[j]
	if !ok {
		return errors.New("unsupported status type, supported values: NOT_STARTED, IN_PROGRESS, COMPLETED")
	}
	*g = r
	return nil
}

type Todo struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      statusType `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
