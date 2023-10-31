package main

import (
	"github.com/kijimaD/wei/cmd"
)

func main() {
	cli := cmd.New()
	err := cli.Execute()
	if err != nil {
		panic(err)
	}
}
