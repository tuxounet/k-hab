package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/tuxounet/k-hab/cmd"
	"github.com/tuxounet/k-hab/hab"
)

func main() {
	cmd := &cli.Command{
		EnableShellCompletion: true,
		Name:                  "hab",
		Usage:                 "personal hab cli",
		Version:               "24.8.0",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "quiet",
				Value: false,
				Usage: "be queit",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "provision",
				Usage: "provision the hab",
				Action: func(_ context.Context, ocmd *cli.Command) error {
					hab := hab.NewHab(ocmd.Bool("quiet"))
					return cmd.ProvisionCmd(hab)
				},
			},
			{
				Name:  "up",
				Usage: "create and/or launch the hab",
				Action: func(_ context.Context, ocmd *cli.Command) error {
					hab := hab.NewHab(ocmd.Bool("quiet"))
					return cmd.UpCmd(hab)
				},
			},
			{
				Name:  "shell",
				Usage: "create and/or launch the hab",
				Action: func(_ context.Context, ocmd *cli.Command) error {
					hab := hab.NewHab(ocmd.Bool("quiet"))
					return cmd.ShellCmd(hab)
				},
			},
			{
				Name:  "down",
				Usage: "stop the hab",
				Action: func(_ context.Context, ocmd *cli.Command) error {
					hab := hab.NewHab(ocmd.Bool("quiet"))
					return cmd.DownCmd(hab)
				},
			}, {
				Name:  "rm",
				Usage: "rm the hab",
				Action: func(_ context.Context, ocmd *cli.Command) error {
					hab := hab.NewHab(ocmd.Bool("quiet"))
					return cmd.RmCmd(hab)
				},
			},
			{
				Name:  "unprovision",
				Usage: "unprovision the hab",
				Action: func(_ context.Context, ocmd *cli.Command) error {
					hab := hab.NewHab(ocmd.Bool("quiet"))
					return cmd.UnprovisionCmd(hab)
				},
			},
			{
				Name:  "nuke",
				Usage: "destroy the hab",
				Action: func(_ context.Context, ocmd *cli.Command) error {
					hab := hab.NewHab(ocmd.Bool("quiet"))
					return cmd.NukeCmd(hab)
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
