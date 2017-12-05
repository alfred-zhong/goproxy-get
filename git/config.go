package git

import (
	"os"
	"os/exec"
)

// Config set the git config key with value. If global is true, git config will
// be set "--global".
//
// It's the same as "git config [--global] key value"
func Config(key, value string, global bool) error {
	g := ""
	if global {
		g = "--global"
	}

	cmd := exec.Command("git", "config", g, key, value)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Unset delete the git config. If global is true, git config will be deleted
// "--global".
//
// It's the same as "git config [--global] --unset key"
func Unset(key string, global bool) error {
	g := ""
	if global {
		g = "--global"
	}

	cmd := exec.Command("git", "config", g, "--unset", key)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
