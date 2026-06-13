package roles

type ReqRole struct {
	Name      string `json:"name"`
	Hierarchy int64  `json:"hierarchy"`
}

type ResponseRole struct {
	Name      string `json:"name"`
	Hierarchy int64  `json:"hierarchy"`
}
