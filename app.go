package main

// For more info check out:
// https://www.iron.io/an-easier-way-to-create-tiny-golang-docker-images/

import (
    "fmt"
    "os"
    "github.com/urfave/cli"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
)

func buildDockerImage() {
    // https://godoc.org/github.com/docker/docker/client#Client.ContainerExecStart
    //cli.ContainerCreate()
    //cli.ContainerExecStart()
}

func buildBinary() {
    //docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app golang:1.4.2 sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o hello'
}

func main() {
    app := cli.NewApp()

    app.Name = "gwf"
    app.Usage = "manage a Go web application"

    /*
    app.Action = func(c *cli.Context) error {
        return nil
    }
    */

    app.Commands = []cli.Command{
        {
            Name:    "build",
            Aliases: []string{"b"},
            Usage:   "build a docker image",
            Action:  func(c *cli.Context) error {
                fmt.Println("image built.")
                buildDockerImage()
                return nil
            },
        },
        {
            Name:    "up",
            Aliases: []string{"u"},
            Usage:   "upload image to cloud",
            Action:  func(c *cli.Context) error {
                fmt.Println("uploading image...")
                buildBinary()
                return nil
            },
        },
    }

    app.Run(os.Args)
}
