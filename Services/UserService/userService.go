package UserService

import (
	"Sistem-Absensi-Backend-Go/Database"
	"Sistem-Absensi-Backend-Go/Dto/RoleDto"
	"Sistem-Absensi-Backend-Go/Dto/UserDto"
	"Sistem-Absensi-Backend-Go/Models/Entity"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func CreateUser(userReq *UserDto.UserRequest) (UserDto.UserResponse, error) {
	db := Database.DB
	var existingUser Entity.User
	var validRoles []Entity.Role

	// 1. Cek user berdasarkan email
	if err := db.Where("email = ?", userReq.Email).First(&existingUser).Error; err == nil {
		return UserDto.UserResponse{}, errors.New("user already exists")
	} else if err != gorm.ErrRecordNotFound {
		return UserDto.UserResponse{}, err
	}

	// 2. Validasi role-role yang ada
	if err := db.Where("id IN ?", userReq.Role).Find(&validRoles).Error; err != nil {
		return UserDto.UserResponse{}, fmt.Errorf("failed to validate roles: %v", err)
	}

	if len(validRoles) != len(userReq.Role) {
		return UserDto.UserResponse{}, fmt.Errorf("one or more role IDs are invalid")
	}

	// 3. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserDto.UserResponse{}, errors.New("failed to hash password")
	}

	// 4. Buat user baru
	userID := uuid.New().String()
	newUser := Entity.User{
		Id:        userID,
		Name:      userReq.Name,
		Email:     userReq.Email,
		Password:  string(hashedPassword),
		IsActive:  "A",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 5. Simpan user ke DB
	if err := db.Create(&newUser).Error; err != nil {
		return UserDto.UserResponse{}, err
	}

	// 6. Simpan relasi user-role
	for _, role := range validRoles {
		userRole := Entity.UserRole{
			UserID:    userID,
			RoleID:    strconv.Itoa(role.Id), // pastikan RoleID string, kalau bukan sesuaikan
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := db.Create(&userRole).Error; err != nil {
			return UserDto.UserResponse{}, fmt.Errorf("failed to assign role: %v", err)
		}
	}

	// 7. Buat slice RoleResponse untuk response
	var rolesResp []RoleDto.RoleResponse
	for _, r := range validRoles {
		rolesResp = append(rolesResp, RoleDto.RoleResponse{
			Id:   r.Id,
			Name: r.Name,
		})
	}

	// 8. Return response dengan slice Role
	return UserDto.UserResponse{
		Id:       newUser.Id,
		Name:     newUser.Name,
		Email:    newUser.Email,
		IsActive: newUser.IsActive,
		Role:     rolesResp,
	}, nil
}

func GetAllUser() ([]UserDto.UserResponse, error) {
	db := Database.DB

	// 1. Ambil semua user
	var users []Entity.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	// 2. Ambil semua user ID
	var userIDs []string
	for _, u := range users {
		userIDs = append(userIDs, u.Id)
	}

	// 7. Bangun response
	var resp []UserDto.UserResponse
	for _, user := range users {
		resp = append(resp, UserDto.UserResponse{
			Id:       user.Id,
			Name:     user.Name,
			Email:    user.Email,
			IsActive: user.IsActive,
		})
	}

	return resp, nil
}

func GetUserById(id string) (UserDto.UserResponse, error) {

	db := Database.DB

	var user Entity.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return UserDto.UserResponse{}, err
	}
	var userRoles []Entity.UserRole
	if err := db.Where("user_id = ?", id).Find(&userRoles).Error; err != nil {
		return UserDto.UserResponse{}, err
	}
	var roles []Entity.Role
	if err := db.Find(&roles).Error; err != nil {
		return UserDto.UserResponse{}, err
	}
	roleMap := map[string]Entity.Role{}
	for _, r := range roles {
		roleMap[strconv.Itoa(r.Id)] = r
	}
	userRoleMap := map[string][]RoleDto.RoleResponse{}
	for _, ur := range userRoles {
		if role, ok := roleMap[ur.RoleID]; ok {
			userRoleMap[ur.UserID] = append(userRoleMap[ur.UserID], RoleDto.RoleResponse{
				Id:   role.Id,
				Name: role.Name,
			})
		}
	}
	return UserDto.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		IsActive: user.IsActive,
		Role:     userRoleMap[user.Id],
	}, nil
}

func UpdateUser(userID string, userReq *UserDto.UserRequest) (UserDto.UserResponse, error) {
	db := Database.DB

	// 1. Ambil data user
	var user Entity.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		return UserDto.UserResponse{}, err
	}

	// 2. Validasi role yang dikirim
	var validRoles []Entity.Role
	if err := db.Where("id IN ?", userReq.Role).Find(&validRoles).Error; err != nil {
		return UserDto.UserResponse{}, err
	}
	if len(validRoles) != len(userReq.Role) {
		return UserDto.UserResponse{}, errors.New("invalid role")
	}

	// 3. Update data user
	user.Name = userReq.Name
	user.Email = userReq.Email
	user.UpdatedAt = time.Now()

	// Optional: update password jika diberikan
	if userReq.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
		if err != nil {
			return UserDto.UserResponse{}, err
		}
		user.Password = string(hashedPassword)
	}

	// 4. Simpan update user
	if err := db.Save(&user).Error; err != nil {
		return UserDto.UserResponse{}, err
	}

	// 5. Hapus semua role user sebelumnya
	if err := db.Where("user_id = ?", userID).Delete(&Entity.UserRole{}).Error; err != nil {
		return UserDto.UserResponse{}, err
	}

	// 6. Simpan relasi UserRole baru
	for _, role := range validRoles {
		userRole := Entity.UserRole{
			UserID:    userID,
			RoleID:    strconv.Itoa(role.Id), // asumsi RoleID string
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := db.Create(&userRole).Error; err != nil {
			return UserDto.UserResponse{}, err
		}
	}

	// 7. Build role response
	var rolesResp []RoleDto.RoleResponse
	for _, r := range validRoles {
		rolesResp = append(rolesResp, RoleDto.RoleResponse{
			Id:   r.Id,
			Name: r.Name,
		})
	}

	// 8. Return response
	return UserDto.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		IsActive: user.IsActive,
		Role:     rolesResp,
	}, nil
}
