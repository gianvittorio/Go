package appliances

import "errors"

type appliance interface {
	Start()
	GetPurpose() string
}

const (
	STOVE = iota
	FRIDGE
	MICROWAVE
)

func CreateAppliance(myType int) (appliance, error) {
	switch myType {
	case STOVE:
		return new(stove), nil
	case FRIDGE:
		return new(fridge), nil
	case MICROWAVE:
		return new(microwave), nil
	default:
		return nil, errors.New("Invalid Appliance Type")
	}
}