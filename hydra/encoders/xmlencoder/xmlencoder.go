package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type CrewMember struct {
	ID int `xml:"id,omitempty"`
	Name string `xml:"name,attr"`
	SecurityClearance int `xml:"clearancelevel"`
	AccessCodes []string `xml:"accesscodes>code"`
}

type ShipInfo struct {
	XMLName xml.Name `xml:"ship"`
	ShipID int `xml:"ShipDetails>ShipId"`
	ShipClass string `xml:"ShipDetails>ShipClass"`
	Captain CrewMember
}

func main() {
	file, err := os.Create("xmlfile.xml")
	if err != nil {
		log.Fatal("Could not create file", err)
	}
	defer file.Close()

	cm := CrewMember{ID:1,Name:"Jaro", SecurityClearance:10,AccessCodes: []string{"ADA","LAL"}}
	si := ShipInfo{ShipID:1,ShipClass:"Fighter",Captain:cm}

	b, err := xml.MarshalIndent(&si, " ", " ")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(xml.Header, string(b))

	enc := xml.NewEncoder(file)
	enc.Indent(" ", "    ")
	enc.Encode(si)
	if err != nil {
		log.Fatal("Could not encode xml file", err)
	}
}