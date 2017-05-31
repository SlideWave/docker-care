package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
)

type images []docker.APIImages
type filter func(image docker.APIImages) bool

func (i images) Filter(f filter) images {
	ret := images{}
	for _, image := range i {
		if f(image) {
			ret = append(ret, image)
		}
	}
	return ret
}

func listImages(client *docker.Client) (images, error) {
	images := images{}
	apiImages, err := client.ListImages(docker.ListImagesOptions{All: false})
	if err != nil {
		return nil, err
	}
	for i := range apiImages {
		images = append(images, apiImages[i])
	}
	return apiImages, nil
}

func filterByName(name string) filter {
	return func(image docker.APIImages) bool {
		for i := range image.RepoTags {
			if strings.HasPrefix(image.RepoTags[i], name) {
				return true
			}
		}
		return false
	}
}
func filterByAge(age int) filter {
	return func(image docker.APIImages) bool {
		age_in_sec := time.Second * time.Duration(age)
		return time.Since(time.Unix(image.Created, 0)) > age_in_sec
	}
}

func doImages(c *cli.Context) {
	client, err := docker.NewClient(c.GlobalString("endpoint"))
	if err != nil {
		log.Fatal(err)
	}

	images, err := listImages(client)
	if err != nil {
		log.Fatal(err)
	}

	filtered_imgs := images.Filter(filterByName(c.String("name"))).Filter(filterByAge(c.Int("age")))

	//total_size := 0
	for i := range filtered_imgs {
		var err error
		err = client.RemoveImage(filtered_imgs[i].ID)
		fmt.Println("removed:", filtered_imgs[i].ID, filtered_imgs[i].RepoTags)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: failed to delete image %v %s", err, filtered_imgs[i].ID, filtered_imgs[i].RepoTags)
		}
	}
}
