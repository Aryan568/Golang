package main

import "github.com/google/wire"

func Initialize() User {
	wire.Build(newUser, uName)
	return User{}
}