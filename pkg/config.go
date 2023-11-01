package wei

import (
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	CsvPath string `yaml:"csvpath"`
}

func LoadConfigForYaml(configPath string) (*config, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg config
	err = yaml.NewDecoder(f).Decode(&cfg)
	cfg.CsvPath = ExpandHomedir(cfg.CsvPath)
	return &cfg, err
}
