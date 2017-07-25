package main

import (
	"github.com/mkideal/cli"
	"os"
	"fmt"
	"errors"
	"io/ioutil"
	sCli "github.com/sinlov/golang_utils/cli"
	"github.com/sinlov/golang_utils/cfg"
	nig "github.com/ShubNig/AubNig/aubnig"
	"path/filepath"
)

const (
	VERSION_NAME string = "0.1.0"
)

func readFileAsString(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("read file error %s\n", err)
	}
	s := string(b)
	return s, nil
}

func isFileExist(filePath string) bool {
	f, err := os.Open(filePath)
	if err != nil || os.IsNotExist(err) {
		return false
	}
	defer f.Close()
	return true
}

type makerT struct {
	cli.Helper
	TempURL       string `cli:"tempUrl" usage:"Choose Temple git URL, this has default read by config.conf" dft:""`
	TempTAG       string `cli:"t,tempTag" usage:"Choose Temple git tag, this has default read by config.conf" dft:""`
	ProjectName   string `cli:"p,projectName" usage:"maker new project name" prompt:"Input want build project name"`
	Group         string `cli:"g,group" usage:"maker group" prompt:"Input group code(default: com.sinlov.android)"`
	ModuleName    string `cli:"m,moduleName" usage:"maker new out module name" prompt:"Input want build module name"`
	ArtifactId    string `cli:"i,artifactId" usage:"maker group" prompt:"Input artifact id (default: modulename)"`
	DeveloperName string `cli:"d,developerName" usage:"maker developer name" prompt:"Input developer name"`
	VersionName   string `cli:"n,versionName" usage:"maker version name" prompt:"Input version name(default: 0.0.1)"`
	VersionCode   int `cli:"c,versionCode" usage:"maker version code" prompt:"Input version code(default: 1)"`
}

var maker = &cli.Command{
	// child cli must has Name
	Name: nig.CLI_CHILD_MAKER_NAME,
	Desc: nig.CLI_CHILD_MAKER_DESC,
	Argv: func() interface{} {
		return new(makerT)
	},
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*makerT)
		tempUrl := argv.TempURL
		projectName := argv.ProjectName
		err := checkInputStringParams(projectName, "projectName")
		projectPath := filepath.Join(sCli.CurrentDirectory(), projectName)
		group := argv.Group
		if group == "" {
			group = nig.DEFAULT_GROUP

		}
		moduleName := argv.ModuleName
		err = checkInputStringParams(moduleName, "moduleName")
		developerName := argv.DeveloperName
		err = checkInputStringParams(developerName, "developerName")
		if err != nil {
			return err
		}
		artifactId := argv.ArtifactId
		if artifactId == "" {
			artifactId = moduleName
		}
		versionName := argv.VersionName
		if versionName == "" {
			versionName = nig.DEFAULT_VERSION_NAME
		}
		versionCode := argv.VersionCode
		if versionCode == 0 {
			versionCode = nig.DEFAULT_VERSION_CODE
		}
		if tempUrl == "" {
			gitTempUrl := cfgFile.Read(nig.KEY_NODE_GIT, nig.KEY_GIT_URL)
			if gitTempUrl == "" {
				// for test
				gitTempUrl = nig.DEFAULT_GIT_URL
			}
			tempUrl = gitTempUrl
		}

		ctx.String("\n=== Your setting start ===\n")
		ctx.String("temp Url: %v\n", tempUrl)
		ctx.String("group : %v\n", group)
		ctx.String("project name: %v\n", projectName)
		ctx.String("project Path: %v\n", projectPath)
		ctx.String("module name: %v\n", moduleName)
		ctx.String("artifact_id : %v\n", artifactId)
		ctx.String("developer Name : %v\n", developerName)
		ctx.String("version Name : %v\n", versionName)
		ctx.String("version Code : %v\n", versionCode)
		ctx.String("=== Your setting end ===\n")
		return nil
	},
}

func checkInputStringParams(stringParams string, showParams string) error {
	if stringParams == "" {
		return errors.New("\nYou are not setting [ " + showParams + " ] exit!")
	}
	return nil
}

type rootT struct {
	cli.Helper
	Version bool `cli:"!v" usage:"force flag, note the !"`
}

var root = &cli.Command{
	Desc: "This is ShubNig Utils command",
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} {
		return new(rootT)
	},
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		if argv.Version {
			ctx.String("Thanks to use ShubNig Utils \nNow version v%s", VERSION_NAME)
		}
		return nil
	},
}

var help = cli.HelpCommand("display help information")

var cfgFile = new(cfg.Cfg)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Warning you input is error please use -h or help to see help")
		os.Exit(1)
	}
	configFilePath := sCli.CommandPath() + "/config.conf"
	if isFileExist(configFilePath) {
		cfgFile.InitCfg(configFilePath)
	} else {
		fmt.Println("can not find config.conf file at root path: " + sCli.CommandPath())
		//os.Exit(1)
	}

	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(maker),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
