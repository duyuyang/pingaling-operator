package controller

import (
	"github.com/yuyangd/pingaling-operator/pkg/controller/pingaling"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, pingaling.Add)
}
