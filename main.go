package main

import (
	"fmt"
	"os"

	"github.com/mmzou/geektime-dl/cli/cmds"
	"github.com/mmzou/geektime-dl/config"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {

	err := config.Instance.Init()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	app := cmds.NewApp()
	app.Commands = []cli.Command{
		cmds.NewLoginCommand(),
		cmds.NewProductCommand(),
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}