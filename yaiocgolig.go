package main

import (
	"github.com/go-martini/martini"
)

func main() {
	server := martini.Classic()
	server.Get("/", Hello)

	server.Run()
}
