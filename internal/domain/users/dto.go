package users

type ReqUser struct {
	Name   string `json:"name"`
	RoleID int64  `json:"role_id"`
}
