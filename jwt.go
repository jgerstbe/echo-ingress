package main

import (
    "fmt"
	"strings"
    "github.com/dgrijalva/jwt-go"
)

func getUniqueUsername(bearerToken string) (string, error) {
    // Parse the token without verification
    token, _, err := new(jwt.Parser).ParseUnverified(strings.TrimPrefix(bearerToken, "Bearer "), jwt.MapClaims{})
    if err != nil {
        return "", fmt.Errorf("Token parsing failed: %v", err)
    }

    // Access the claim
    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        if username, found := claims["unique_username"].(string); found {
            return username, nil
        }
    }

    return "", fmt.Errorf("Username claim not found in the JWT")
}
