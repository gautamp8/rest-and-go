package store

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes {
    Route {
        "Authentication",
        "POST",
        "/get-token",
        controller.GetToken,
    },
    Route {
        "Index",
        "GET",
        "/",
        controller.Index,
    },
    Route {
        "AddProduct",
        "POST",
        "/AddProduct",
        AuthenticationMiddleware(controller.AddProduct),
    },
    Route {
        "UpdateProduct",
        "PUT",
        "/UpdateProduct",
        AuthenticationMiddleware(controller.UpdateProduct),
    },
    // Get Product by {id}
    Route {
        "GetProduct",
        "GET",
        "/products/{id}",
        controller.GetProduct,
    },
    // Delete Product by {id}
    Route {
        "DeleteProduct",
        "DELETE",
        "/deleteProduct/{id}",
        AuthenticationMiddleware(controller.DeleteProduct),
    },
    // Search product with string
    Route {
        "SearchProduct",
        "GET",
        "/Search/{query}",
        controller.SearchProduct,
    }}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes { 
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc
        
        router.
         Methods(route.Method).
         Path(route.Pattern).
         Name(route.Name).
         Handler(handler)
    }
    return router
}