package appliances

type fridge struct {
	typeName string
}

func (fr *fridge) Start() {
	fr.typeName = " Fridge"
}

func (fr *fridge) GetPurpose() string {
	return "I am a " + fr.typeName + " I cool stuff down!!!"
}
