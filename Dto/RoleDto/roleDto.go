package RoleDto

type RoleDto struct {
	Name string `json:"name"`
}

type RoleResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
