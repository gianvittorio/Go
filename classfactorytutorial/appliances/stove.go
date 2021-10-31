package appliances

type stove struct {
	typeName string
}

func (fr *stove) Start() {
	fr.typeName = " Stove"
}

func (fr *stove) GetPurpose() string {
	return "I am a " + fr.typeName + " I cook food!!!"
}
