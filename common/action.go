package common

import "time"

type ActionType int

const (
	Prescription ActionType = iota
	Symptom
	Sentiment
)

type Action interface {
	Type() ActionType
	Time() time.Time

	OnAnswer(s string)
	OnNoAnswer()
	Timeout() time.Duration
}
