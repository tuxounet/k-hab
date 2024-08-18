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
		Name:                  "k-hab",
		Usage:                 "k-hab cli",
		Version:               "24.8.1",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "quiet",
				Value: false,
				Usage: "be quiet",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "provision",
				Usage: "provision the hab",
				Action: func(_ context.Context, ocmd *cli.Command) error {
					return cmd.ProvisionCmd(hab.NewHab(ocmd.Bool("quiet")))
				},
			},
			{
				Name:  "up",
				Usage: "create and/or launch the hab",
				Action: func(_ context.Context, ocmd *cli.Command) error {
					return cmd.UpCmd(hab.NewHab(ocmd.Bool("quiet")))
				},
			},
			{
				Name:  "shell",
				Usage: "create and/or launch the hab",
				Action: func(_ context.Context, ocmd *cli.Command) error {
					return cmd.ShellCmd(hab.NewHab(ocmd.Bool("quiet")))
				},
			},
			{
				Name:  "down",
				Usage: "stop the hab",
				Action: func(_ context.Context, ocmd *cli.Command) error {
					return cmd.DownCmd(hab.NewHab(ocmd.Bool("quiet")))
				},
			}, {
				Name:  "rm",
				Usage: "rm the hab",
				Action: func(_ context.Context, ocmd *cli.Command) error {
					return cmd.RmCmd(hab.NewHab(ocmd.Bool("quiet")))
				},
			},
			{
				Name:  "unprovision",
				Usage: "unprovision the hab",
				Action: func(_ context.Context, ocmd *cli.Command) error {
					return cmd.UnprovisionCmd(hab.NewHab(ocmd.Bool("quiet")))
				},
			},
			{
				Name:  "nuke",
				Usage: "destroy the hab",
				Action: func(_ context.Context, ocmd *cli.Command) error {
					return cmd.NukeCmd(hab.NewHab(ocmd.Bool("quiet")))
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
