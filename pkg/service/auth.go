package service

import (
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
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
