package handler

import "encoding/json"

type userConfig struct {
	LensVersion       string   `json:"lensVersion"`
	LensTheme         string   `json:"lensTheme"`
	UserName          string   `json:"userName"`
	Token             string   `json:"token"`
	AllowedNamespaces []string `json:"allowedNamespaces"`
	IsClusterAdmin    bool     `json:"isClusterAdmin"`
	ChartEnable       bool     `json:"chartEnable"`
	KubectlAccess     bool     `json:"kubectlAccess"`
	DefaultNamespace  string   `json:"defaultNamespace"`
}

func (uc *userConfig) String() string {
	bytesData, _ := json.Marshal(uc)
	return string(bytesData)
}

func newUserConfig(user string, token string, allowedNamespaces []string, defaultNamespace string) *userConfig {
	isClusterAdmin := false
	if user == "admin" {
		isClusterAdmin = true
	}
	return &userConfig{
		LensVersion:       "1.0",
		LensTheme:         "",
		UserName:          user,
		Token:             token,
		AllowedNamespaces: allowedNamespaces,
		IsClusterAdmin:    isClusterAdmin,
		ChartEnable:       true,
		KubectlAccess:     true,
		DefaultNamespace:  defaultNamespace,
	}
}
