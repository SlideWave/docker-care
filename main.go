package main

import (
	"os"
	"path/filepath"
	"github.com/codegangsta/cli"
)

const Version string = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "docker-care"
	app.Version = Version
	app.Usage = "A tool for the tending and feeding of dockerd"
	app.Author = "SlideWave, LLC"
	app.Email = "info@slidewave.com"
	app.Commands = Commands
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "directory, d",
			Value: "/var/lib/docker",
			Usage: "Specifies docker path",
		},
		cli.StringFlag{
			Name: "endpoint, e",
			Value: "unix:///var/run/docker.sock",
			Usage: "Specifies a docker endpoint",
		},
	}
	app.Run(os.Args)
}

func joinDir(c *cli.Context, dirs ...string) string {
	return filepath.Join(c.GlobalString("directory"), filepath.Join(dirs...))
}
