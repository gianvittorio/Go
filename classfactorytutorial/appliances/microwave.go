package appliances

type microwave struct {
	typeName string
}

func (fr *microwave) Start() {
	fr.typeName = " microwave"
}

func (fr *microwave) GetPurpose() string {
	return "I am a " + fr.typeName + " I warm stuff up!!!"
}
