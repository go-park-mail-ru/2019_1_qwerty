package models

import "errors"

var EUserAE = errors.New("User already exists")
var ENE = errors.New("Not Exist")
var EWrongPassword = errors.New("EWrongPassword")

var EUserNE = errors.New("No User")
