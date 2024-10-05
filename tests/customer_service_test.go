package tests

import (
    "testing"
    "merchant-bank-api/services"
)

func TestLogin(t *testing.T) {
    token, err := services.Login("test@test.com", "password123")
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if token == "" {
        t.Fatalf("Expected token, got empty string")
    }
}
