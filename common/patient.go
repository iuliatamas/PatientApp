package PatientApp

type Person struct {
	Name  string
	Phone string
	Email string
}

type Patient struct {
	Person
	Conditions []Condition
	Contacts   []Person
	Drugs      []Drug
	Symptoms   map[string]SymptomStats
	DrugUsage  map[string]DrugStats
}
