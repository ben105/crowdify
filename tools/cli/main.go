package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:    "track",
				Aliases: []string{"t"},
				Usage:   "options for Spotify tracks",
				Commands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new track",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							trackName := cmd.Args().First()
							fmt.Println("new track:", trackName)
							addTrack(trackName)
							return nil
						},
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
