package main

import (
    "fmt"
    "os"
    "github.com/urfave/cli"
//    "github.com/docker/docker/api/types"
//    "github.com/docker/docker/client"
)

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
                return nil
            },
        },
        {
            Name:    "up",
            Aliases: []string{"u"},
            Usage:   "upload image to cloud",
            Action:  func(c *cli.Context) error {
                fmt.Println("...")
                return nil
            },
        },
    }

    app.Run(os.Args)
}
