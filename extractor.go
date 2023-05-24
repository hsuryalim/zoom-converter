package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	var extractionLink string

	app := cli.NewApp()
	app.Name = "zoomlify"
	app.Usage = "extract zoom meeting link"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "extract",
			Destination: &extractionLink,
			Usage:       "Set link",
			Aliases:     []string{"e"},
			Action: func(ctx *cli.Context, link string) error {
				printOut(link)
				os.Exit(0)
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
	}
}

func printOut(link string) {
	re := regexp.MustCompile(`https.*?\n`)
	value := re.Find([]byte(link))
	str := string(value)
	fmt.Println(strings.Trim(str, "\n"))
}
