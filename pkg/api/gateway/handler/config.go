package handler

type UserConfig struct {
	LensVersion       string   `json:"lensVersion"`
	LensTheme         string   `json:"lensTheme"`
	UserName          string   `json:"userName"`
	Token             string   `json:"token"`
	AllowedNamespaces []string `json:"allowedNamespaces"`
	IsClusterAdmin    bool     `json:"isClusterAdmin"`
	ChartEnable       bool     `json:"chartEnable"`
	KubectlAccess     bool     `json:"kubectlAccess"`
}

func newUserConfig(user string, token string, allowedNamespaces []string) *UserConfig {
	isClusterAdmin := false
	if user == "admin" {
		isClusterAdmin = true
	} else {
		allowedNamespaces = []string{"dxp", "dxp2"}
	}
	return &UserConfig{
		LensVersion:       "1.0",
		LensTheme:         "",
		UserName:          user,
		Token:             token,
		AllowedNamespaces: allowedNamespaces,
		IsClusterAdmin:    isClusterAdmin,
		ChartEnable:       true,
		KubectlAccess:     true,
	}
}
