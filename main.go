package main

import (
	"fmt"
	"github.com/urfave/cli"
	"math/rand"
	"os"
	"time"
)

var flips = []string{
	"[ ╯ ' □']╯︵┻━┻)",
	"[ ╯  `^´]╯︵┻━┻)",
	"[ ╯ ´･ω･]╯︵┻━┻)",
}

var puts = []string{
	"┬──┬ノ['-' ノ ]",
	"┬──┬ノ[･ω･ ノ ]",
	"┬──┬ノ['~' ノ ]",
}

var flipCount = len(flips)

func RandomPosition() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(flipCount)
}

func FlipTable() {
	fmt.Println(flips[RandomPosition()])
}

func PutTable() {
	fmt.Println(puts[RandomPosition()])
}

func main() {
	app := cli.NewApp()
	app.Name = "flipz"
	app.Usage = "command-line flip table utility"
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) error {
		FlipTable()
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:    "flip",
			Aliases: []string{"f"},
			Usage:   "random flip table",
			Action: func(c *cli.Context) error {
				FlipTable()
				return nil
			},
		},
		{
			Name:    "put",
			Aliases: []string{"p"},
			Usage:   "random put table",
			Action: func(c *cli.Context) error {
				PutTable()
				return nil
			},
		},
	}
	app.Run(os.Args)
}
