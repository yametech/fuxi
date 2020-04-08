package template

type DeploymentRequest struct {
	Namespace *string `json:"namespace" form:"namespace" binding:"exists"`
	Name      string  `json:"name" form:"namespace" binding:"exists"`
}
