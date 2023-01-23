package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"
)

type StatusType int

func NewStatusType(sType any) (*StatusType, error) {
	t := reflect.TypeOf(sType).String()
	switch t {
	case "string":
		str := sType.(string)
		val, ok := toEnum[str]
		if !ok {
			return nil, fmt.Errorf("invalid status key. valid keys are NOT_STARTED, IN_PROGRESS, COMPLETED. your input %v", sType)
		}
		return &val, nil
	case "int":
		i := sType.(int)
		st := StatusType(i)
		_, ok := toString[st]
		if !ok {
			return nil, fmt.Errorf("invalid status value. valid keys are 0, 1, 2. your input %v", sType)
		}
		return &st, nil
	default:
		return nil, fmt.Errorf("invalid type")
	}

}

const (
	NotStarted StatusType = iota
	InProgress
	Completed
)

var toString = map[StatusType]string{
	NotStarted: "NOT_STARTED",
	InProgress: "IN_PROGRESS",
	Completed:  "COMPLETED",
}

var toEnum = map[string]StatusType{
	"NOT_STARTED": NotStarted,
	"IN_PROGRESS": InProgress,
	"COMPLETED":   Completed,
}

func (s *StatusType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	val, ok := toString[*s]
	if !ok {
		return nil, errors.New("unsupported status enum, supported values: 0, 1, 2")
	}
	buffer.WriteString(val)
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (g *StatusType) UnmarshalJSON(b []byte) error {
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
	Status      StatusType `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
