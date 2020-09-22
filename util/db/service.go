package db

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
)

const (
	host     = "localhost"
	port     = 27017
	password = ""     //DB password
	user     = "user" //user name of DB
)

var sessions map[string]*mgo.Session

// GetSession - get a db session handle
func GetSession(db string) *mgo.Session {
	initSessions()
	uri := getConnectionURI(db)
	if sessions[db] == nil {
		mgoSession, err := dail(uri)
		if err != nil {
			panic(err) // no, not really
		}
		mgoSession.SetSafe(&mgo.Safe{WMode: "majority"})
		sessions[db] = mgoSession
	}
	return sessions[db].Clone()
}

func initSessions() {
	if sessions == nil {
		sessions = make(map[string]*mgo.Session)
	}
}

//trying to get connection by 5 times.
func dail(uri string) (*mgo.Session, error) {
	var err error
	var mgoSession *mgo.Session

	for index := 1; index < 5; index++ {
		mgoSession, err = mgo.Dial(uri)
		if err == nil {
			return mgoSession, err
		}
	}
	return mgoSession, err
}

func getConnectionURI(db string) string {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", user, password, host, port, db)
	// fmt.Println("Mongo connection uri", uri)
	return uri
}

// Create -- Insert a doc and return document with if any error.
func Create(db string, collection string, data map[string]interface{}) (map[string]interface{}, error) {
	data["_id"] = GenerateID()
	session := GetSession(db)
	defer session.Close()
	dbCollection := session.DB(db).C(collection)
	err := dbCollection.Insert(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//Update --Update the documents
func Update(db string, collection string, query, upd map[string]interface{}) error {
	session := GetSession(db)
	defer session.Close()
	prog := session.DB(db).C(collection)
	if upd["$set"] == nil {
		upd["$set"] = upd
	}
	err := prog.Update(query, upd)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

//Delete - Delete documnet
func Delete(db string, collection string, query map[string]interface{}) error {
	session := GetSession(db)
	defer session.Close()
	dbCollection := session.DB(db).C(collection)
	_, err := dbCollection.RemoveAll(query)
	return err
}

//GetCollectionWithoutPaging - Get documents from a collection based on given query
func GetCollectionWithoutPaging(db string, collection string, query map[string]interface{}) ([]interface{}, error) {
	result := make([]interface{}, 0)
	session := GetSession(db)
	defer session.Close()
	coll := session.DB(db).C(collection)
	q := coll.Find(query).All(&result)
	if q != nil && q.Error() != "" {
		log.Println(q.Error())
	}
	return result, q
}

// GenerateID - Generate a new time based ID
func GenerateID() string {
	return strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
}
