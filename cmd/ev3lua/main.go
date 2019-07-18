package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/martin42/goev3"
)

func main() {
	scriptFile := flag.String("s", "", "script file to execute")
	flag.Parse()
	bs, err := ioutil.ReadFile(*scriptFile)
	if err != nil {
		fmt.Println("ERROR: read-script:", err)
		os.Exit(1)
	}

	ctrl, err := goev3.NewController()
	if err != nil {
		fmt.Println("ERROR: new-controller:", err)
		os.Exit(2)
	}
	err = goev3.RunLua(string(bs), ctrl)
	if err != nil {
		fmt.Println("ERROR: run-lua:", err)
		os.Exit(3)
	}
}
