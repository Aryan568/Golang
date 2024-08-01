package main

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserData struct {
	Email string `json:"email"`
}

var (
	users     = make(map[string]string) // email -> hashed password
	userMutex sync.RWMutex
)

const (
	adminEmail    = "admin@example.com"
	adminPassword = "admin1000023"         // In a real application, use a strong, hashed password
	jwtSecret     = "hsuhJIHgJH890hjBSJFB" // Use a strong, randomly generated key in production
)

func main() {
	r := gin.Default()

	r.POST("/register", registerHandler)
	r.POST("/login", loginHandler)
	r.GET("/admin/users", authMiddleware(), adminUsersHandler)

	r.Run(":8080")
}

func registerHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userMutex.Lock()
	defer userMutex.Unlock()

	if _, exists := users[user.Email]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	users[user.Email] = string(hashedPassword)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func loginHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userMutex.RLock()
	hashedPassword, userExists := users[user.Email]
	userMutex.RUnlock()

	// For admin, we compare directly with the plaintext admin password.
	var err error
	if user.Email == adminEmail {
		// Hash the admin password for comparison
		adminHashedPassword, _ := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
		err = bcrypt.CompareHashAndPassword(adminHashedPassword, []byte(user.Password))
	} else if userExists {
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	} else {
		err = bcrypt.ErrMismatchedHashAndPassword
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)
			if email != adminEmail {
				c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
				c.Abort()
				return
			}
			c.Set("email", email)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func adminUsersHandler(c *gin.Context) {
	userMutex.RLock()
	defer userMutex.RUnlock()

	var userData []UserData
	for email := range users {
		userData = append(userData, UserData{Email: email})
	}

	c.JSON(http.StatusOK, userData)
}
