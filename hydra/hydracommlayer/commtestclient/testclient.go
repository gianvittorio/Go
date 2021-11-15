package main

import (
	"flag"
	"hydra/hydra/hydracommlayer"
	"hydra/hydra/hydracommlayer/hydraproto"
	"log"
	"strings"
)

func main() {
	op := flag.String("type", "", "Server (s) or client (c) ?")
	address := flag.String("addr", ":8080", "address? host:port")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runServer(*address)
	case "C":
		runClient(*address)
	}
}

func runServer(dest string) {
	c := hydracommlayer.NewConnection(hydracommlayer.Protobuf)
	recvChan, err := c.ListenAndDecode(dest)
	if err != nil {
		log.Fatal(err)
	}
	for msg := range recvChan {
		log.Println("Received: ", msg)
	}
}

func runClient(dest string) {
	c := hydracommlayer.NewConnection(hydracommlayer.Protobuf)
	ship := &hydraproto.Ship{
		Shipname: "Hydra",
		CaptainName: "Jala",
		Crew: []*hydraproto.Ship_CrewMember{
			{Id: 1, Name: "Kevin", SecClearance: 5, Position: "Pilot"},
			{Id: 2, Name: "Jade", SecClearance: 4, Position: "Tech"},
			{Id: 3, Name: "Wally", SecClearance: 5, Position: "Engineer"},
		},
	}

	if err := c.EncodeAndSend(ship, dest); err != nil {
		log.Println("Error occurred while sending message", err)
	} else {
		log.Println("Send operation succeeded")
	}
}
