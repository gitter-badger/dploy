package main

import (
	"flag"
	"fmt"
	dploy "github.com/mhausenblas/dploy/lib"
	"os"
	"strings"
)

const (
	BANNER = `    .___        .__                  
  __| _/______  |  |    ____  ___.__.
 / __ | \____ \ |  |   /  _ \<   |  |
/ /_/ | |  |_> >|  |__(  <_> )\___  |
\____ | |   __/ |____/ \____/ / ____|
     \/ |__|                  \/     
`
	VERSION = "0.6.2"
)

var (
	cmd       string
	workspace string
	all       bool
)

func about() {
	fmt.Fprint(os.Stderr, BANNER)
	fmt.Fprint(os.Stderr, fmt.Sprintf("This is dploy version %s, using workspace [%s]\n", VERSION, workspace))
	fmt.Fprint(os.Stderr, fmt.Sprintf("Please visit http://dploy.sh to learn more about me,\n"))
	fmt.Fprint(os.Stderr, fmt.Sprintf("report issues and also how to contribute to this project.\n"))
	fmt.Fprint(os.Stderr, strings.Repeat("=", 57), "\n")
}

func init() {
	cwd, _ := os.Getwd()
	flag.StringVar(&workspace, "workspace", cwd, "directory in which to operate")
	flag.BoolVar(&all, "all", false, "output all available data, semantics are sub-command dependent")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage: dploy [args] <command>\n")
		fmt.Fprint(os.Stderr, "\nThe following commands are available:\n")
		fmt.Fprint(os.Stderr, "\tinit\t... creates a new app descriptor and inits `specs/`\n")
		fmt.Fprint(os.Stderr, "\tdryrun\t... validates app deployment using Marathon API\n")
		fmt.Fprint(os.Stderr, "\trun\t... launches the app using `dploy.app` and the content of `specs/`\n")
		fmt.Fprint(os.Stderr, "\tdestroy\t... tears down the app\n")
		fmt.Fprint(os.Stderr, "\tls\t... lists the app's resources\n")
		fmt.Fprint(os.Stderr, "\tps\t... lists runtime properties of the app\n")
		fmt.Fprint(os.Stderr, "\nValid (optional) arguments are:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() < 1 {
		about()
		flag.Usage()
		os.Exit(1)
	} else {
		cmd = flag.Args()[0]
	}
}

func main() {
	about()
	switch cmd {
	case "init":
		dploy.Init(workspace, all)
	case "dryrun":
		dploy.DryRun(workspace, all)
	case "run":
		dploy.Run(workspace, all)
	case "destroy":
		dploy.Destroy(workspace, all)
	case "ls":
		dploy.ListResources(workspace, all)
	case "ps":
		dploy.ListRuntimeProperties(workspace, all)
	default:
		fmt.Fprint(os.Stderr, flag.Args()[0], " is not a valid command\n")
		flag.Usage()
		os.Exit(2)
	}
}
