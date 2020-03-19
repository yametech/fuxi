package controller

import (
	"github.com/yametech/fuxi/pkg/controller/nodeliveness"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, nodeliveness.Add)
}
