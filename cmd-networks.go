package main

import (
	"github.com/fsouza/go-dockerclient"
	"github.com/codegangsta/cli"
	"github.com/mcuadros/go-version"
	"log"
)

func doNetworks(c *cli.Context) {
	client, err := docker.NewClient(c.GlobalString("endpoint"))
	if err != nil {
		log.Fatal(err)
	}

	ver, err := client.Version()
	if version.Compare(ver.Get("ApiVersion"),"1.10", "<") {
		log.Fatal("Network clean only works on Docker 1.10 and newer")
	}
	client.PruneNetworks(docker.PruneNetworksOptions{})
	return
}
