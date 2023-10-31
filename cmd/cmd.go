package cmd

import (
	"errors"
	"flag"
	"strconv"

	wei "github.com/kijimaD/wei/pkg"
)

var (
	NotExistSubCommand = errors.New("wei need subcommand.\nbuild build image\nrec record weight")
)

type CLI struct{}

func New() *CLI {
	return &CLI{}
}

func (c *CLI) Execute() error {
	cnf, err := wei.LoadConfigForYaml()
	if err != nil {
		return err
	}
	args := flag.Args()
	if len(args) <= 1 {
		return NotExistSubCommand
	}
	argSubcmd := args[0]
	if argSubcmd == "build" {
		w := wei.New()
		w.Plot()
	} else if argSubcmd == "rec" {
		argWeight := args[1]
		weight, err := strconv.ParseFloat(argWeight, 64)
		if err != nil {
			return err
		}
		e := wei.NewEntry(cnf, weight)
		err = e.Record()
		if err != nil {
			return err
		}
	}

	return nil
}
