package data

import "errors"

var ErrLoginTaken = errors.New("this login is already taken")
var ErrLoggedInAlready = errors.New("user logged in already")
var ErrCredentialsDontMatch = errors.New("credentials do not match")
