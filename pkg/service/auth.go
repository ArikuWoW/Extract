package service

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ArikuWoW/extract/models"
	"github.com/ArikuWoW/extract/pkg/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {

		logrus.Errorf("Failed to hash password: %s", err)
		return ""
	}
	return string(hPass)
}

func (s *AuthService) GetUser(login, password string, c *gin.Context) {

	user, err := s.repo.GetUser(login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid login",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	fmt.Print(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{

			"error": "Invalid password",
		})
		fmt.Print(err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		UserId: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			Subject:   fmt.Sprint(user.Id),
		},
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
