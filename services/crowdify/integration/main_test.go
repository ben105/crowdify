package integration_tests

import (
	"os"
	"testing"

	"github.com/ben105/crowdify/packages/emulator"
)

func TestMain(m *testing.M) {
	emulator.Load()
	emulator.Emulate()
	code := m.Run()
	emulator.Destroy()
	os.Exit(code)
}
