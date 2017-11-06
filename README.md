# rest-and-go
A basic online store API written to learn Go Programming Language

This API is a pretty basic implementation of an online(e-commerce) store.
- You can perform basic CRUD(CREATE, READ, UPDATE and DELETE) operations
- SEARCH on a predefined database of products 
- Only Authenticated users can Add, Update and Delete products from database
- Authentication is based on JWT(JSON web Tokens) Tokens
- API is backed by a predefined Mongo DB database hosted on [mLab](https://mLab.com)
- This API also lives on [Heroku](https://www.heroku.com) - https://gruesome-monster-22811.herokuapp.com/ 

See [API Documentation and Usage](#api-documentation-and-usage) below on how to use it.

## Directory Structure
```
rest-and-go/
    |- Godeps/             - Contains info about all dependencies of the project
    |- store/              - Contains main API logic files 
        |- controller.go  - Defines methods handling calls at various endpoints
        |- model.go       - User and Product models
        |- repository.go  - Methods interacting with the database
        |- router.go      - Defines routes and endpoints
    |- vendor/             - Dependency packages, necessary for deployment
    |- .gitignore
    |- LICENSE
    |- Procfile             - Procfile for herkou deployment
    |- README.md
    |- dummyData.js         - Script to populate local mongodb with dummy data
    |- main.go              - Entry point of the API
  
```

## Setup

### Golang Development Setup

You can use this bash script to automate the Golang development setup - https://github.com/canha/golang-tools-install-script

**Steps**
1. Download the repository using wget 
`wget https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh`
2. According to the OS you're on
    - Linux 64 bit -> `bash goinstall.sh --64`
    - Linux 32 bit -> `bash goinstall.sh --32`
    - macOS -> `bash goinstall.sh --darwin`

You can also follow the official [docs](https://golang.org/doc/install) of installation if you want to know the complete process.

### Project setup

1. Clone the repository in your `$GOPATH/src/` directory. If you have used the bash script for setup, your `$GOPATH` variable should point to `$HOME/go`
2. **Follow the steps 2-6 only if you have to set-up databse by yourself**. The MongoDB database is hosted on mLab free trial account for now and might expire. In that case, you'll need the steps below.
3. To run project locally, Install Mongo DB - https://www.mongodb.com/download-center?jmp=nav#community
4. After installing Mongo DB, start it's server by typing `mongod` in Terminal.
5. Open a new tab in terminal and type `mongo < dummyData.js` to insert the dummmy product data.
6. Open file `store/repository.go`, find the `SERVER` variable and replace the URL. 

```
const SERVER = "http://localhost:27017"
```
7. Last thing required to run the project, install all the go dependencies
```
// Library to handle jwt authentication 
$ go get "github.com/dgrijalva/jwt-go"

// Libraries to handle network routing
$ go get "github.com/gorilla/mux"
$ go get "github.com/gorilla/context"
$ go get "github.com/gorilla/handlers"

// mgo library for handling Mongo DB
$ go get "gopkg.in/mgo.v2"
```
Yay! Now we're ready to run the API :tada: <br>
8. Type `export PORT=8000` in Terminal and open http://localhost:8000 in your browser to see the products.

## API Documentation and Usage

It is **recommended** to install some extension to beautify JSON(like [JSON Formatter](https://chrome.google.com/webstore/detail/json-formatter/bcjindcccaagfpapjjmafapmmgkkhgoa)) if you're trying in a browser.

**Important** - Don't forget to define $PORT in your shell variables. <br>Example: `export PORT=8000`

```sh
BASE_URL = "http://localhost:$PORT"
'OR'
BASE_URL = https://gruesome-monster-22811.herokuapp.com/
```

### 1. View Products

- **Endpoint Name** - `Index`      <br>
- **Method** - `GET`               <br>
- **URL Pattern** - `/`            <br>
- **Usage** 
    - Open BASE_URL in browser
    - **Terminal/CURL**
    ```sh
    curl -X GET BASE_URL
    ```
- **Expected Response** - JSON containing all the products in database <br>
- **Example**
![Screenshot](/screenshots/All-Products.png?raw=true)

### 2. View Single Product

- **Endpoint Name** - `GetProduct`    <br>
- **Method** - `GET`                  <br>
- **URL Pattern** - `/products/{id}`  <br>
- **Usage**
    - Open BASE_URL/products/{id} in browser
    - **Terminal/CURL**
```
curl -X GET BASE_URL/products/{id} 
```
- **Expected Response** - Product with the {id} in database
- **NOTE** - There are only six(6) ids in the database, so 1 <= {id} <= 6   
- **Example**
![Screenshot](/screenshots/GetProduct-Request.png)

### 3. Search Product

- **Endpoint Name** - `SearchProduct`  <br>
- **Method** - `GET`                   <br>
- **URL Pattern** - `/Search/{query}`  <br>
- **Usage** - Browser OR curl        
- **BROWSER**
    - Open BASE_URL/Search/{query} in browser
    - **Terminal/CURL**
    ```sh
    curl -X GET BASE_URL/Search/{query}
    ```
- **Expected Response** - Products matching the search query <br>
- **Example**
![Screenshot](/screenshots/Search-Request.png)

### 4. Authentication
For **Adding**, **Updating** and **Deleting** products from database you must send a JWT token in Authentication header.

- **Endpoint Name** - `GetToken` <br>
- **Method** - `POST`            <br>
- **URL Pattern** - `/get-token` <br>
- **Usage** - CURL OR POSTMAN ONLY
    - **Terminal/CURL**
    ```sh
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{ username: "<YOUR_USERNAME>", password: "<RANDOM_PASSWORD>"}' \
    BASE_URL/get-token
    ```
- **Expected Response** - A JWT Authentication Token as shown below
- **Example**
![Screenshot](/screenshots/Authentication-Request.png)

### 5. Add Product

- **Endpoint Name** - `AddProduct` <br>
- **Method** - `POST`              <br>
- **URL Pattern** - `/AddProduct`  <br>
- **Usage** - CURL OR POSTMAN ONLY
    - **Terminal/CURL**
    ```sh
    curl -X POST \
    -H "Authorization: Bearer <ACCESS_TOKEN>" \
    -d '{ "_id": 11, 
        "title": "Memes",
        "image": "I am selling memes, hehe.",          
        "price": 1,
        "rating": 5
        }' \
    BASE_URL/AddProduct
    ```
- **Expected Response** - Addition successful without any error message. Check the logs in Terminal window which is running server. 
- **Example**
![Screenshot](/screenshots/AddProduct-Request.png)

### 6. Update Product

- **Endpoint Name** - `UpdateProduct` <br>
- **Method** - `PUT`                  <br>
- **URL Pattern** - `/UpdateProduct`  <br>
- **Usage** - CURL OR POSTMAN ONLY
    - **Terminal/CURL**
    ```sh
    curl -X PUT \
    -H "Authorization: Bearer <ACCESS_TOKEN>" \
    -d '{ "ID": 14, 
        "title": "Memes",
        "image": "I am not selling memes to you, hehe.",          
        "price": 1000,
        "rating": 5
        }' \
    BASE_URL/UpdateProduct
    ```
- **Expected Response** - Update successful without any error message. Check the logs in Terminal window which is running server. <br>
- **Example**
![Screenshot](/screenshots/Update-Request.png)

### 7. Delete Product

- **Endpoint Name** - `DeleteProduct` <br>
- **Method** - `DELETE` <br>
- **URL Pattern** - `/deleteProduct/{id}` <br>
- **Usage** - CURL OR POSTMAN ONLY
    - **Terminal/CURL**
    ```sh
    curl -X DELETE \
    -H "Authorization: Bearer <ACCESS_TOKEN>" \
    BASE_URL/deleteProduct/{id}
    ```
- **Expected Response** - Deletion successful without any error message. Check the logs in Terminal window which is running server. <br>
- **Example**
![Screenshot](/screenshots/Delete-Request.png)

## TODO
* [ ] Write unit tests to test every method
* [ ] Improve the code by proper exception handling
* [ ] Add repository badges like TravisCI, Better Code, Codacy etc.
* [ ] Create a REST API server project using this package as a boilerplate
* [ ] User and roles management
* [ ] Session management using JWT tokens



