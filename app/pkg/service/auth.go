package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/SaRgEX/Diplom/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "asdf422985asdji"
	signingKey = "asd4qwe@qwioqtuqpte"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserPayload model.UserPayload
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) CreateUserWithRole(user model.UserWithRole) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUserWithRole(user)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		model.UserPayload{
			UserId: user.Id,
			Role:   user.Role,
		},
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accesstoken string) (model.UserPayload, error) {
	token, err := jwt.ParseWithClaims(accesstoken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return model.UserPayload{}, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return model.UserPayload{}, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserPayload, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) Logout(token string) error {
	return s.repo.Logout(token)
}
