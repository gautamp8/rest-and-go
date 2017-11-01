# rest-and-go
A basic online store API written to learn Go Programming Language

[![Build Status](https://travis-ci.org/jokamjohn/bucket_api.svg?branch=master)](https://travis-ci.org/jokamjohn/bucket_api)
[![BCH compliance](https://bettercodehub.com/edge/badge/jokamjohn/bucket_api?branch=master)](https://bettercodehub.com/)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/cfda51ef2f8946639eb34b11fa8b5480)](https://www.codacy.com/app/jokamjohn/bucket_api?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=jokamjohn/bucket_api&amp;utm_campaign=Badge_Grade)

This API is a pretty basic implementation of an online(e-commerce) store. It enables you to perform basic CRUD(CREATE, READ, UPDATE and DELETE) operations and SEARCH on a predefined database of products. Only Authenticated users can Add, Update and Delete products from database.

API is backed by a Mongodb database hosted on [!mLab](https://mLab.com) which contains the predefined list of products. Authentication is based on JWT(JSON web Tokens) Tokens

If you want to try directly, API is also deployed on [!Heroku](https://www.heroku.com) - https://gruesome-monster-22811.herokuapp.com/ 

See [API Documentation]((#api-documentation)) below on how to use it.

## Directory Structure
```
rest-and-go/
    |- Godeps             - Contains info about all dependencies of the project
    |- store              - Contains main API logic files 
        |- controller.go  - Defines methods handling calls at various endpoints
        |- model.go       - User and Product models
        |- repository.go  - Methods interacting with the database
        |- router.go      - Defines routes and endpoints
    |- vendor             - Dependency packages, necessary for deployment
  |- .gitignore
  |- LICENSE
  |- Procfile             - Procfile for herkou deployment
  |- README.md
  |- dummyData.js         - Script to populate local mongodb with dummy data
  |- main.go              - Entry point of the API
  
```


## API Documentation and Usage

It is **recommended** to install some extension to beautify JSON(like [!JSON Formatter](https://chrome.google.com/webstore/detail/json-formatter/bcjindcccaagfpapjjmafapmmgkkhgoa)) if you're trying in a browser.

// TODO: export PORT
```
BASE_URL = http://localhost:$PORT
'OR'
BASE_URL = https://gruesome-monster-22811.herokuapp.com/
```

### VIEW Products
```
ENDPOINT NAME - Index
METHOD - GET
URL PATTERN - /

USAGE - 
Browser
Open BASE_URL in browser

Terminal/CURL
curl -X GET BASE_URL

EXPECTED RESPONSE - JSON containing all the products in database
    
```

**Example**



### VIEW Product with a particular ID
```
ENDPOINT NAME - GetProduct
METHOD - GET
URL PATTERN - /products/{id}

USAGE - 
Browser
Open BASE_URL/products/{id} in browser

Terminal/CURL
curl -X GET BASE_URL/products/{id}

EXPECTED RESPONSE - Product with the {id} in database

NOTE - There are only six(6) ids in the database, so 1 <= {id} <= 6    
```

**Example**

### SEARCH Product
```
ENDPOINT NAME - SearchProduct
METHOD - GET
URL PATTERN - /Search/{query}

USAGE - Browser OR CURL
Browser -
Open BASE_URL/Search/{query} in browser

Terminal/CURL -
curl -X GET BASE_URL/Search/{query}

EXPECTED RESPONSE - Products matching the search query
```

**Example**


### Authentication
For **Adding**, **Updating** and **Deleting** products from database you must send a JWT token in Authentication header.

```
ENDPOINT NAME - GetToken
METHOD - POST
URL PATTERN - /get-token

USAGE - CURL OR POSTMAN ONLY

Terminal/CURL -

curl -X POST \
-H "Content-Type: application/json" \
-d '{ username: "<YOUR_USERNAME>", password: "<RANDOM_PASSWORD>"}' \
BASE_URL/get-token

EXPECTED RESPONSE - A JWT Authentication Token as shown below
```

**Example**

### ADD Product to Database

```
ENDPOINT NAME - AddProduct
METHOD - POST
URL PATTERN - /AddProduct

USAGE - CURL OR POSTMAN ONLY

Terminal/CURL -

curl -X POST \
-H "Authorization: Bearer <ACCESS_TOKEN>" \
-d '{ "_id": 11, 
    "title": "Memes",
    "image": "I am selling memes, hehe.",          
    "price": 1,
    "rating": 5
    }' \
BASE_URL/AddProduct

EXPECTED RESPONSE - Addition successful without any error message. Check the logs in Terminal window which is running server. 
```

**Example**

### UPDATE Product


ENDPOINT NAME - UpdateProduct
METHOD - PUT
URL PATTERN - /UpdateProduct

USAGE - CURL OR POSTMAN ONLY

Terminal/CURL -
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
EXPECTED RESPONSE - Update successful without any error message. Check the logs in Terminal window which is running server.

**Example**


### DELETE Product


ENDPOINT NAME - DeleteProduct
METHOD - DELETE
URL PATTERN - /deleteProduct/{id}

USAGE - CURL OR POSTMAN ONLY

Terminal/CURL -
```sh
curl -X DELETE \
-H "Authorization: Bearer <ACCESS_TOKEN>" \
BASE_URL/deleteProduct/{id}
```
EXPECTED RESPONSE - Deletion successful without any error message. Check the logs in Terminal window which is running server.

**Example**

