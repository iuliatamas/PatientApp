package common

import "time"

type ActionType int

const (
	Prescription ActionType = iota
	Symptom
	Sentiment
)

type Action interface {
	Type() ActionType       // type of action
	Time() time.Time        // time when action must be taken
	Timeout() time.Duration // time willing to wait for response
	Tries() int             // number of times we've contacted patient with this action

	OnAnswer(s string)
	OnNoAnswer()

	Canceled() bool // when prescription is updated
}

// types that satisfy the action intrface
type PrescriptionAction struct {
	Type    string
	Time    time.Time
	Timeout time.Duration
	Tries   int
}

func (pa *PrescriptionAction) Type() {
	return Type
}

func (pa *PrescriptionAction) Time() {
	return Time
}

func (pa *PrescriptionAction) Timeout() {
	return Timeout
}

func (pa *PrescriptionAction) OnAnswer(s string) {

}

func (pa *PrescriptionAction) OnNoAnswer(s string) {
	// Add same action after timeout, or contact clinician
	if pa.Tries+1 <= MAX_TRIES {
		// add to action queue again
	} else {
		// contact physicians
	}
}
