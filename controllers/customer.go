package controllers

import (
    "encoding/json"
    "net/http"
    "merchant-bank-api/services"
    "merchant-bank-api/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
    var creds struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    json.NewDecoder(r.Body).Decode(&creds)

    token, err := services.Login(creds.Email, creds.Password)
    if err != nil {
        utils.ResponseJSON(w, "Login failed: "+err.Error(), http.StatusUnauthorized)
        return
    }

    utils.ResponseJSON(w, map[string]string{"token": token}, http.StatusOK)
}

func Payment(w http.ResponseWriter, r *http.Request) {
    customerID := r.Context().Value("customerID").(string)
    var req struct {
        Amount float64 `json:"amount"`
    }
    json.NewDecoder(r.Body).Decode(&req)

    err := services.Payment(customerID, req.Amount)
    if err != nil {
        utils.ResponseJSON(w, "Payment failed: "+err.Error(), http.StatusBadRequest)
        return
    }

    utils.ResponseJSON(w, "Payment successful", http.StatusOK)
}
