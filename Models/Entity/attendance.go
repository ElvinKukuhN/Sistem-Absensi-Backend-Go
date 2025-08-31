package Entity

import (
	"time"
)

type Attendance struct {
	ID        string    `json:"id" gorm:"primary_key"`
	UserID    string    `json:"user_id" gorm:"column:user_id"`
	CheckIn   time.Time `json:"check_in" gorm:"column:check_in"`
	CheckOut  time.Time `json:"check_out" gorm:"column:check_out"`
	Location  string    `json:"location"`
	Status    string    `json:"status" gorm:"comment:present, absent, or on leave"`
	CreatedAt time.Time `json:"created_at" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updatedAt"`
	// Relation
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
