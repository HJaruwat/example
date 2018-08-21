package main

import (
	"cabal-api/context"
)

// App instance of application
var App context.App

func main() {
	setup()
	Run()
}

func setup() {
	App = context.App{}
	App.InitializeDatabase()
	App.InitializeRoute()

}

// Run is start instance and listener
func Run() {
	App.Run(":90")
}
