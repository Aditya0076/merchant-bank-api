package middlewares

import (
    "context"
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt"
    "merchant-bank-api/services"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing authorization header", http.StatusUnauthorized)
            return
        }

        tokenStr := strings.Split(authHeader, " ")[1]
        claims := &services.Claims{}

        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return services.JwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        ctx := context.WithValue(r.Context(), "customerID", claims.Email)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
