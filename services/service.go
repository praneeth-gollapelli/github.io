package services

import (
	"errors"
	"log"

	"github.io/constants"
	mongo "github.io/util/db"
)

//Create
func Create(data map[string]interface{}) error {
	if len(data) == 0 {
		return errors.New("Invalid input received")
	}
	_, err := mongo.Create(constants.DB, constants.COLL, data)
	return err
}

//Update
func Update(query, data map[string]interface{}) error {
	if len(data) == 0 {
		return errors.New("Invalid input received")
	}
	return mongo.Update(constants.DB, constants.COLL, query, data)
}

//Destroy
func Destroy(params map[string]interface{}) error {
	err := mongo.Delete(constants.DB, constants.COLL, params)
	if err != nil {
		log.Println(err)
	}
	return err
}

func Find(query map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{}, 0)
	result["rows"] = make([]interface{}, 0)
	result["records"] = 0
	result["total"] = 0
	rows, err := mongo.GetCollectionWithoutPaging(constants.DB, constants.COLL, query)
	if err != nil {
		log.Println(err)
	}
	result["rows"] = rows
	return result
}
