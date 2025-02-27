package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const TokenExpire = 12
const AuthCookie = "auth_token"

func ExpiredCookie() *http.Cookie {
    return &http.Cookie{
        Name:     "auth_token",
        Value:    "",
        Path:     "/",
        Expires:  time.Unix(0, 0), // Expire immediately
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteLaxMode,
    }
}


func SignUser(userID string, userEmail string) (string, error) {
    hmacSecret := os.Getenv("HMAC_SECRET")
    if hmacSecret == "" {
        return "", fmt.Errorf("HMAC_SECRET environment variable is missing")
    }
    expTime := time.Now().UTC().Add(time.Hour * TokenExpire).Unix() // 1-day expiration
    jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "UID":   userID,
        "Email": userEmail,
        "exp": expTime,
    })
    tokenString, err := jwtToken.SignedString([]byte(hmacSecret))
    if err != nil {
        return "", err
    }
    return tokenString, err
}

func ParseUser(tokenString string) (string, error) {
    hmacSecret := os.Getenv("HMAC_SECRET")
    if hmacSecret == "" {
        return "", fmt.Errorf("HMAC_SECRET environment variable is missing")
    }
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(hmacSecret), nil
    })
    if err != nil {
        return "", err
    }
    if !token.Valid {
        return "", fmt.Errorf("invalid token")
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", fmt.Errorf("cannot extract claims")
    }
    uid, ok := claims["UID"].(string)
    if !ok {
        return "", fmt.Errorf("UID not found or invalid")
    }
    return uid, nil
}

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        cookie, err := c.Cookie(AuthCookie)
        if err != nil {
            c.SetCookie(ExpiredCookie())
            return next(c)
        }
        userID, err := ParseUser(cookie.Value)
        if err != nil {
            c.SetCookie(ExpiredCookie())
            return next(c)
        }
        c.Set("userID", userID)
        return next(c)
    }
}

