/*
 * tecna
 * https://github.com/chrisenytc/tecna
 *
 * Copyright (c) 2014, Christopher EnyTC
 * Licensed under the MIT license.
 */

package lib

/*
 * Dependencies
 */

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

type App struct {
	Name        string
	Description string
}

type Database struct {
	Url    string
	DbName string
}

type Config struct {
	App      App
	Database Database
}

func GetConfig() (config Config, err error) {
	var goenv string
	if os.Getenv("MARTINI_ENV") != "" {
		goenv = os.Getenv("MARTINI_ENV")
	} else {
		goenv = "development"
	}
	dirbase, err := os.Getwd()
	if err != nil {
		return Config{}, err
	}
	configData, err := ioutil.ReadFile(path.Join(dirbase, "config", goenv+".yml"))
	if err != nil {
		return Config{}, err
	}
	parsedData := Config{}
	err = yaml.Unmarshal(configData, &parsedData)
	if err != nil {
		return Config{}, err
	}
	return parsedData, nil
}
