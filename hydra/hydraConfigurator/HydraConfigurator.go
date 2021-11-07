package hydraconfigurator

import (
	"errors"
	"reflect"
)

const (
	CUSTOM uint8 = iota
	JSON
	XML
)

var wrongTypeError error = errors.New("Type must be pointer to a struct")

func GetConfiguration(confType uint8, obj interface{}, filename string) (err error) {
	mysRValue := reflect.ValueOf(obj)
	if mysRValue.Kind() != reflect.Ptr || mysRValue.IsNil() {
		return wrongTypeError
	}

	mysRValue = mysRValue.Elem()
	if mysRValue.Kind() != reflect.Struct {
		return wrongTypeError
	}

	switch confType {
	case CUSTOM:
		err = MarshalCustomConfig(mysRValue, filename)
	case JSON:
		err = decodeJSONConfig(obj, filename)
	case XML:
		err = decodeXMLConfig(obj, filename)
	}

	return err
}