package aubnig

import (
	"path/filepath"
	"github.com/sinlov/golang_utils/files"
	sCli "github.com/sinlov/golang_utils/cli"
	"fmt"
	"os"
	"strings"
)

// if not find config Path just try to use GOPATH code github.com/ShubNig/AubNig/config.conf
// if code aubnig.conf and run root path not found, return ""
func Try2FindOutConfigPath() (string, string) {
	configFilePath := filepath.Join(sCli.CommandPath(), "aubnig.conf")
	projectPath := sCli.CurrentDirectory()
	if files.IsFileExist(configFilePath) {
		return configFilePath, projectPath
	}
	fmt.Printf("\nWarning!\ncan not find config.conf file at aubnig path: %s\n", sCli.CommandPath())
	go_path_env := os.Getenv("GOPATH")
	go_path_env_s := strings.Split(go_path_env, ":")
	is_find_dev_conf := false
	for _, path := range go_path_env_s {
		codePath := filepath.Join(path, "src", "github.com", "ShubNig", "AubNig")
		futurePath := filepath.Join(codePath, "aubnig.conf")
		projectPath = filepath.Join(codePath, "build")
		if files.IsFileExist(futurePath) {
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
	return configFilePath, projectPath
}
