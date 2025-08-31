package AuthController

import (
	"Sistem-Absensi-Backend-Go/Dto/AuthDto"
	"Sistem-Absensi-Backend-Go/Services/AuthService"
	"Sistem-Absensi-Backend-Go/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	var req AuthDto.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Utils.Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	user, err := AuthService.SignUp(req.Name, req.Email, req.Password, req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, Utils.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Utils.Response{
		Code:    http.StatusCreated,
		Success: true,
		Message: "User created successfully",
		Data:    user,
	})
}

func SignInHandler(c *gin.Context) {
	var req AuthDto.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Utils.Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	user, err := AuthService.SignIn(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, Utils.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Utils.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "User signed in successfully",
		Data:    user,
	})
}
