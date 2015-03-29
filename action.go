package main

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
	String() string
}

// types that satisfy the action intrface
type PrescriptionAction struct {
	time    time.Time
	timeout time.Duration
	tries   int
	str     string

	patient *Patient
}

func (pa *PrescriptionAction) String() string {
	return pa.str
}

func (pa *PrescriptionAction) Patient() *Patient {
	return pa.patient
}

func (pa *PrescriptionAction) Type() ActionType {
	return Prescription
}

func (pa *PrescriptionAction) Time() time.Time {
	return pa.time
}

func (pa *PrescriptionAction) Timeout() time.Duration {
	return pa.timeout
}

func (pa *PrescriptionAction) OnAnswer(answ string, s *Server) {

}

var PRESCRIPTION_ALERT = "Patient not taking medicine!"
var MAX_TRIES int = 5

func (pa *PrescriptionAction) OnNoAnswer(s *Server) {
	// Add same action after timeout, or contact clinician
	if pa.tries+1 <= MAX_TRIES {
		// add to action queue again

	} else {
		// contact physicians
		s.SendSMS(pa.Patient().Clinician, PRESCRIPTION_ALERT)
	}
}
