package jwt

import (
	"time"

	j "github.com/golang-jwt/jwt"
)

// Payload is payload of JWT
type Payload map[string]string

// Issue issues a JWT
func Issue(p Payload, secret string, expire time.Duration) (string, error) {
	claims := j.MapClaims{}
	claims["exp"] = time.Now().Add(expire).Unix()
	claims["iat"] = time.Now().Unix()
	for k, v := range p {
		claims[k] = v
	}
	return j.NewWithClaims(j.SigningMethodHS256, claims).SignedString([]byte(secret))
}

// Parse parses a JWT
func Parse(token string, secret string) (Payload, error) {
	claims, err := j.Parse(token, func(token *j.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claimsMap := claims.Claims.(j.MapClaims)
	p := Payload{}
	for k, v := range claimsMap {
		if k == "exp" || k == "iat" {
			continue
		}
		p[k] = v.(string)
	}

	return p, nil
}

// Refresh refreshes a JWT
func Refresh(token string, secret string, expire time.Duration) (string, error) {
	p, err := Parse(token, secret)
	if err != nil {
		return "", err
	}
	return Issue(p, secret, expire)
}
