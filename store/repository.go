package store

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"strings"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "mongodb://gautam:gautam@ds157233.mlab.com:57233/dummystore"

// DBNAME the name of the DB instance
const DBNAME = "dummyStore"

// COLLECTION is the name of the collection in DB
const COLLECTION = "store"

var productId = 10;

// GetProducts returns the list of Products
func (r Repository) GetProducts() Products {
	session, err := mgo.Dial(SERVER)

	if err != nil {
	 	fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	results := Products{}

	if err := c.Find(nil).All(&results); err != nil {
	  	fmt.Println("Failed to write results:", err)
	}

	return results
}

// GetProductById returns a unique Product
func (r Repository) GetProductById(id int) Product {
	session, err := mgo.Dial(SERVER)

	if err != nil {
	 	fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	var result Product

	fmt.Println("ID in GetProductById", id);

	if err := c.FindId(id).One(&result); err != nil {
	  	fmt.Println("Failed to write result:", err)
	}

	return result
}

// GetProductsByString takes a search string as input and returns products
func (r Repository) GetProductsByString(query string) Products {
	session, err := mgo.Dial(SERVER)

	if err != nil {
	 	fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	result := Products{}

	// Logic to create filter
	qs := strings.Split(query, " ")
	and := make([]bson.M, len(qs))
	for i, q := range qs {
    	and[i] = bson.M{"title": bson.M{
        	"$regex": bson.RegEx{Pattern: ".*" + q + ".*", Options: "i"},
    	}}
	}
	filter := bson.M{"$and": and}

	if err := c.Find(&filter).Limit(5).All(&result); err != nil {
	  	fmt.Println("Failed to write result:", err)
	}

	return result
}

// AddProduct adds a Product in the DB
func (r Repository) AddProduct(product Product) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	productId += 1
	product.ID = productId
	session.DB(DBNAME).C(COLLECTION).Insert(product)
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added New Product ID- ", product.ID)

	return true
}

// UpdateProduct updates a Product in the DB
func (r Repository) UpdateProduct(product Product) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	err = session.DB(DBNAME).C(COLLECTION).UpdateId(product.ID, product)
	
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Updated Product ID - ", product.ID)

	return true
}

// DeleteProduct deletes an Product
func (r Repository) DeleteProduct(id int) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Remove product
	if err = session.DB(DBNAME).C(COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	fmt.Println("Deleted Product ID - ", id)
	// Write status
	return "OK"
}
