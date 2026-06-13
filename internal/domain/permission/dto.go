package permission

type ReqPermission struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}
