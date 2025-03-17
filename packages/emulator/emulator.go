package emulator

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

var ctx = context.Background()
var compose tc.ComposeStack

func Load() {
	fmt.Println("Starting environment...")
	identifier := tc.StackIdentifier("integration_tests")

	_, filename, _, _ := runtime.Caller(0)
	moduleDir := filepath.Dir(filename)

	var err error
	compose, err = tc.NewDockerComposeWith(tc.WithStackFiles(filepath.Join(moduleDir, "../../emulator/compose.yaml")), identifier)
	if err != nil {
		log.Fatalf("Failed to start environment: %v\n", err)
	}
}

func Emulate() {
	err := compose.Up(ctx, tc.Wait(true))
	if err != nil {
		log.Fatalf("Failed to start environment: %v\n", err)
	}
}

func Destroy() {
	if compose == nil {
		return
	}
	err := compose.Down(ctx, tc.RemoveOrphans(true), tc.RemoveVolumes(true))
	if err != nil {
		log.Fatalf("Failed to stop environment: %v\n", err)
	}
}
