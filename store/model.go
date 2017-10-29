package store

// Product represents an e-comm item
type Product struct {
	ID     int 		 	 `bson:"_id"`
	Title  string        `json:"title"`
	Image  string        `json:"image"`
	Price   uint64       `json:"price"`
	Rating  uint8        `json:"rating"`
}

// Products is an array of Product objects
type Products []Product