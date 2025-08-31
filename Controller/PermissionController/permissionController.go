package PermissionController

import (
	"Sistem-Absensi-Backend-Go/Dto/PermissionDto"
	"Sistem-Absensi-Backend-Go/Services/PermissionService"
	"Sistem-Absensi-Backend-Go/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateHandler(c *gin.Context) {
	var req PermissionDto.PermissionDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Utils.Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	permission, err := PermissionService.CreatePermission(&req)
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
		Message: "Permission created successfully",
		Data:    permission,
	})

}

func GetHandler(c *gin.Context) {
	id := c.Query("id")
	var resp interface{}
	if id != "" {
		permission, err := PermissionService.GetPermissionById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, Utils.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		resp = permission
	} else {
		allPermissions, err := PermissionService.GetAllPermissions()
		if err != nil {
			c.JSON(http.StatusBadRequest, Utils.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		resp = allPermissions
	}

	c.JSON(http.StatusOK, Utils.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "Permission retrieved successfully",
		Data:    resp,
	})
}

func DeleteHandler(c *gin.Context) {
	idParam := c.Query("id")

	if idParam == "" {
		c.JSON(http.StatusBadRequest, Utils.Response{
			Success: false,
			Message: "Missing 'id' query parameter",
		})
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, Utils.Response{
			Success: false,
			Message: "Invalid 'id' value, must be an integer",
		})
		return
	}

	permission, err := PermissionService.DeletePermission(id)
	if err != nil || permission == nil {
		c.JSON(http.StatusBadRequest, Utils.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Utils.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "Permission deleted successfully",
		Data:    permission.Code,
	})
}
