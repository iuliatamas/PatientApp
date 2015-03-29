package main

import "sync"

type Server struct {
	PhoneNumber string

	// queue of actions
	ActLock sync.Mutex
	Actions ActionQueue
}

func NewServer(pn string) *Server {
	return &Server{PhoneNumber: pn, Actions: make([]Action, 0)}
}

func (s *Server) ActOnResponse(p *Patient, resp string) {
	dt := p.DecisionTree
	p.DecisionTree = dt.Do(s, resp)
}
