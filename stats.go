package main

type SymptomReport struct {
	Reported bool
	Severity int32
}

type SymptomStats struct {
	Symptom SymptomData
	Reports []SymptomReport
	// SymptomReport can be aggregated into
	// Frequency
	// AvgFrequency
	// StdDevFrequency
}

type TreatmentReport struct {
	Missed bool
}

type TreatmentStats struct {
	Treatment Treatment
	Reports   []TreatmentReport
	// Frequency
	// AvgFrequency
	// StdDevFrequency
}
