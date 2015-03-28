package main

type Condition struct {
	Id   int
	Name string

	SympProb map[int]float32 // symptom probabilities by symptom id
}

type SymptomData struct {
	Id   int
	Name string

	// questions that can help the patient identify if they have the symptom
	Questions []string
}

type Response struct {
	// message entered by patient: Yes, No, ...
	m string
	// Phone number as identifier
	from string
}

type Treatment struct {
}
