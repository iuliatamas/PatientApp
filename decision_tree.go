package main

import "time"

// Yes, No, Idk

// each level of a decision tree has a
type DecisionTree struct {
	Action   Action
	Decision map[YesNoIdk]*DecisionTree
}

func (d *DecisionTree) Do(s *Server, resp string) *DecisionTree {
	yni := IsYesNoIdk(resp)
	d.Action.Done()
	dnext := d.Decision[yni]
	// if there are no more decisions left
	// return nil, the caller should queue the next pending decision tree
	// of the patient
	if dnext == nil {
		return nil
	}
	msg := dnext.Action.String()
	p := dnext.Action.Patient()
	s.SendSMS(p, msg)

	// setup an action in the server to happen if the action is not completed
	// in the timeout time
	t := dnext.Action.Timeout()

	go func() {
		select {
		case <-time.After(t):
			dnext.Action.OnNoAnswer(s)
		}
	}()
	return dnext
}

func NewDT(a Action) *DecisionTree {
	return &DecisionTree{Action: a, Decision: make(map[YesNoIdk]*DecisionTree, 3)}
}
