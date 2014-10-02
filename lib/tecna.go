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
	"github.com/chrisenytc/tecna/api/controllers"
	"github.com/chrisenytc/tecna/api/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

type DatabaseSession struct {
	*mgo.Session
}

func (session *DatabaseSession) Database() martini.Handler {
	config, err := GetConfig()
	if err != nil {
		panic(err)
	}
	return func(context martini.Context) {
		s := session.Clone()
		context.Map(s.DB(config.Database.DbName))
		defer s.Close()
		context.Next()
	}
}

type Server *martini.ClassicMartini

func CreateServer(session *DatabaseSession) Server {
	m := Server(martini.Classic())
	m.Use(render.Renderer(render.Options{IndentJSON: true}))
	m.Use(session.Database())

	m.Get("/", controllers.HomeIndex)

	m.Post("/users", binding.Form(models.User{}), controllers.UsersCreate)
	m.Put("/users/:id", binding.Form(models.User{}), controllers.UsersUpdate)
	m.Delete("/users/:id", controllers.UsersDelete)
	m.Get("/users", controllers.UsersIndex)

	return m
}
