package system

type CasbinPolicy struct {
	PType  string `json:"p_type" form:"p_type"`
	RoleID string `json:"role_id" form:"role_id"`
	Path   string `json:"path" form:"path"`
	Method string `json:"method" form:"method"`
	Desc   string `json:"desc" form:"desc"`
}
type CasbinPolicyList struct {
	Items []CasbinPolicy `json:"items"`
	Total int            `json:"total"`
}
