package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func parseConfigPath(value string) (string, error) {
	return value, nil
}

func parsePort(value string) (string, error) {
	valueInt, err := strconv.Atoi(value)

	if nil != err {
		return value, errors.New("wrong port, port must be a number")
	}

	if valueInt > 65535 {
		return value, errors.New("wrong port, port must be in the range 0-65535")
	}

	return value, nil
}

func parseAddress(value string) (string, error) {
	var ipRegex = regexp.MustCompile(`^(http://|https://)?(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$`)

	if false == ipRegex.MatchString(value) {
		return value, errors.New("wrong ip address")
	}

	return value, nil
}

func parseConfigLine(value string) (string, string, error) {
	var keyAndValue = regexp.MustCompile(`^([^=]+)=(.+)$`)

	if false == keyAndValue.MatchString(value) {
		return "", "", errors.New("wrong config line")
	}

	tmp := keyAndValue.FindAllString(value, -1)

	return tmp[0], tmp[1], nil
}

func ParseArgs(args []string) (map[string]string, error) {
	values := make(map[string]string)

	index := 1
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-c":
			if index >= len(args) {
				return values, errors.New("-c without value")
			}
			value, err := parseConfigPath(args[index])
			if err != nil {
				return values, err
			}
			values["CONFIG_PATH"] = value
		case "-p":
			if index >= len(args) {
				return values, errors.New("-p without value")
			}
			value, err := parsePort(args[index])
			if err != nil {
				return values, err
			}
			values["PORT"] = value
		case "-a":
			if index >= len(args) {
				return values, errors.New("-a without value")
			}
			value, err := parseAddress(args[index])
			if err != nil {
				return values, err
			}
			values["ADDRESS"] = value
		case "-sc":
			if index >= len(args) {
				return values, errors.New("-sc without value")
			}
			keyA, valueA, err := parseConfigLine(args[index])
			if err != nil {
				return values, err
			}

			values[strings.TrimSpace(keyA)] = strings.TrimSpace(valueA)
		}
		index++
	}

	return values, nil
}
