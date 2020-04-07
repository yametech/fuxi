package template

type CommonTemplate struct {
	UserID    *uint32 `json:"user_id" form:"user_id" binding:"exists"`
	Namespace *string `json:"namespace" form:"namespace" binding:"exists"`
	Name      string
	IsAdmin   *bool `json:"isadmin" form:"isadmin" binding:"exists"`
}
