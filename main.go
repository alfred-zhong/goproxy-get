package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/alfred-zhong/goproxy-get/git"
)

type goGetArgs []struct {
	value bool
	name  string
	desc  string
}

func (args goGetArgs) parseArgs() []string {
	result := []string{}
	for _, arg := range args {
		if arg.value {
			result = append(result, "-"+arg.name)
		}
	}

	return result
}

var ggArgs = goGetArgs{
	{false, "d", "Stop after downloading the packages; that is,it instructs get not to install the packages."},
	{false, "f", "Valid only when -u is set, forces get -u not to verify that each package has been checked out from the source control repository implied by its import path."},
	{false, "fix", "Run the fix tool on the downloaded packages before resolving dependencies or building the code."},
	{false, "insecure", "Permits fetching from repositories and resolving custom domains using insecure schemes such as HTTP. Use with caution"},
	{false, "t", "Also download the packages required to build the tests for the specified packages."},
	{false, "u", "Use the network to update the named packages and their dependencies. By default, get uses the network to check out missing packages but does not use it to look for updates to existing packages."},
	{false, "v", "Enables verbose progress and debug output."},
}

func main() {
	for i := range ggArgs {
		flag.BoolVar(&ggArgs[i].value, ggArgs[i].name, ggArgs[i].value, ggArgs[i].desc)
	}

	proxy := flag.String("p", "127.0.0.1:1087", "proxy address:port to use")
	flag.Parse()

	// export system env
	os.Setenv("http_proxy", fmt.Sprintf("http://%s", *proxy))
	os.Setenv("https_proxy", fmt.Sprintf("http://%s", *proxy))

	// git config --global
	git.Config("http.proxy", fmt.Sprintf("http://%s", *proxy), true)
	git.Config("https.proxy", fmt.Sprintf("https://%s", *proxy), true)

	// run "go get ..."
	arg := []string{"get"}
	arg = append(arg, ggArgs.parseArgs()...)
	arg = append(arg, flag.Args()...)
	cmd := exec.Command("go", arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "run go get fail: %v", err)
	}

	// unset the proxy env after "go get ..."
	os.Unsetenv("http_proxy")
	os.Unsetenv("https_proxy")
	git.Unset("http.proxy", true)
	git.Unset("https.proxy", true)
	git.RemoveSectionIfEmpty("http")
	git.RemoveSectionIfEmpty("https")
}
