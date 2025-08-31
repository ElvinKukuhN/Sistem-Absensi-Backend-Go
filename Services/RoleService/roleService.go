package RoleService

import (
	"Sistem-Absensi-Backend-Go/Database"
	"Sistem-Absensi-Backend-Go/Dto/RoleDto"
	"Sistem-Absensi-Backend-Go/Models/Entity"
	"errors"
	"time"
)

func CreateRole(role *RoleDto.RoleDto) (*RoleDto.RoleResponse, error) {
	db := Database.DB

	var existing Entity.Role
	if err := db.Where("upper(name) = ?", role.Name).First(&existing).Error; err == nil {
		return nil, errors.New("role already exists")
	}
	if role.Name == "" {
		return nil, errors.New("name is required")
	}

	roles := &Entity.Role{
		Name:      role.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&roles).Error; err != nil {
		return nil, err
	}

	resp := &RoleDto.RoleResponse{
		Id:   roles.Id,
		Name: roles.Name,
	}

	return resp, nil
}

func GetAllRoles() ([]RoleDto.RoleResponse, error) {
	db := Database.DB

	var roles []Entity.Role
	if err := db.Find(&roles).Error; err != nil {
		return nil, err
	}

	var resp []RoleDto.RoleResponse
	for _, role := range roles {
		resp = append(resp, RoleDto.RoleResponse{
			Id:   role.Id,
			Name: role.Name,
		})
	}

	return resp, nil
}

func GetRoleById(id string) (*RoleDto.RoleResponse, error) {
	db := Database.DB

	var role Entity.Role
	if err := db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}

	resp := &RoleDto.RoleResponse{
		Id:   role.Id,
		Name: role.Name,
	}

	return resp, nil
}

func DeleteRole(id int) (*RoleDto.RoleResponse, error) {
	db := Database.DB

	var role Entity.Role
	if err := db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}

	if err := db.Delete(&role).Error; err != nil {
		return nil, err
	}
	resp := &RoleDto.RoleResponse{
		Name: role.Name,
	}
	return resp, nil
}
