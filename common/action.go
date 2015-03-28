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

	OnAnswer(answ string, s *Server)
	OnNoAnswer()

	Canceled() bool // when prescription is updated

	Patient() *Patient // patient we are acting for
}

// types that satisfy the action intrface
type PrescriptionAction struct {
	Time    time.Time
	Timeout time.Duration
	Tries   int
}

func (pa *PrescriptionAction) Type() ActionType {
	return Prescription
}

func (pa *PrescriptionAction) Time() {
	return Time
}

func (pa *PrescriptionAction) Timeout() {
	return Timeout
}

func (pa *PrescriptionAction) OnAnswer(answ string, s *Server) {

}

var PRESCRIPTION_ALERT = "Patient not taking medicine!"

func (pa *PrescriptionAction) OnNoAnswer(s *Server) {
	// Add same action after timeout, or contact clinician
	if pa.Tries+1 <= MAX_TRIES {
		// add to action queue again
	} else {
		// contact physicians
		s.sendSMS(pa.Patient().Clinician(), PRESCRIPTION_ALERT)
	}
}
