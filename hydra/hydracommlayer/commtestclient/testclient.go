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
			&hydraproto.Ship_CrewMember{1, "Kevin", 5, "Pilot"},
			&hydraproto.Ship_CrewMember{2, "Jade", 4, "Tech"},
			&hydraproto.Ship_CrewMember{3, "Wally", 5, "Engineer"},
		},
	}

	if err := c.EncodeAndSend(ship, dest); err != nil {
		log.Println("Error occurred while sending message", err)
	}
}
