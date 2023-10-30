package cmd

import (
	"errors"

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
	if len(args) <= 1 {
		return NotExistSubCommand
	}

	if args[1] == "build" {
		w := wei.New()
		w.Plot()
	} else if args[1] == "rec" {
		cnf, err := wei.LoadConfigForYaml()
		if err != nil {
			return err
		}
		e := wei.NewEntry(cnf, 55.55)
		err = e.Record()
		if err != nil {
			return err
		}
	}

	return nil
}
