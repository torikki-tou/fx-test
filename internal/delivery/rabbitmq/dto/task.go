package dto

import "encoding/json"

type TaskEvent struct {
	Type string
	Data struct {
		SomeFields json.RawMessage
	}
}
