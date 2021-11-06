package jwt

import (
	"errors"
	"time"

	stdjwt "github.com/golang-jwt/jwt/v4"
)

const (
	expiresIn = 24 * time.Hour
)

var (
	ErrJWTAlgMismatch      = errors.New("JWT algorithm mismatch")
	ErrJWTInvalid          = errors.New("Invalid JWT provided")
	ErrJWTExpired          = errors.New("Expired JWT")
	ErrJWTUnknownAlgorithm = errors.New("unknown JWT signing algorithm")
)

type UserClaims struct {
	*stdjwt.RegisteredClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

func NewRegisteredClaims(expiresIn time.Duration) stdjwt.RegisteredClaims {
	now := time.Now()
	return stdjwt.RegisteredClaims{
		ExpiresAt: stdjwt.NewNumericDate(now.Add(expiresIn)),
		IssuedAt:  stdjwt.NewNumericDate(now),
	}
}

func NewUserClaims(username string, role string) UserClaims {
	rc := NewRegisteredClaims(expiresIn)
	uc := UserClaims{
		RegisteredClaims: &rc,
		Username:         username,
		Role:             role,
	}
	return uc
}

type Wrapper struct {
	Algorithm stdjwt.SigningMethod
	Secret    string
}

func NewHS256Wrapper(secret string) Wrapper {
	return Wrapper{
		Algorithm: stdjwt.SigningMethodHS256,
		Secret:    secret,
	}
}

func (w Wrapper) Encode(claims stdjwt.Claims) (string, error) {
	switch w.Algorithm {
	case stdjwt.SigningMethodHS256:
		token := stdjwt.NewWithClaims(w.Algorithm, claims)
		return token.SignedString([]byte(w.Secret))
	default:
		return "", ErrJWTUnknownAlgorithm
	}
}

func (w Wrapper) Decode(tokenStr string, claims stdjwt.Claims) (*stdjwt.Token, error) {
	token, err := stdjwt.ParseWithClaims(tokenStr, claims, func(t *stdjwt.Token) (interface{}, error) {
		return []byte(w.Secret), nil
	})

	if err != nil {
		if e, ok := err.(*stdjwt.ValidationError); ok {
			if e.Errors == stdjwt.ValidationErrorExpired {
				return nil, ErrJWTExpired
			}
		}
		return nil, err
	}

	if !token.Valid {
		return nil, ErrJWTInvalid
	}

	if token.Method.Alg() != w.Algorithm.Alg() {
		return nil, ErrJWTAlgMismatch
	}

	return token, nil
}
