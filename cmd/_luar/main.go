package main

import (
	"fmt"
	"time"

	"github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

var script = `
	t = Thing()
	t.Name = "fridge"
	t.Purpose = "cooling"
	did = t:Do(5)	
	print("did it", did, "times")
	print("make other thing ...")
	ot = t:MakeOtherThing("car", "driving")	
	ot:Do(1)
`

func main() {
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("Thing", luar.NewType(L, Thing{}))

	fmt.Println("run script ...")
	start := time.Now()
	if err := L.DoString(script); err != nil {
		panic(err)
	}
	fmt.Println("run script took", time.Since(start))
}

type Thing struct {
	Name    string
	Purpose string
}

func (t *Thing) Do(times int) int {
	for i := 0; i < times; i++ {
		fmt.Println(t.Name, "is", t.Purpose)
	}
	return times
}

func (t *Thing) MakeOtherThing(name, purpose string) *Thing {
	return &Thing{
		Name:    name,
		Purpose: purpose,
	}
}
