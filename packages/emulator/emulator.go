package emulator

import (
	"context"
	"fmt"
	"log"
	"time"

	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

var ctx = context.Background()
var compose tc.ComposeStack

func init() {
	fmt.Println("Starting environment...")
	identifier := tc.StackIdentifier("integration_tests")

	var err error
	compose, err = tc.NewDockerComposeWith(tc.WithStackFiles("../../emulator/compose.yaml"), identifier)
	if err != nil {
		log.Fatalf("Failed to start environment: %v\n", err)
	}
}

func Emulate() {
	err := compose.Up(ctx, tc.Wait(true))

	// Wait for our migrations to run, otherwise we try to connect to a keyspace that doesn't exist.
	time.Sleep(30 * time.Second)

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
