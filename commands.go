package main

import "github.com/codegangsta/cli"

var Commands = []cli.Command {
	commandVolume,
	commandImage,
	commandNetwork,
}

var commandVolume = cli.Command {
	Name: "volume",
	ShortName: "v",
	Usage: "Removed orphaned volumes from the host",
	Action: doVolumes,
	Flags: []cli.Flag {
		cli.BoolFlag {
			Name: "force, f",
			Usage: "Force orphaned volumes to be removed",
		},
	},
}

var commandImage = cli.Command {
	Name: "image",
	ShortName: "i",
	Usage: "Removes orphaned images from the host",
	Action: doImages,
	Flags: []cli.Flag {
		cli.BoolFlag {
			Name: "force, f",
			Usage: "Force orphaned volumes to be removed",
		},
		cli.StringFlag {
			Name: "name",
			Usage: "Delete image specified by name",
		},
		cli.IntFlag {
			Name: "age, a",
			Usage: "Delete images whose age Created time is older than specified age in seconds",
		},
	},
}

var commandNetwork = cli.Command {
	Name: "network",
	ShortName: "n",
	Usage: "Removes empty networks from the host",
	Action: doNetworks,
	Flags: []cli.Flag {
		cli.BoolFlag{
			Name: "force, f",
			Usage: "Force empty networks to be removed",
		},
	},
}