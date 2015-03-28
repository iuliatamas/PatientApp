package PatientApp

type SymptomReport struct {
	Reported bool
	Severity int32
}

type SymptomStats struct {
	Symptom Symptom
	Reports []SymptomReport
	// SymptomReport can be aggregated into
	// Frequency
	// AvgFrequency
	// StdDevFrequency
}

type DrugReport struct {
	Missed bool
}

type DrugStats struct {
	Drug    Drug
	Reports []DrugReport
	// Frequency
	// AvgFrequency
	// StdDevFrequency
}
