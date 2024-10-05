package main

import (
    "log"
    "net/http"
    "merchant-bank-api/routes"
)

func main() {
    r := routes.RegisterRoutes()

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
