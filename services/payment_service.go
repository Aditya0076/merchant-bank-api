package services

import (
    "errors"
    "merchant-bank-api/repositories"
    "merchant-bank-api/models"
    "time"
)

func Payment(customerID string, amount float64) error {
    customer, err := repositories.FindCustomerByID(customerID)
    if err != nil || !customer.LoggedIn {
        return errors.New("customer not logged in")
    }

    // Simulasi transfer (hanya mencatat di history)
    action := "Payment of $" + fmt.Sprintf("%.2f", amount)
    log := models.History{
        CustomerID: customer.ID,
        Action:     action,
        Timestamp:  time.Now().Format(time.RFC3339),
    }

    repositories.SaveHistory(log)
    return nil
}
