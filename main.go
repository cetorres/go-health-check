package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "Website Health Checker",
		Usage: "A small tool to check if a website is running and reachable or not.",
		Version: "1.0",
		Authors: []*cli.Author{
			{Name:"Carlos E. Torres", Email: "cetorres@cetorres.com"},
		},
		Copyright: "2023 Carlos E. Torres",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "domain",
				Aliases: []string{"d"},
				Usage: "domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name: "port",
				Aliases: []string{"p"},
				Usage: "port number to check",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error{
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}