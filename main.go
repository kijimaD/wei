package main

import (
	"os"

	"github.com/kijimaD/wei/cmd"
)

func main() {
	cli := cmd.New()
	err := cli.Execute(os.Args)
	if err != nil {
		panic(err)
	}
}
