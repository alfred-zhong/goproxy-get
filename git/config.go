package git

import (
	"strings"
	"io"
	"bufio"
	"bytes"
	"fmt"
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

// RemoveSection remove the section in git config.
//
// Same as "git config [--global] --remove-section section.key"
func RemoveSection(sectionKey string, global bool) error {
	g := ""
	if global {
		g = "--global"
	}

	cmd := exec.Command("git", "config", g, "--remove-section", sectionKey)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ListNameOnly list all the names exist in the git config file.
//
// Same as "git config [--global] --list --name-only"
func ListNameOnly(global bool) ([]string, error) {
	g := ""
	if global {
		g = "--global"
	}

	cmd := exec.Command("git", "config", g, "--list", "--name-only")
	b, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git config list name only fail: %v", err)
	}
	
	reader := bufio.NewReader(bytes.NewReader(b))
	names := []string{}
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return names, nil
			}
			return names, err
		}

		if len(line) > 0 {
			names = append(names, string(line))
		}
	}
}

// RemoveSectionIfEmpty removes section if the section is empty.
func RemoveSectionIfEmpty(sectionKey string) error {
	names, err := ListNameOnly(true)
	if err != nil {
		return err
	}

	exist := false
	for _, name := range names {
		if len(name) > len(sectionKey) && strings.HasPrefix(name, sectionKey) && name[len(sectionKey)] == '.' {
			exist = true
			break
		}
	}

	if !exist {
		if err := RemoveSection(sectionKey, true); err != nil {
			return err
		}
	}
	return nil
}