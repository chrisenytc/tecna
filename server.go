/*
 * tecna
 * https://github.com/chrisenytc/tecna
 *
 * Copyright (c) 2014, Christopher EnyTC
 * Licensed under the MIT license.
 */

package main

/*
 * Dependencies
 */

import (
	"fmt"
	"github.com/chrisenytc/tecna/lib"
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

/*
 * Main application
 */

func main() {
	config, err := lib.GetConfig()
	if err != nil {
		panic(err)
	}
	session, err := mgo.Dial(config.Database.Url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	DbSession := &lib.DatabaseSession{session}
	server := lib.CreateServer(DbSession)
	fmt.Printf("[mongodb] mongodb connected successfully on %s\n", config.Database.Url)
	server.Run()
}
