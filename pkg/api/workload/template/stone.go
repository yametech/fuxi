package template

import (
	nuwav1 "github.com/yametech/nuwa/api/v1"
)

type StoneRequest struct {
	Model CommonTemplate   `json:"model" from:"model"`
	Spec  nuwav1.StoneSpec `json:"spec" form:"spec"`
}

type StoneResponse struct {
	StoneList nuwav1.StoneList `json:"stone_list`
}
