package UserDto

import "Sistem-Absensi-Backend-Go/Dto/RoleDto"

type UserResponse struct {
	Id        string                 `json:"id"`
	Name      string                 `json:"name"`
	Email     string                 `json:"email"`
	Role      []RoleDto.RoleResponse `json:"role,omitempty"`
	IsActive  string                 `json:"is_active"`
	CreatedAt string                 `json:"created_at,omitempty" gorm:"column:createdAt"`
	UpdatedAt string                 `json:"updated_at,omitempty" gorm:"column:updatedAt"`
}

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     []int  `json:"role" validate:"required"`
}
