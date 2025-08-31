package Entity

import "time"

type Permission struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updatedAt"`
}

type RolePermission struct {
	RoleID       int `json:"role_id" gorm:"column:role_id"`
	PermissionID int `json:"permission_id" gorm:"column:permission_id"`
	//relation
	Role       Role       `json:"role,omitempty" gorm:"foreignKey:RoleID"`
	Permission Permission `json:"permission,omitempty" gorm:"foreignKey:PermissionID"`
}

type MenuPermission struct {
	MenuID       int `json:"menu_id" gorm:"column:menu_id"`
	PermissionID int `json:"permission_id" gorm:"column:permission_id"`
	//relation
	Menu       Menu       `json:"menu,omitempty" gorm:"foreignKey:MenuID"`
	Permission Permission `json:"permission,omitempty" gorm:"foreignKey:PermissionID"`
}
