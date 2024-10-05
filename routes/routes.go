package routes

import (
    "net/http"
    "merchant-bank-api/controllers"
    "merchant-bank-api/middlewares"
    "github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
    r := mux.NewRouter()

    // Customer routes
    r.HandleFunc("/login", controllers.Login).Methods("POST")
    r.Handle("/payment", middlewares.AuthMiddleware(http.HandlerFunc(controllers.Payment))).Methods("POST")

    return r
}
