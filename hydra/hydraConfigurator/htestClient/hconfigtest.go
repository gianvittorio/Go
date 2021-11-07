package main

import (
	"fmt"
	hydraconfigurator "hydra/hydra/hydraConfigurator"
)

type ConfStruct struct {
	TS string `name:"testString" xml:"testString" json:"testString"`
	TB bool `name:"testBool" xml:"testBool" json:"testBool"`
	TF float64 `name:"testFloat" xml:"testFloat" json:"testFloat"`
	TestInt int
}

func main() {
	configstruct := new(ConfStruct)
	hydraconfigurator.GetConfiguration(hydraconfigurator.XML, configstruct, "configfile.xml")
	fmt.Println(*configstruct)

	if configstruct.TB {
		fmt.Println("bool is true")
	}

	fmt.Println(float64(4.8 * configstruct.TF))

	fmt.Println(5 * configstruct.TestInt)

	fmt.Println(configstruct.TS)
}
