package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/tee"
)

const (
	flagAppend          = "append"
	flagIgnoreInterrupt = "ignore-interrupts"
)

func main() {
	app := &cli.App{
		Name:  "tee",
		Usage: "read from standard input and write to standard output and files",
		UsageText: `tee [OPTIONS] [FILE...]

   Copy standard input to each FILE, and also to standard output.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagAppend,
				Aliases: []string{"a"},
				Usage:   "append to the given FILEs, do not overwrite",
			},
			&cli.BoolFlag{
				Name:    flagIgnoreInterrupt,
				Aliases: []string{"i"},
				Usage:   "ignore interrupt signals",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "tee: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments
	for i := 0; i < c.NArg(); i++ {
		params = append(params, gloo.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.Bool(flagAppend) {
		params = append(params, Append)
	}
	if c.Bool(flagIgnoreInterrupt) {
		params = append(params, IgnoreInterrupt)
	}

	// Create and execute the tee command
	cmd := Tee(params...)
	return gloo.Run(cmd)
}
