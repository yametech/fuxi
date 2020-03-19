package template

type CommonTemplate struct {
	UserId      *uint32 `json:"user_id" form:"user_id" binding:"exists"`
	Namespace   *string `json:"user_id" form:"user_id" binding:"exists"`
	Name        string
	IsNamespace *bool `json:"user_id" form:"user_id" binding:"exists"`
}
