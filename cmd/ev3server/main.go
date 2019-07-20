package main

import "github.com/martin42/goev3"

func main() {
	s, err := goev3.NewServer(":8080")
	if err != nil {
		panic(err)
	}
	s.ListenAndServe()
}
