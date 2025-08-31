package Entity

import "time"

type Role struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updatedAt"`
}
