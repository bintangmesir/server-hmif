package utils

import (
	"errors"
	"server/initialize"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
    refreshToken = []byte(initialize.ENV_REFRESH_TOKEN)
    accessToken = []byte(initialize.ENV_ACCESS_TOKEN)    
)

func GenerateToken(userID uuid.UUID, role string, tokenType string) (string, error) {
    claims := jwt.MapClaims{
        "id":   userID,
        "role":   role,
        "type":   tokenType,
        "exp":    time.Now().Add(time.Hour * 1).Unix(),
    }
     
    if tokenType == "refresh" {
        claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        return token.SignedString(refreshToken)    
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(accessToken)
}

func ParseToken(tokenStr string, tokenType string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {        
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }        
        if tokenType != "refresh"{
            return accessToken, nil    
        }
        return refreshToken, nil
    })

    if err != nil {
        return nil, err
    }
    
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, errors.New("invalid token")
}
