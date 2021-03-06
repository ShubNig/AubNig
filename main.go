package main

import (
	"github.com/mkideal/cli"
	"os"
	"fmt"
	"github.com/ShubNig/AubNig/aubnig"
	"github.com/ShubNig/AubNig/childcliaubnig"
)

const (
	VERSION_NAME string = "0.1.0"
)

var runMode string
var projectPath string
var codeCatchPath string

var help = cli.HelpCommand("display help information")
var jsonCfg = new(aubnig.CfgAub)

type rootT struct {
	cli.Helper
	Version bool `cli:"!v" usage:"force flag, note the !"`
}

var root = &cli.Command{
	Desc: "This is AubNig Utils command",
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} {
		return new(rootT)
	},
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		if argv.Version {
			ctx.String("Thanks to use AubNig Utils \nNow version v%s", VERSION_NAME)
		}
		return nil
	},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Warning you input is error please use -h or help to see help")
		os.Exit(1)
	}
	configFilePath, findProjectPath := aubnig.Try2FindOutConfigPath()
	if configFilePath == "" {
		os.Exit(1)
	} else {
		err := jsonCfg.InitJsonCfg(configFilePath)
		if err != nil {
			fmt.Printf("init JsonConfig err %s\n", err.Error())
			os.Exit(1)
		}
	}
	if jsonCfg.IsDebug {
		fmt.Printf("===> now in %v mode all setting will be default <===\n", jsonCfg.ConfAubNig.RunMode)
		projectPath = findProjectPath
	}
	catchPath, err := aubnig.InitCatchPath()
	if err != nil {
		fmt.Printf("init catch err %s\n", err.Error())
		os.Exit(1)
	}
	codeCatchPath = catchPath
	makerDef := childcliaubnig.Maker{
		RunMode:     runMode,
		CodePath:    codeCatchPath,
		ProjectPath: projectPath,
		Config:      jsonCfg.ReadConfig(),
	}
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(makerDef.MakeCliDef()),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
