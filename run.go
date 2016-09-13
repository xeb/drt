package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os/exec"
	"os"
	"os/user"
	"strings"
	"github.com/urfave/cli"
)


// NewRunCommand creates new run cmd for cli
func NewRunCommand() cli.Command {
	cmd := cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "run a given manifest",
		Action:  RunCli,
	}

	return cmd
}

// RunCli executes run from CLI context
func RunCli(c *cli.Context) error {
	args := c.Args()
	if len(args) == 0 {
		et := "run requires at least one arg: \n\t-a file path,\n\t-URI or\n\t-registered alias\nto a manifest YAML file"
		return errors.New(et)
	}
	path := args.First()
	params := []string{}
	if len(args) > 1 {
		params = args[1:]
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
  	return RunAlias(path, params)
	}

	mnf, e := ParseYaml(path)
	if e != nil {
		return e
	}
	return Run(mnf, params)
}

// RunAlias executes run with a given manifest alias in ~/.drt and args list
func RunAlias(mnfName string, args []string) error {
	usr, err := user.Current()
	destPath := usr.HomeDir + InstallPath + "/" + mnfName
  if err != nil {
      return err
  }
	mnf, e := ParseYaml(destPath)
	if e != nil {
		return e
	}
	return Run(mnf, args)
}

// Run executes run with a given manifest and args list
func Run(mnf *Manifest, args []string) error {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}

	params := BuildParameters(dir, mnf.WorkingDir, args)
	_, err = exec.LookPath("docker")
	if err != nil {
		return errors.New("docker not found in PATH")
	}

	rp := []string{
		"run",
		"--rm",
		"-v", fmt.Sprintf("%s:%s", dir, mnf.WorkingDir),
	}

	if mnf.Params != "" {
 		extraParams := strings.Split(mnf.Params, " ")
		rp = append(rp, extraParams...)
	}
	rp = append(rp, mnf.Image)

	if mnf.Cmd != "" {
		rp = append(rp, mnf.Cmd)
	}

	params = append(rp, params...)
	fmt.Printf("docker %s\n", strings.Join(params, " "))

	cmd := exec.Command("docker", params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// BuildParameters creates a string of parameters for docker
func BuildParameters(hostDir string, containerDir string, args []string) []string {
	files, _ := ioutil.ReadDir(hostDir)

	if len(containerDir) > 0 && containerDir[len(containerDir)-1] != '/' {
		containerDir = containerDir + "/"
	}

	for i, e := range(args) {
		if(len(e) == 0) {
			continue
		}

		// Determine if this is a relative path
		// that we need to add the mounted PWD volume path to
		relpath := false
		if(e[0] != '/' && // not absolute path
			strings.Contains(e, "/")) { // but it is a path...
			relpath = true
		} else {
			for _, f := range files {
				if f.Name() == e { // a file is in current pwd
					relpath = true
				}
			}
		}

		if relpath {
			args[i] = containerDir + e
		}
	}
	return args
}
