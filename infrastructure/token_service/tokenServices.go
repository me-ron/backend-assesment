package tokenservice

import (
	"errors"
	"loan_tracker/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenService_imp struct {
	AccessTokenSecret       string
	RefreshTokenSecret      string
	VerificationTokenSecret string
}

func NewTokenService(accessSecret, refreshSecret, verificationsecret string) *TokenService_imp {
	return &TokenService_imp{
		AccessTokenSecret:       accessSecret,
		RefreshTokenSecret:      refreshSecret,
		VerificationTokenSecret: verificationsecret,
	}
}

func (t *TokenService_imp) GenerateAccessToken(userId string) (string, error) {
	claims := domain.UserClaims{
        ID:      userId,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
        },
    }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.AccessTokenSecret))
}

func (t *TokenService_imp) GenerateRefreshToken(userId string) (string, error) {
	claims := domain.UserClaims{
        ID:      userId,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
        },
    }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.RefreshTokenSecret))
}

func (t *TokenService_imp) ValidateAccessToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.AccessTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid access token")
	}

	claims, ok := token.Claims.(*domain.UserClaims)
    if !ok {
        return "", errors.New("invalid token claims")
    }

	return claims.ID, nil
}


func (t *TokenService_imp) ValidateRefreshToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.RefreshTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(*domain.UserClaims)
    if !ok {
        return "", errors.New("invalid token claims")
    }

	return claims.ID, nil
}

func (t *TokenService_imp) GenerateVerificationToken(userId string) (string, error) {
	claims := domain.UserClaims{
		ID:      userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.VerificationTokenSecret))
}

func (t *TokenService_imp) ValidateVerificationToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.VerificationTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid verification token")
	}

	claims, ok := token.Claims.(*domain.UserClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	return claims.ID, nil
}

func (t *TokenService_imp) GeneratePasswordToken(userId string) (string, error) {
	claims := domain.UserClaims{
		ID:      userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.VerificationTokenSecret))
}

func (t *TokenService_imp) ValidatePasswordToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.VerificationTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid password token")
	}

	claims, ok := token.Claims.(*domain.UserClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	return claims.ID, nil
}