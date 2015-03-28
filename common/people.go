package PatientApp

type Person interface {
	Name()
	Phone()
	Email()
}

type PersonInfo struct {
	Name  string
	Phone string
	Email string
}

type Patient struct {
	PersonInfo
	Conditions []Condition
	Contacts   []Person
	Drugs      []Drug
	Symptoms   map[string]SymptomStats
	DrugUsage  map[string]DrugStats

	// actions we have taken for the patient
	Actions []Action
	// limited to one clinician for now
	Clinician Clinician
}

func (p *Patient) Name() string {
	return p.Name
}
func (p *Patient) Phone() string {
	return p.Phone()
}
func (p *Patient) Email() string {
	return p.Email
}

type Clinician struct {
	PersonInfo
}

func (p *Clinician) Name() string {
	return p.Name
}
func (p *Clinician) Phone() string {
	return p.Phone()
}
func (p *Clinician) Email() string {
	return p.Email
}
