/*
 * tecna
 * https://github.com/chrisenytc/tecna
 *
 * Copyright (c) 2014, Christopher EnyTC
 * Licensed under the MIT license.
 */

package controllers

import (
	"fmt"
	"github.com/chrisenytc/tecna/api/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

func getAll(db *mgo.Database) []models.User {
	var users []models.User
	db.C("users").Find(nil).All(&users)
	return users
}

/*
 * GET /users
 */

func UsersIndex(req *http.Request, r render.Render, db *mgo.Database) {
	r.JSON(200, getAll(db))
}

/*
 * POST /users
 */

func UsersCreate(req *http.Request, r render.Render, db *mgo.Database, user models.User) {
	userData := models.User{Id: bson.NewObjectId(), Name: user.Name, Email: user.Email, CreatedAt: time.Now()}
	err := db.C("users").Insert(userData)
	if err != nil {
		r.JSON(500, map[string]interface{}{"msg": "User not found", "error": err})
	} else {
		r.JSON(200, map[string]interface{}{"msg": "User created successfully", "user": userData})
	}
}

/*
 * PUT /users/:id
 */

func UsersUpdate(params martini.Params, req *http.Request, r render.Render, db *mgo.Database, user models.User) {
	change := bson.M{"$set": bson.M{"name": user.Name, "email": user.Email}}
	err := db.C("users").UpdateId(bson.ObjectIdHex(params["id"]), change)
	if err != nil {
		fmt.Println(err)
		r.JSON(500, map[string]interface{}{"msg": "User not found", "error": err})
	} else {
		r.JSON(200, map[string]interface{}{"msg": "User updated successfully"})
	}
}

/*
 * DELETE /users/:id
 */

func UsersDelete(params martini.Params, req *http.Request, r render.Render, db *mgo.Database) {
	err := db.C("users").RemoveId(bson.ObjectIdHex(params["id"]))
	if err != nil {
		r.JSON(500, map[string]interface{}{"msg": "User not found", "error": err})
	} else {
		r.JSON(200, map[string]interface{}{"msg": "User removed successfully"})
	}
}
