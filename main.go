package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danilobuerger/buderus-km200/api"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	godotenv.Load(".env.local")
	godotenv.Load()

	app := &cli.App{
		Name:  "buderus-km200",
		Usage: "access the km200 via cli",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "hostport",
				Usage:    "hostport of gateway",
				Required: true,
				EnvVars:  []string{"KM200_HOSTPORT"},
			},
			&cli.StringFlag{
				Name:     "gateway-password",
				Usage:    "gateway password (without dashes)",
				Required: true,
				EnvVars:  []string{"KM200_GATEWAY_PASSWORD"},
			},
			&cli.StringFlag{
				Name:     "private-password",
				Usage:    "private password that was created on setup",
				Required: true,
				EnvVars:  []string{"KM200_PRIVATE_PASSWORD"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "get",
				Usage:  "calls an endpoint",
				Action: get,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func get(c *cli.Context) error {
	client := api.NewClient(
		c.String("hostport"),
		c.String("gateway-password"),
		c.String("private-password"),
	)

	res, err := client.Get(c.Args().First())
	if err != nil {
		return err
	}
	fmt.Println(string(res))

	return nil
}
