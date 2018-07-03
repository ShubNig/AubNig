package aubnig

import (
	"path/filepath"
	"fmt"
	"os"
	"strings"
	sFiles "github.com/sinlov/golang_utils/files"
	sCli "github.com/sinlov/golang_utils/cli"
)

const (
	defaultConfPath = "aubnig.json"
	gitRepo         = "AubNig"
	gitUser         = "ShubNig"
	gitHost         = "github.com"
)

// if not find config Path just try to use GOPATH code github.com/ShubNig/AubNig/config.conf
// if code aubnig.conf and run root path not found, return ""
func Try2FindOutConfigPath() (string, string) {
	configFilePath := filepath.Join(sCli.CommandPath(), "aubnig.json")
	projectPath := sCli.CurrentDirectory()
	if sFiles.IsFileExist(configFilePath) {
		return configFilePath, projectPath
	}
	fmt.Printf("\nWarning!\nCan not find aubnig.json file at aubnig path: %s\n", sCli.CommandPath())
	goPathEnv := os.Getenv("GOPATH")
	goPathEnvS := strings.Split(goPathEnv, ":")
	isFindDevConf := false
	for _, path := range goPathEnvS {
		codePath := filepath.Join(path, "src", gitHost, gitUser, gitRepo)
		futurePath := filepath.Join(codePath, defaultConfPath)
		projectPath = filepath.Join(codePath, "build")
		if sFiles.IsFileExist(futurePath) {
			configFilePath = futurePath
			isFindDevConf = true
			break
		}
	}
	if isFindDevConf {
		fmt.Printf("just use dev config at path: %s\n", configFilePath)
	} else {
		fmt.Printf("can not load config at path: %s\nExit 1\n", configFilePath)
		configFilePath = ""
	}
	return configFilePath, projectPath
}
