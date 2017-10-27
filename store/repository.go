package store

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "store"

// DOCNAME the name of the document
const DOCNAME = "products"

// GetProducts returns the list of Products
func (r Repository) GetProducts() Products {
	session, err := mgo.Dial(SERVER)

	if err != nil {
	 	fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(DOCNAME)
	results := Products{}

	if err := c.Find(nil).All(&results); err != nil {
	  	fmt.Println("Failed to write results:", err)
	}

	return results
}

// AddProduct inserts an Product in the DB
func (r Repository) AddProduct(product Product) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	product.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(product)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

// UpdateProduct updates an Product in the DB (not used for now)
func (r Repository) UpdateProduct(product Product) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	product.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).UpdateId(product.ID, product)
	
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

// DeleteProduct deletes an Product (not used for now)
func (r Repository) DeleteProduct(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	// Write status
	return "OK"
}