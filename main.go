package main

import (
	"fmt"
	"github.com/kravetsd/go-clitool/flgs"
	"github.com/kravetsd/go-clitool/secret"
	"os/exec"
)

func main() {
	cloneUrl, path, branch := flgs.GetClone()

	cmd := exec.Command("git", "clone", fmt.Sprintf("-b %s", branch), cloneUrl, path)
	_ = cmd.Run()

	secret.NewEnvironment()
}
