package main

import (
	"cabal-api/context"
	"os"
	"testing"
)

var a context.App

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
