package cmd

import (
	"errors"
	"flag"
	"fmt"

	wei "github.com/kijimaD/wei/pkg"
)

var (
	NotExistSubCommand = errors.New("wei need subcommand.\nbuild build image\nrec record weight")
)

type CLI struct{}

func New() *CLI {
	return &CLI{}
}

func (c *CLI) Execute(args []string) error {
	flag.Parse()

	if len(args) <= 1 {
		return NotExistSubCommand
	}

	if args[1] == "build" {
		w := wei.New()
		w.Plot()
	} else if args[1] == "rec" {
		// TODO: 未実装
		fmt.Println("rec")
	}

	return nil
}
