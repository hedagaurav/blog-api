package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hedagaurav/blog-api/config"
	"github.com/hedagaurav/blog-api/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

// todo: Add otp verification for mobile number and email.

// Register handles user registration
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Save the user to the database
	db := config.ConnectDatabase()
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// todo: Send registration success email to the user.

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login handles user login
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Find the user by email
	var user models.User
	db := config.ConnectDatabase()
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		// Log the error for debugging
		log.Printf("Error finding user: %s, ERROR: %v", input.Email, err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		// Log the error for debugging
		log.Printf("Error finding user: %s, ERROR: %v", input.Email, err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		// Log the error for debugging
		log.Printf("Error generating token for user: %s, ERROR: %v", input.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// log the successful login
	log.Printf("User logged in successfully: %s, Login Time: %s.", input.Email, time.Now().Format(time.RFC3339))

	// save the token in the database with login time (optional)
	// db.Model(&user).Update("token", tokenString)
	// db.Model(&user).Update("login_time", time.Now())

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Logout handles user logout
func Logout(c *gin.Context) {
	// Invalidate the token by removing it from the client side
	// save the logout time in the database (optional)
	// db.Model(&user).Update("logout_time", time.Now())
	// log the successful logout
	log.Printf("User logged out successfully: %s, Logout Time: %s.", c.GetHeader("Authorization"), time.Now().Format(time.RFC3339))
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
