package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

const (
	InstallPath string = "/.drt"
)

// Manifest denotes the drt properties from the YAML file
type Manifest struct {
	Name           string
	Description    string
	Image          string
	WorkingDir     string	`yaml:"workingDir,omitempty"`
	Params      	 string
	Cmd        		 string
}

// ParseYaml parses []bytes into a YAML doc
func ParseYaml(path string) (*Manifest, error) {
	manifest := &Manifest{}
	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}

	e = yaml.Unmarshal(bytes, manifest)
	if e != nil {
		return nil, e
	}
	return manifest, nil
}
