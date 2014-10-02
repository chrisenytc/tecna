/*
 * tecna
 * https://github.com/chrisenytc/tecna
 *
 * Copyright (c) 2014, Christopher EnyTC
 * Licensed under the MIT license.
 */

package controllers

import (
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"net/http"
)

/*
 * GET /
 */

func HomeIndex(req *http.Request, r render.Render, db *mgo.Database) {
	r.JSON(200, map[string]interface{}{"msg": "Hello World", "body": req, "query": req.URL.Query().Get("query")})
}
