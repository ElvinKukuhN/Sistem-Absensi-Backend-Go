package UserController

import (
	"Sistem-Absensi-Backend-Go/Dto/UserDto"
	"Sistem-Absensi-Backend-Go/Services/UserService"
	"Sistem-Absensi-Backend-Go/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateHandler(c *gin.Context) {
	var req UserDto.UserRequest

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Utils.Response{
			Success: false,
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	// Validasi minimal field yang dibutuhkan
	if req.Name == "" || req.Email == "" || req.Password == "" || len(req.Role) == 0 {
		c.JSON(http.StatusBadRequest, Utils.Response{
			Success: false,
			Message: "Name, email, password, and at least one role are required.",
		})
		return
	}

	// Panggil service untuk buat user
	userResp, err := UserService.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Utils.Response{
			Success: false,
			Message: "Failed to create user: " + err.Error(),
		})
		return
	}

	// Return sukses
	c.JSON(http.StatusOK, Utils.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "User created successfully",
		Data:    userResp,
	})
}

func GetHandler(c *gin.Context) {
	id := c.Query("id")
	var resp interface{}
	if id != "" {
		user, err := UserService.GetUserById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, Utils.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		resp = user
	} else {
		allUsers, err := UserService.GetAllUser()
		if err != nil {
			c.JSON(http.StatusBadRequest, Utils.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		resp = allUsers
	}

	c.JSON(http.StatusOK, Utils.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "User retrieved successfully",
		Data:    resp,
	})

}

func UpdateHandler(c *gin.Context) {
	userId := c.Param("id")

	if userId == "" {
		c.JSON(http.StatusBadRequest, Utils.Response{
			Success: false,
			Message: "User ID is required",
		})
		return
	}

	var userReq UserDto.UserRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, Utils.Response{
			Success: false,
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	userResp, err := UserService.UpdateUser(userId, &userReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Utils.Response{
			Success: false,
			Message: "Failed to update user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Utils.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "User updated successfully",
		Data:    userResp,
	})
}
