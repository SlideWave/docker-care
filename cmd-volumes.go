package main

import (
	"log"
	"io/ioutil"
	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
	"github.com/mcuadros/go-version"
	"path/filepath"
	"fmt"
	"os"
)

func listOndiskVolumes(volDir string) (map[string]bool, error) {
	dirs, err := ioutil.ReadDir(volDir)
	if err != nil {
		return nil, err
	}

	ondiskVolumes := map[string]bool{}
	for _, dir := range dirs {
		if dir.IsDir() && len(dir.Name()) == 64 {
			ondiskVolumes[dir.Name()] = true
		}
	}
	return ondiskVolumes, nil
}

func listOnContainerVolumes(client *docker.Client) (map[string]bool, error) {
	containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		return nil, err
	}

	onContainerVolumes := map[string]bool{}
	for _, container := range containers {
		ci, err := client.InspectContainer(container.ID)
		if err != nil {
			return nil, err
		}
		for _, volDir := range ci.Volumes {
			onContainerVolumes[filepath.Base(volDir)] = true
		}
	}
	return onContainerVolumes, nil;
}

func doVolumes(c *cli.Context) {
	client, err := docker.NewClient(c.GlobalString("endpoint"))
	if err != nil {
		log.Fatal(err)
	}

	ver, err := client.Version()
	if version.Compare(ver.Get("ApiVersion"),"1.23", ">=") {
		client.PruneVolumes(docker.PruneVolumesOptions{})
		return
	}
	ondiskVolumes, err := listOndiskVolumes(joinDir(c, "volumes"))
	if err != nil {
		log.Fatal(err)
	}

	onContainerVolumes, err := listOnContainerVolumes(client)
	if err != nil {
		log.Fatal(err)
	}

	for n := range onContainerVolumes {
		if ondiskVolumes[n] {
			delete(ondiskVolumes, n)
		}
	}

	for volumeId := range ondiskVolumes {
		dirs := []string{
			joinDir(c, "volumes", volumeId),
			joinDir(c, "vfs", "dir", volumeId),
		}
		for _, dir := range dirs {
			var err error
			err = os.RemoveAll(dir)
			fmt.Println("removed:", dir)

			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: failed to delete volume id: %s", err, volumeId)
			}
		}
	}
}