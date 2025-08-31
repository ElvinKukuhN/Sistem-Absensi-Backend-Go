package Entity

import "time"

type User struct {
	Id        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	IsActive  string    `json:"is_active" gorm:"default:A"`
	CreatedAt time.Time `json:"created_at" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updatedAt"`
}

type UserRole struct {
	UserID    string    `json:"user_id" gorm:"column:user_id"`
	RoleID    string    `json:"role_id" gorm:"column:role_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updatedAt"`
	//relation
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Role Role `json:"role,omitempty" gorm:"foreignKey:RoleID"`
}
