package handlers

import (
	"fmt"
	"iis_server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	//RefreshToken string `json:"refresh_token"`
	Message string `json:"message"`
}

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
	Message     string `json:"message"`
}

var RefreshToken string

func LoginHandler(c *gin.Context) {
	var creds LoginCredentials

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	fmt.Printf("Login attempt for user: %s\n", creds.Username)

	const hardcodedUsername = "admin"
	const hardcodedPassword = "12345678"
	const hardcodedUserID = "1"

	if creds.Username != hardcodedUsername || creds.Password != hardcodedPassword {
		fmt.Printf("Login failed for user: %s (invalid credentials)\n", creds.Username)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"}) // 401
		return
	}

	accessToken, err := utils.GenerateAccessToken(hardcodedUserID)
	if err != nil {
		fmt.Printf("Error generating access token: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(hardcodedUserID)
	if err != nil {
		fmt.Printf("Error generating refresh token: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}
	RefreshToken = refreshToken
	fmt.Printf("Login successful for user: %s\n", creds.Username)

	response := LoginResponse{
		AccessToken: accessToken,
		//RefreshToken: refreshToken,
		Message: "Login successful",
	}
	c.JSON(http.StatusOK, response)
}

func RefreshTokenHandler(c *gin.Context) {

	refreshTokenString := RefreshToken
	claims, err := utils.ValidateToken(refreshTokenString)
	if err != nil {
		fmt.Printf("Refresh token validation failed: %v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	fmt.Printf("Refresh token validated successfully for user: %s\n", claims.UserID)

	newAccessToken, err := utils.GenerateAccessToken(claims.UserID)
	if err != nil {
		fmt.Printf("Error generating new access token during refresh: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new access token"})
		return
	}

	newRefreshToken, err := utils.GenerateRefreshToken(claims.UserID)
	if err != nil {
		fmt.Printf("Error generating new refresh token during refresh: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new refresh token"})
		return
	}

	RefreshToken = newRefreshToken

	response := RefreshResponse{
		AccessToken: newAccessToken,
		Message:     "Access token refreshed successfully",
	}
	c.JSON(http.StatusOK, response)
}

func LogoutHandler(c *gin.Context) {
	RefreshToken = ""
	c.AbortWithStatus(http.StatusNoContent)
}
