package goev3

import (
	lua "github.com/yuin/gopher-lua"
	"layeh.com/gopher-luar"
)

func RunLua(script string, ctrl *Controller) error {
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("ev3", luar.New(L, ctrl))
	return L.DoString(script)
}
