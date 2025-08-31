package Entity

type Menu struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at" gorm:"column:createdAt"`
	UpdatedAt string `json:"updated_at" gorm:"column:updatedAt"`
}
