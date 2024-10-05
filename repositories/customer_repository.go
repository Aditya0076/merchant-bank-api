package repositories

import (
    "encoding/json"
    "errors"
    "io/ioutil"
    "os"
    "merchant-bank-api/models"
)

var customersFile = "data/customers.json"

// Temukan customer berdasarkan email
func FindCustomerByEmail(email string) (*models.Customer, error) {
    customers, err := readCustomers()
    if err != nil {
        return nil, err
    }

    for _, customer := range customers {
        if customer.Email == email {
            return &customer, nil
        }
    }

    return nil, errors.New("customer not found")
}

// Update customer data
func UpdateCustomer(updatedCustomer *models.Customer) error {
    customers, err := readCustomers()
    if err != nil {
        return err
    }

    for i, customer := range customers {
        if customer.ID == updatedCustomer.ID {
            customers[i] = *updatedCustomer
        }
    }

    return writeCustomers(customers)
}

// Membaca data customer dari file JSON
func readCustomers() ([]models.Customer, error) {
    file, err := os.Open(customersFile)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var customers []models.Customer
    byteValue, _ := ioutil.ReadAll(file)
    json.Unmarshal(byteValue, &customers)
    return customers, nil
}

// Menulis data customer ke file JSON
func writeCustomers(customers []models.Customer) error {
    data, err := json.Marshal(customers)
    if err != nil {
        return err
    }
    return ioutil.WriteFile(customersFile, data, 0644)
}
