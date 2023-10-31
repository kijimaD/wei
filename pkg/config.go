package wei

import (
	"flag"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const defaultConfigPath = ".wei/config.yml"

type config struct {
	CsvPath string `yaml:"csvpath"`
}

func LoadConfigForYaml() (*config, error) {
	homedir, _ := os.UserHomeDir()
	expanded := filepath.Join(homedir, defaultConfigPath)

	var configPath = flag.String("c", expanded, "config path")
	flag.Parse()

	f, err := os.Open(*configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg config
	err = yaml.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}
