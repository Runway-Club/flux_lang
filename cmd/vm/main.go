package main

import (
	"fmt"
	"github.com/Runway-Club/flux_lang/shared"
	"github.com/Runway-Club/flux_lang/vm"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	virtualMachine := vm.NewFluxVirtualMachine()
	app := &cli.App{
		Name:  "flux",
		Usage: "Flux Virtual Machine",
		Action: func(*cli.Context) error {
			fmt.Println("A flexible programming language for Web Development")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Run a Flux project",
				Action: func(c *cli.Context) error {
					// get the project name -i flag
					entryPoint := c.String("i")
					verbose := c.Bool("v")
					fmt.Printf("running %s\n", entryPoint)
					// execute the virtual machine
					result := virtualMachine.Execute(&shared.ExecutionParams{EntryPoint: entryPoint, Verbose: verbose})
					fmt.Printf("Execution completed in %dms\n", result.ElapsedTime)
					if result.Error != "" {
						fmt.Printf("\u001B[31m 8=D Error: %s \u001B[0m \n", result.Error)
					} else {
						fmt.Println("\u001B[34m 8=D Execution with no errors! \u001B[0m")
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "i",
						Value: "main.flux",
						Usage: "entry point file",
					},
					&cli.BoolFlag{
						Name:  "v",
						Value: false,
						Usage: "verbose",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
