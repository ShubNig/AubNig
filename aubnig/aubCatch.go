package aubnig

import (
	"path/filepath"
	"os"
	sCli "github.com/sinlov/golang_utils/cli"
	"github.com/sinlov/golang_utils/files"
)

func InitCatchPath() (string, error) {
	homePath, err := sCli.Home()
	if err != nil {
		return "", err
	}
	catchPath := filepath.Join(homePath, ".ShubNig", "AubNig")
	if !files.IsFileExist(catchPath) {
		err := os.MkdirAll(catchPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	return catchPath, nil
}
