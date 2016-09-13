package main

import (
	"errors"
	"os"
	"path"
	"fmt"
	"os/user"
	"github.com/urfave/cli"
)

// NewInstallCommand creates new run cmd for cli
func NewInstallCommand() cli.Command {
	cmd := cli.Command{
		Name:    "install",
		Aliases: []string{"i"},
		Usage:   "install a given manifest",
		Action:  InstallCli,
	}

	return cmd
}

// InstallCli installs a manifest from a given CLI context
func InstallCli(c *cli.Context) error {
	args := c.Args()
	if len(args) == 0 {
		et := "install requires at least one arg: \n\t-a file path or,\n\t-URI\nto a manifest YAML file"
		return errors.New(et)
	}
	path := args.First()
	mnf, e := ParseYaml(path)
	if e != nil {
		return e
	}
	fmt.Printf("Installing %s\n", mnf.Name)
	return Install(mnf, path)
}

// Install installs a given manifest
func Install(mnf *Manifest, origin string) error {
	usr, err := user.Current()
	destPath := usr.HomeDir + InstallPath

  if err != nil {
      return err
  }
	err = os.MkdirAll(destPath, 0700)
	if err != nil {
		return err
	}

	new := path.Join(destPath, mnf.Name)
	fmt.Println(new)
	if _, err := os.Stat(new); os.IsNotExist(err) {
  	err = os.Link(origin, new)
		if err != nil {
			return err
		}
	} else {
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		fmt.Fprintf(os.Stdout, "Already installed %s\n", mnf.Name)
	}
	return nil
}
