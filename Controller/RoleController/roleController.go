package RoleController

import (
	"Sistem-Absensi-Backend-Go/Dto/RoleDto"
	"Sistem-Absensi-Backend-Go/Services/RoleService"
	"Sistem-Absensi-Backend-Go/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateHandler(c *gin.Context) {
	var req RoleDto.RoleDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Utils.Response{
			Success: false,
			Message: "Invalid request",
		})
		return
	}

	role, err := RoleService.CreateRole(&req)
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
		Message: "Role created successfully",
		Data:    role,
	})

}

func GetHandler(c *gin.Context) {
	id := c.Query("id")
	var resp interface{}
	if id != "" {
		role, err := RoleService.GetRoleById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, Utils.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		resp = role
	} else {
		allRoles, err := RoleService.GetAllRoles()
		if err != nil {
			c.JSON(http.StatusBadRequest, Utils.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		resp = allRoles
	}

	c.JSON(http.StatusOK, Utils.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "Role retrieved successfully",
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

	role, err := RoleService.DeleteRole(id)
	if err != nil || role == nil {
		c.JSON(http.StatusBadRequest, Utils.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Utils.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: "Role deleted successfully",
		Data:    role.Name,
	})
}
