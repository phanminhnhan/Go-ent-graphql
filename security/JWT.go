package security

import (
	"github.com/golang-jwt/jwt"
	"time"
	"errors"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

const SecretKey = "fshjofjsdfo8oi3wyuf98wyu9876uhzxiou#@"

type JwtCustomClaims struct {
	Id string
	Email string
	UserName string
	CExpiresAt time.Time
	jwt.StandardClaims
}

func (payload *JwtCustomClaims) Valid() error {
	if time.Now().After(payload.CExpiresAt) {
		return errors.New("Expire token")
	}
	return nil
}

func Gentoken(user_id, user_email, user_name string)(map[string]string, error){
	claims_access := &JwtCustomClaims{
		Id: user_id,
		Email : user_email,
		UserName: user_name,
		CExpiresAt: time.Now().Add(time.Minute * 10),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	}
	claims_refresh := &JwtCustomClaims{
		Id: user_id,
		Email : user_email,
		UserName: user_name,
		CExpiresAt: time.Now().Add(time.Hour * 24),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims_access)
	t, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return nil, err
	}
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims_refresh)
	rt, err := refresh_token.SignedString([]byte(SecretKey))
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"access_token": t,
		"refresh_token": rt,
	}, nil
}


func VerifyToken(token string) (*JwtCustomClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(SecretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*JwtCustomClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
