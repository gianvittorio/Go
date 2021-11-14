package hydrablayer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoDataStore struct {
	*mgo.Session
}

func NewMongoStore(conn string) (*mongoDataStore, error) {
	session, err := mgo.Dial(conn)
	if err != nil {
		return nil, err
	}
	return &mongoDataStore{Session: session}, nil
}

func (ms *mongoDataStore) AddMember(cm *CrewMember) error {
	session := ms.Copy()
	defer session.Close()
	personnel := session.DB("admin").C("personnel")
	return personnel.Insert(cm)
}

func (ms *mongoDataStore) FindMember(id int) (CrewMember, error) {
	session := ms.Copy()
	defer session.Close()
	personnel := session.DB("admin").C("personnel")
	cm := CrewMember{}
	err := personnel.Find(bson.M{"id": id}).One(&cm)
	return cm, err
}

func (ms *mongoDataStore) AllMembers() (crew, error) {
	session := ms.Copy()
	defer session.Close()
	personnel := session.DB("admin").C("personnel")
	members := crew{}
	err := personnel.Find(nil).All(&members)
	return members, err
}
