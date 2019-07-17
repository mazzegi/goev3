package main

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

var script = `
	f = thing.new("fridge", "cooling")
	print(f.name(f), f:what_about())	

	v = thing.new("cleaner", "vacuum")
	print(v:name(), v:what_about())	

	v:set_what_about("wet")
	print(v:name(), v:what_about())	
`

var script2 = `
	motorA = "outA"

	ev3_move_to_rel_pos{
		motor = motorA,
		pos = 50,
		speed = 10,
		stop = "brake"
	}
`

func main() {
	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("ev3_move_to_rel_pos", L.NewFunction(LMoveToRelPos))
	err := L.DoString(script2)
	if err != nil {
		panic(err)
	}

	// mtt := L.NewTypeMetatable(luaTypeThing)
	// L.SetGlobal(luaTypeThing, mtt)
	// L.SetField(mtt, "new", L.NewFunction(LNewThing))
	// L.SetField(mtt, "__index", L.SetFuncs(L.NewTable(), thingLMethods()))
	// // L.SetField(mtt, "what_about", L.NewFunction(LThingWhatAbout))
	// // L.SetField(mtt, "name", L.NewFunction(LThingName))
	// err := L.DoString(script)
	// if err != nil {
	// 	panic(err)
	// }
}

type Thing struct {
	name      string
	whatAbout string
}

const luaTypeThing = `thing`

func (t Thing) WhatAbout() string {
	return t.whatAbout
}

func (t Thing) Name() string {
	return t.name
}

func newThing(name, whatAbout string) *Thing {
	return &Thing{
		name:      name,
		whatAbout: whatAbout,
	}
}

func thingLMethods() map[string]lua.LGFunction {
	return map[string]lua.LGFunction{
		"name":           LThingName,
		"what_about":     LThingWhatAbout,
		"set_what_about": LThingSetWhatAbout,
	}
}

func LNewThing(L *lua.LState) int {
	name := L.ToString(1)
	whatAbout := L.ToString(2)
	t := newThing(name, whatAbout)
	ud := L.NewUserData()
	ud.Value = t
	L.SetMetatable(ud, L.GetTypeMetatable(luaTypeThing))
	L.Push(ud)
	return 1
}

func LThingWhatAbout(L *lua.LState) int {
	ud := L.CheckUserData(1)
	t, ok := ud.Value.(*Thing)
	if !ok {
		return 0
	}
	L.Push(lua.LString(t.WhatAbout()))
	return 1
}

func LThingSetWhatAbout(L *lua.LState) int {
	ud := L.CheckUserData(1)
	t, ok := ud.Value.(*Thing)
	if !ok {
		return 0
	}
	wa := L.CheckString(2)
	t.whatAbout = wa
	return 0
}

func LThingName(L *lua.LState) int {
	ud := L.CheckUserData(1)
	t, ok := ud.Value.(*Thing)
	if !ok {
		return 0
	}
	L.Push(lua.LString(t.Name()))
	return 1
}

func LMoveToRelPos(L *lua.LState) int {
	arg := L.ToTable(1)
	motor := ""
	pos := 0
	speed := 0
	stop := ""
	arg.ForEach(func(k lua.LValue, v lua.LValue) {
		switch k.String() {
		case "motor":
			motor = v.String()
		case "pos":
			pos = int(lua.LVAsNumber(v))
		case "speed":
			speed = int(lua.LVAsNumber(v))
		case "stop":
			stop = v.String()
		}
	})
	fmt.Println("motor:", motor, "pos:", pos, "speed:", speed, "stop:", stop)
	return 0
}
