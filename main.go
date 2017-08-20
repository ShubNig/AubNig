package main

import (
	"github.com/mkideal/cli"
	"os"
	"fmt"
	"errors"
	"io/ioutil"
	sCli "github.com/sinlov/golang_utils/cli"
	"github.com/sinlov/golang_utils/cfg"
	"github.com/ShubNig/AubNig/aubnig"
	"path/filepath"
	"strings"
)

const (
	VERSION_NAME string = "0.1.0"
)

var runMode string

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

// if not find config Path just try to use GOPATH code github.com/ShubNig/AubNig/config.conf
// if code aubnig.conf and run root path not found, return ""
func try2FindOutConfigPath() string {
	configFilePath := filepath.Join(sCli.CommandPath(), "aubnig.conf")
	if isFileExist(configFilePath) {
		cfgFile.InitCfg(configFilePath)
		return configFilePath
	}
	fmt.Printf("\nWarning!\ncan not find config.conf file at aubnig path: %s\n", sCli.CommandPath())
	go_path_env := os.Getenv("GOPATH")
	go_path_env_s := strings.Split(go_path_env, ":")
	is_find_dev_conf := false
	for _, path := range go_path_env_s {
		futurePath := filepath.Join(path, "src", "github.com", "ShubNig", "AubNig", "aubnig.conf")
		if isFileExist(futurePath) {
			configFilePath = futurePath
			is_find_dev_conf = true
			break
		}
	}
	if is_find_dev_conf {
		fmt.Printf("just use dev config at path: %s\n", configFilePath)
	} else {
		fmt.Printf("can not load config at path: %s\nExit 1\n", configFilePath)
		configFilePath = ""
	}
	return configFilePath
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
	Name: aubnig.CLI_CHILD_MAKER_NAME,
	Desc: aubnig.CLI_CHILD_MAKER_DESC,
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
			group = aubnig.DEFAULT_GROUP

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
			versionName = aubnig.DEFAULT_VERSION_NAME
		}
		versionCode := argv.VersionCode
		if versionCode == 0 {
			versionCode = aubnig.DEFAULT_VERSION_CODE
		}
		if tempUrl == "" {
			gitTempUrl := cfgFile.Read(aubnig.KEY_NODE_GIT, aubnig.KEY_GIT_URL)
			tempUrl = gitTempUrl
		}

		if runMode == "dev" {
			tempUrl = aubnig.DEFAULT_GIT_URL
			projectPath = aubnig.DEV_PROJECT_PATH
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
		return errors.New(fmt.Sprintf("\nYou are not setting [ %s ] exit!", showParams))
	}
	return nil
}

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

var help = cli.HelpCommand("display help information")

var cfgFile = new(cfg.Cfg)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Warning you input is error please use -h or help to see help")
		os.Exit(1)
	}
	configFilePath := try2FindOutConfigPath()
	if configFilePath == "" {
		os.Exit(1)
	} else {
		cfgFile.InitCfg(configFilePath)
	}
	runMode = cfgFile.Read(aubnig.KEY_NODE_AUBNIG, aubnig.KEY_RUN_MODE)
	if runMode == "dev" {
		fmt.Printf("===> now in %s mode all setting will be default <===\n", runMode)
	}
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(maker),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
