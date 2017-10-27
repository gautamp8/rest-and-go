package store

import (
    "encoding/json"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "github.com/gorilla/mux"
)

//Controller ...
type Controller struct {
    Repository Repository
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
    products := c.Repository.GetProducts() // list of all products
    log.Println(products)
    data, _ := json.Marshal(products)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
    return
}

// AddProduct POST /
func (c *Controller) AddProduct(w http.ResponseWriter, r *http.Request) {
    var product Product
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
    
    if err != nil {
        log.Fatalln("Error AddProduct", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    if err := r.Body.Close(); err != nil {
        log.Fatalln("Error AddProduct", err)
    }

    if err := json.Unmarshal(body, &product); err != nil { // unmarshall body contents as a type Candidate
    w.WriteHeader(422) // unprocessable entity

    if err := json.NewEncoder(w).Encode(err); err != nil {
        log.Fatalln("Error AddProduct unmarshalling data", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    }

    success := c.Repository.AddProduct(product) // adds the product to the DB
    if !success {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    return
}

// UpdateProduct PUT /
func (c *Controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
    var product Product
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
    if err != nil {
        log.Fatalln("Error UpdateProduct", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    if err := r.Body.Close(); err != nil {
        log.Fatalln("Error AddaUpdateProductlbum", err)
    }

    if err := json.Unmarshal(body, &product); err != nil { // unmarshall body contents as a type Candidate
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            log.Fatalln("Error UpdateProduct unmarshalling data", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
    }

    success := c.Repository.UpdateProduct(product) // updates the product in the DB
    
    if !success {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    return
}

// DeleteProduct DELETE /
func (c *Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"] // param id
    if err := c.Repository.DeleteProduct(id); err != "" { // delete a product by id
        if strings.Contains(err, "404") {
            w.WriteHeader(http.StatusNotFound)
        } else if strings.Contains(err, "500") {
            w.WriteHeader(http.StatusInternalServerError)
        }
        return
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    return
}