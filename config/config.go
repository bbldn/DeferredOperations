package config

import (
	"errors"
	"io/ioutil"
	"strings"
)

type Config struct {
	Values map[string]string
}

func (c *Config) Load(config map[string]string) error {
	return c.load(config)
}

func (c *Config) load(config map[string]string) error {
	data, err := ioutil.ReadFile("/etc/deferred-operation/config.conf")
	if err != nil {
		return errors.New("error open config file")
	}

	if nil == c.Values {
		c.Values = make(map[string]string)
	}

	dataString := strings.TrimSpace(string(data))

	if len(dataString) > 0 {
		lines := strings.Split(dataString, "\n")
		for _, value := range lines {
			//KEY=VALUE
			values := strings.Split(value, "=")
			if 2 != len(values) {
				return errors.New("error config parse")
			}
			c.Values[strings.TrimSpace(values[0])] = strings.TrimSpace(values[1])
		}
	}

	for key, value := range config {
		c.Values[key] = value
	}

	return c.validate()
}

func (c *Config) validate() error {
	_, exists := c.Values["APP_PATH"]
	if true != exists {
		return errors.New("`APP_PATH` not found in config")
	}

	_, exists = c.Values["PORT"]
	if true != exists {
		return errors.New("`PORT` not found in config")
	}

	_, exists = c.Values["ADDRESS"]
	if true != exists {
		return errors.New("`ADDRESS` not found in config")
	}

	return nil
}
