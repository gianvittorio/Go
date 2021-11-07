package hydraconfigurator

import (
	"encoding/xml"
	"fmt"
	"os"
)

func decodeXMLConfig(v interface{}, filename string) error {
	fmt.Println("Decoding XML")
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	return xml.NewDecoder(file).Decode(v)
}