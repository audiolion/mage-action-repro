// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Build mg.Namespace

func (Build) Windows() error {
	env := map[string]string{
		"GOOS":        "windows",
		"GOARCH":      "386",
		"CGO_ENABLED": "1",
	}

	err := sh.RunWith(env, "go", "build", "-o", "agent-linux", "-i", "./cmd/agent")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building agent-linux: %q\n", err)
	}

	return err

}
