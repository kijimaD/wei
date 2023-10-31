package cmd

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
	"strconv"

	wei "github.com/kijimaD/wei/pkg"
)

var (
	NotExistSubCommand = errors.New("wei need subcommand.\nbuild build image\nrec record weight")
)

const defaultConfigPath = ".wei/config.yml"

type CLI struct{}

func New() *CLI {
	return &CLI{}
}

func (c *CLI) Execute() error {
	homedir, _ := os.UserHomeDir()
	expanded := filepath.Join(homedir, defaultConfigPath)
	var configPath = flag.String("c", expanded, "config path")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		return NotExistSubCommand
	}
	argSubcmd := args[0]
	if argSubcmd == "build" {
		w := wei.New()
		w.Plot()
	} else if argSubcmd == "rec" {
		cnf, err := wei.LoadConfigForYaml(*configPath)
		if err != nil {
			return err
		}
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
