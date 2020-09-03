// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Build mg.Namespace

func (Build) Linux() error {
	env := map[string]string{
		"GOOS":   "linux",
		"GOARCH": "amd64",
	}

	err := sh.RunWith(env, "go", "build", "-o", "agent-linux", "-i", "./cmd/agent")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building agent-linux: %q\n", err)
	}

	return err

}
