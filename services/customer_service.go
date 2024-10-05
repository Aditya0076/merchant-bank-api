package services

import (
    "errors"
    "time"
    "merchant-bank-api/models"
    "merchant-bank-api/repositories"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

// Login service untuk otentikasi customer
func Login(email, password string) (string, error) {
    customer, err := repositories.FindCustomerByEmail(email)
    if err != nil {
        return "", errors.New("customer not found")
    }

    // Bandingkan password hash
    err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
    if err != nil {
        return "", errors.New("invalid password")
    }

    // Buat token JWT
    expirationTime := time.Now().Add(5 * time.Hour)
    claims := &Claims{
        Email: customer.Email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", errors.New("error creating token")
    }

    // Update status logged in
    customer.LoggedIn = true
    repositories.UpdateCustomer(customer)

    return tokenString, nil
}
