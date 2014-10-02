/*
 * tecna
 * https://github.com/chrisenytc/tecna
 *
 * Copyright (c) 2014, Christopher EnyTC
 * Licensed under the MIT license.
 */

package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id        bson.ObjectId `json:"_id" bson:"_id"`
	Name      string        `form:"name" json:"name" bson:"name" binding:"required"`
	Email     string        `form:"email" json:"email" bson:"email" binding:"required"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
}
