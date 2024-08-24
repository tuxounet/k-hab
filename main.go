package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	habContext "github.com/tuxounet/k-hab/context"
)

var version = "DEVELOPEMENT"

type Author struct {
	Name    string
	Contact string
}

func main() {

	cmd := &cli.Command{
		Name:                  "k-hab",
		Version:               version,
		Authors:               []any{"Christophe TIRAOUI <github:tuxounet>"},
		Description:           "a single executable configuring and executing one or more containers, whose network interactions are controlled",
		EnableShellCompletion: true,
		HideHelp:              false,
		HideVersion:           false,
		Flags: []cli.Flag{

			&cli.BoolFlag{
				Name:  "quiet",
				Value: false,
				Usage: "be quiet",
			},
			&cli.StringFlag{
				Name:  "loglevel",
				Value: "INFO",
				Usage: "set minimal level to produce log (DEBUG, INFO, WARN, ERROR, FATAL)",
				Validator: func(s string) error {
					switch s {
					case "DEBUG", "INFO", "WARN", "ERROR", "FATAL":
						return nil
					default:
						return errors.New("loglevel must be one of DEBUG, INFO, WARN, ERROR, FATAL")
					}
				},
			},
		},
		Commands: []*cli.Command{
			buildCommand("provision", "provision the hab", habContext.ProvisionVerb),
			buildCommand("up", "create and/or launch the hab", habContext.UpVerb),
			buildCommand("shell", "create and/or launch the hab", habContext.ShellVerb),
			buildCommand("down", "stop the hab", habContext.DownVerb),
			buildCommand("rm", "rm the hab", habContext.RmVerb),
			buildCommand("unprovision", "unprovision the hab", habContext.UnprovisionVerb),
			buildCommand("nuke", "destroy the hab", habContext.NukeVerb),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}

func buildCommand(name string, usage string, verb habContext.HabVerbs) *cli.Command {
	return &cli.Command{
		Name:  name,
		Usage: usage,
		Action: func(ctx context.Context, ocmd *cli.Command) error {
			habCtx := habContext.NewHabContext(ctx)
			err := habCtx.ParseCli(ocmd)
			if err != nil {
				return err
			}
			err = habCtx.Init()
			if err != nil {
				return err
			}

			switch verb {
			case habContext.ProvisionVerb:
				return habCtx.Provision()
			case habContext.UpVerb:
				return habCtx.Start()
			case habContext.ShellVerb:
				return habCtx.Shell()
			case habContext.DownVerb:
				return habCtx.Stop()
			case habContext.RmVerb:
				return habCtx.Rm()
			case habContext.UnprovisionVerb:
				return habCtx.Unprovision()
			case habContext.NukeVerb:
				return habCtx.Nuke()
			default:
				return errors.New("unknown verb " + string(verb))
			}
		},
	}
}
