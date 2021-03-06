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
	OnNoAnswer(s *Server)

	Patient() *Patient // patient we are acting for
	String() string
	Done()
}

// types that satisfy the action intrface
type PrescriptionAction struct {
	time    time.Time
	timeout time.Duration
	tries   int
	str     string

	patient *Patient
	done    bool
}

func NewPA(msg string, p *Patient) *PrescriptionAction {
	return &PrescriptionAction{
		time.Now(),
		1 * time.Minute,
		3,
		msg,
		p,
		false,
	}
}

func (pa *PrescriptionAction) String() string {
	return pa.str
}

func (pa *PrescriptionAction) Done() {
	pa.done = true
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

func (pa *PrescriptionAction) Tries() int {
	return pa.tries
}

func (pa *PrescriptionAction) Timeout() time.Duration {
	return pa.timeout
}

func (pa *PrescriptionAction) OnAnswer(answ string, s *Server) {

}

var PRESCRIPTION_ALERT = "Patient not taking medicine!"
var MAX_TRIES int = 5

// XXX: should return time to reschedule it
func (pa *PrescriptionAction) OnNoAnswer(s *Server) {
	if pa.done {
		return
	}
	// Add same action after timeout, or contact clinician
	if pa.tries+1 <= MAX_TRIES {
		// add to action queue again

	} else {
		// contact physicians
		s.SendSMS(pa.Patient().Clinician, PRESCRIPTION_ALERT)
	}
}
