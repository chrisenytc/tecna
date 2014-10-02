/*
 * tecna
 * https://github.com/chrisenytc/tecna
 *
 * Copyright (c) 2014, Christopher EnyTC
 * Licensed under the MIT license.
 */

package main

import (
	"github.com/chrisenytc/tecna/lib"
	"testing"
)

func TestGetConfig(t *testing.T) {
	config, err := lib.GetConfig()
	if err != nil {
		panic(err)
	}
	testConfig := lib.Config{
		lib.App{
			Name:        "Tecna - Test",
			Description: "The powerful stack for making awesome apis with go lang",
		},
		lib.Database{
			Url:    "mongodb://localhost:27017",
			DbName: "tecnadb-test",
		},
	}
	if config != testConfig {
		t.Errorf("'%s' is different of '%s'", config, testConfig)
	}
}

func TestGetConfigError(t *testing.T) {
	config, err := lib.GetConfig()
	if err != nil {
		panic(err)
	}
	testConfig := lib.Config{
		lib.App{
			Name:        "Tecna - Error",
			Description: "The powerful stack for making awesome apis with go lang",
		},
		lib.Database{
			Url:    "mongodb://localhost:27017",
			DbName: "tecnadb-test",
		},
	}
	if config == testConfig {
		t.Errorf("'%s' is equal to '%s'", config, testConfig)
	}
}
