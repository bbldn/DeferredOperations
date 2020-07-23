package config

import (
	"errors"
	"fmt"
	GoConfigParser "github.com/bigkevmcd/go-configparser"
)

type Config struct {
	Values *GoConfigParser.ConfigParser
}

func (c *Config) Load() error {
	data, err := GoConfigParser.NewConfigParserFromFile("config/config.cfg")
	if err != nil {
		return errors.New("error open config file")
	}

	c.Values = data

	return c.validate()
}

func (c *Config) validate() error {
	if 0 == len(c.Values.Sections()) {
		return errors.New("config file is empty")
	}

	if false == c.Values.HasSection("DEFAULTS") {
		return errors.New("`configuration file does not contain `DEFAULTS` section")
	}

	exists, _ := c.Values.HasOption("DEFAULTS", "PORT")
	if false == exists {
		return errors.New("`PORT` not found in config")
	}

	exists, _ = c.Values.HasOption("DEFAULTS", "ADDRESS")
	if false == exists {
		return errors.New("`ADDRESS` not found in config")
	}

	for _, section := range c.Values.Sections() {
		if "DEFAULTS" == section {
			continue
		}

		exists, _ = c.Values.HasOption(section, "APP_PATH")
		if false == exists {
			return errors.New(fmt.Sprintf("`%s` section not contain `APP_PATH`", section))
		}
	}

	return nil
}
