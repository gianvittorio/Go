package main

import (
	_ "encoding/csv"
	_ "io"
	"log"
	_ "os"
	_ "strconv"
	"sync"
	_ "sync"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type crewMember struct {
	ID int `bson: "id"`
	Name string `bson: "name"`
	SecClearance int `bson: "security_clearance"`
	Position string `bson: "position"`
}

type Crew []crewMember

func main() {
	session, err := mgo.Dial("mongo1:9042?connect=direct")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	
	personnel := session.DB("admin").C("personnel")

	n, _ := personnel.Count()
	log.Println("Number of personnel is ", n)

	cm := crewMember{}
	personnel.Find(bson.M{"position": "Technician"}).One(&cm)
	log.Println(cm)

	query := bson.M{
		"security_clearance": bson.M{
			"$gt": 3,
		},
		"position": bson.M{
			"$in": []string{"Mechanic", "Biologist"},
		},
	}

	var crew Crew
	log.Println(query)
	err = personnel.Find(query).All(&crew)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Query results: ", crew)

	names := []struct {
		Name string `bson:"name"`
	}{}

	err = personnel.Find(query).Select(bson.M{"name": 1}).All(&names)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(names)


	// Insert
	newcr := crewMember{ID: 18, Name: "Kaya Gal", SecClearance: 4, Position: "Biologist"}
	if err = personnel.Insert(newcr); err != nil {
		log.Fatal(err)
	}

	// Update
	err = personnel.Update(bson.M{"id": 18}, bson.M{"$set": bson.M{"position": "Engineer iV"}})
	if err != nil {
		log.Fatal(err)
	}

	// Remove
	if err = personnel.Remove(bson.M{"id": 18}); err != nil {
		log.Fatal(err)
	}

	// Concurrent access
	var wg sync.WaitGroup
	count, _ := personnel.Count()
	wg.Add(count)
	for i := 1; i <= count; i++ {
		go readId(i, session.Copy(), &wg)
	}
	wg.Wait()
}

func readId(id int, sessionCopy *mgo.Session, wg *sync.WaitGroup) {
	defer func() {
		sessionCopy.Close()
		wg.Done()
	}()
	p := sessionCopy.DB("admin").C("personnel")
	cm := crewMember{}
	err := p.Find(bson.M{"id": id}).One(&cm)
	if err != nil {
		log.Printf("Could not retrieve id %d, error %s \n", id, err.Error())
		return
	}
	log.Println(cm)
}