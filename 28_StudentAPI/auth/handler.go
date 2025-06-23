package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

// Handles POST /auth/signup
func (h *AuthHandler) Signup(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.Service.Signup(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", token, 1*60*60, "/", "localhost", false, true)
	// "token" : The name of the cookie
	// token : token value that need to store
	// maxAge : LifeTime of coolie in seconds(3600 seconds = 1 hour)
	// path "/" : The path for which the cookie is valid(root means all path)
	// "localhost" : the domain , the cookie is valid for
	// secure (false) : if true, cookie is sent inlt over HTTPS
	// httpOnly : if true, cookie is inaccessible to javaScript(helps prevent XSS attacks)

	c.JSON(http.StatusOK, gin.H{"token": token, "message": "Login successful"})
}
