package token

import (
    "fmt"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var signlePassword = []byte("linklobby-go")

func createToken(usr string) *string {
    tokhead := map[string]interface{}{
        "typ": "JWT",
        "alg": "HS256",
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "iss":   "/.well-known/public.key.json",
        "usr":   usr,
        "nbf":   time.Now().Unix(),
        "iat":   time.Now().Unix(),
        "aud":   usr, 
        "exp":   time.Now().Add(time.Hour * 48).Unix(),
        "scope": "service.socket.user.use service.socker.user.create lobby.user.create lobby.user.delete lobby.report lobby.user.blocklist",
    })

    token.Header = tokhead
    tokenString, err := token.SignedString(mySigningKey)
    if err != nil {
        fmt.Println("Error generating token:", err)
        return
    }

    return tokenString
}

