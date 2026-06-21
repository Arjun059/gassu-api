package users

type ReqUser struct {
	Name   string `json:"name"`
	RoleID int64  `json:"role_id"`
}

type CreateUserResponseDTO struct {
	Token  string `json:"token"`
	UserID int64  `json:"userId"`
}
