package config

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	os.Exit(retCode)
}
