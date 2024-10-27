package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

type SignupResponse struct {
	Message string `json:"message"`
}

// LoginResponse represents the response structure for the login endpoint.
type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// signup godoc
// @Summary Create a new user
// @Description Register a new user with the provided data
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User information"
// @Success 201 {object} SignupResponse
// @Failure 400 {object} SignupResponse "Could not parse request data."
// @Failure 500 {object} SignupResponse "Could not save user."
// @Router /signup [post]
func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	fmt.Print(user)

	err = user.Save()
	fmt.Println("Error saving user:", err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// login godoc
// @Summary Authenticate a user
// @Description Log in a user and generate a token
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User login information"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} SignupResponse "Could not parse request data."
// @Failure 401 {object} SignupResponse "Could not authenticate user."
// @Failure 500 {object} SignupResponse "Could not authenticate user."
// @Router /login [post]
func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
