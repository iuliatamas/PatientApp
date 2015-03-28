package main

type Person interface {
	Name() string
	Phone() string
	Email() string
}

type PersonInfo struct {
	name  string
	phone string
	email string
}

type Patient struct {
	PersonInfo
	Conditions []Condition
	Contacts   []Person
	Treatments []Treatment
	Symptoms   map[string]SymptomStats
	DrugUsage  map[string]TreatmentStats

	// actions we have taken for the patient
	Actions []Action
	// limited to one clinician for now
	Clinician *Clinician
}

func (p *Patient) Name() string {
	return p.name
}
func (p *Patient) Phone() string {
	return p.phone
}
func (p *Patient) Email() string {
	return p.email
}

type Clinician struct {
	PersonInfo
}

func (p *Clinician) Name() string {
	return p.name
}
func (p *Clinician) Phone() string {
	return p.phone
}
func (p *Clinician) Email() string {
	return p.email
}
