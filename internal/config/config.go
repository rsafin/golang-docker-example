package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	App struct{
		Bot struct {
			Key string `yaml:"key"`
		} `yaml:"bot"`
	} `yaml:"app"`
}

func NewConfig(filePath string) (*Config, error)  {
	config := &Config{}

	file, err := os.Open("config.yml")
	if err != nil {
		return nil, err

	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
