package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/alfred-zhong/goproxy-get/git"
)

type goGetArgs struct {
	d        bool
	f        bool
	fix      bool
	insecure bool
	t        bool
	u        bool
	v        bool
}

func (args goGetArgs) parseArgs() []string {
	result := []string{}
	if args.d {
		result = append(result, "-d")
	}
	if args.f {
		result = append(result, "-f")
	}
	if args.fix {
		result = append(result, "-fix")
	}
	if args.insecure {
		result = append(result, "-insecure")
	}
	if args.t {
		result = append(result, "-t")
	}
	if args.u {
		result = append(result, "-u")
	}
	if args.v {
		result = append(result, "-v")
	}

	return result
}

func main() {
	ggArgs := new(goGetArgs)
	flag.BoolVar(&ggArgs.d, "d", false, "Stop after downloading the packages; that is,it instructs get not to install the packages.")
	flag.BoolVar(&ggArgs.f, "f", false, "Valid only when -u is set, forces get -u not to verify that each package has been checked out from the source control repository implied by its import path.")
	flag.BoolVar(&ggArgs.fix, "fix", false, "Run the fix tool on the downloaded packages before resolving dependencies or building the code.")
	flag.BoolVar(&ggArgs.insecure, "insecure", false, "Permits fetching from repositories and resolving custom domains using insecure schemes such as HTTP. Use with caution")
	flag.BoolVar(&ggArgs.t, "t", false, "Also download the packages required to build the tests for the specified packages.")
	flag.BoolVar(&ggArgs.u, "u", false, "Use the network to update the named packages and their dependencies. By default, get uses the network to check out missing packages but does not use it to look for updates to existing packages.")
	flag.BoolVar(&ggArgs.v, "v", false, "Enables verbose progress and debug output.")

	proxy := flag.String("p", "127.0.0.1:1087", "proxy address:port to use")
	flag.Parse()

	// export system env
	os.Setenv("http_proxy", fmt.Sprintf("http://%s", *proxy))
	os.Setenv("https_proxy", fmt.Sprintf("http://%s", *proxy))

	// git config --global
	git.Config("http.proxy", fmt.Sprintf("http://%s", *proxy), true)
	git.Config("https.proxy", fmt.Sprintf("https://%s", *proxy), true)

	defer func() {
		os.Unsetenv("http_proxy")
		os.Unsetenv("https_proxy")

		git.Unset("http.proxy", true)
		git.Unset("https.proxy", true)
	}()

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
}
