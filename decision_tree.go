package main

// Yes, No, Idk

// each level of a decision tree has a
type DecisionTree struct {
	Server   Server
	Action   Action
	Decision map[YesNoIdk]*DecisionTree
}

func (d *DecisionTree) Do(s *Server, resp string) *DecisionTree {
	yni := IsYesNoIdk(resp)
	dnext := d.Decision[yni]
	if dnext == nil {
		return dnext
	}
	msg := dnext.Action.String()
	p := dnext.Action.Patient()
	s.SendSMS(p, msg)
	return dnext
}
