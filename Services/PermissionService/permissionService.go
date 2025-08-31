package PermissionService

import (
	"Sistem-Absensi-Backend-Go/Database"
	"Sistem-Absensi-Backend-Go/Dto/PermissionDto"
	"Sistem-Absensi-Backend-Go/Models/Entity"
	"errors"
	"time"
)

func CreatePermission(permission *PermissionDto.PermissionDto) (*PermissionDto.PermissionResponse, error) {
	db := Database.DB

	var existing Entity.Permission
	if err := db.Where("upper(code) = ?", permission.Code).First(&existing).Error; err == nil {
		return nil, errors.New("permission already exists")
	}
	if permission.Code == "" {
		return nil, errors.New("code is required")
	}

	permissions := &Entity.Permission{
		Code:      permission.Code,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&permissions).Error; err != nil {
		return nil, err
	}

	resp := &PermissionDto.PermissionResponse{
		Id:        permissions.Id,
		Code:      permissions.Code,
		CreatedAt: permissions.CreatedAt.Format(time.RFC3339),
		UpdatedAt: permissions.UpdatedAt.Format(time.RFC3339),
	}

	return resp, nil
}

func GetAllPermissions() ([]PermissionDto.PermissionResponse, error) {
	db := Database.DB

	var permissions []Entity.Permission
	if err := db.Find(&permissions).Error; err != nil {
		return nil, err
	}

	var resp []PermissionDto.PermissionResponse
	for _, permission := range permissions {
		resp = append(resp, PermissionDto.PermissionResponse{
			Id:        permission.Id,
			Code:      permission.Code,
			CreatedAt: permission.CreatedAt.Format(time.RFC3339),
			UpdatedAt: permission.UpdatedAt.Format(time.RFC3339),
		})
	}

	return resp, nil
}

func GetPermissionById(id string) (*PermissionDto.PermissionResponse, error) {
	db := Database.DB

	var permission Entity.Permission
	if err := db.Where("id = ?", id).First(&permission).Error; err != nil {
		return nil, err
	}

	resp := &PermissionDto.PermissionResponse{
		Id:        permission.Id,
		Code:      permission.Code,
		CreatedAt: permission.CreatedAt.Format(time.RFC3339),
		UpdatedAt: permission.UpdatedAt.Format(time.RFC3339),
	}

	return resp, nil
}

func DeletePermission(id int) (*PermissionDto.PermissionResponse, error) {
	db := Database.DB

	var permission Entity.Permission
	if err := db.Where("id = ?", id).First(&permission).Error; err != nil {
		return nil, err
	}

	if err := db.Delete(&permission).Error; err != nil {
		return nil, err
	}
	resp := &PermissionDto.PermissionResponse{
		Code: permission.Code,
	}
	return resp, nil
}
