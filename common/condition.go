package main

type Condition struct {
	Id   int
	Name string

	SympProb map[Symptom]float32
}

type Symptom struct {
	Id   int
	Name string

	// questions that can help the patient identify if they have the symptom
	// and their responses
	Questions [string]Response
}

type Response struct {
	// message entered by patient: Yes, No, ...
	m string
	// Phone number as identifier
	from string
}
