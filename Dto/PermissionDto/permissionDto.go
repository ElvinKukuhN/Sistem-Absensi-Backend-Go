package PermissionDto

type PermissionDto struct {
	Code string `json:"code"`
}

type PermissionResponse struct {
	Id        int    `json:"id"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at" gorm:"column:createdAt"`
	UpdatedAt string `json:"updated_at" gorm:"column:updatedAt"`
}
