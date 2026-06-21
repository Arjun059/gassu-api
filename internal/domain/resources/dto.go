package resources

type ReqPermission struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}
