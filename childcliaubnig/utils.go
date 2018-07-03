package childcliaubnig

import (
	"fmt"
	"errors"
	"regexp"
	"strings"
	"github.com/sinlov/golang_utils/jstring"
	"os"
)

const REG_JAVA_PKG string = `^[a-z]+[a-z0-9_]+[\.][a-z]+[a-z0-9_]+[\.][a-z]+[a-z0-9_]*$`

// check cli input, if input empty will return err
// return nil is pass
func checkCliInputStringParams(stringParams string, showParams string) error {
	if stringParams == "" {
		return errors.New(fmt.Sprintf("\nYou are not setting [ %s ] exit!", showParams))
	}
	return nil
}

// input check string use reg check /^[a-z]+[a-z0-9_]+[\.][a-z]+[a-z0-9_]+[\.][a-z]+[a-z0-9_]*$/
// return nil is pass
func checkPackageNameAsJava(forCheckPKGName string) error {
	forCheckPKGName = strings.TrimSpace(forCheckPKGName)
	if forCheckPKGName == "" {
		return errors.New("package name is empty")
	} else {
		javaExp, err := regexp.Compile(REG_JAVA_PKG)
		if err != nil {
			return err
		}
		if javaExp.MatchString(forCheckPKGName) {
			return nil
		} else {
			return errors.New(fmt.Sprintf("package [ %s ] not java package name", forCheckPKGName))
		}
	}

}

// check module name not in can not setting name like gradle app test
// return nil is pass
func checkModuleNameAsGradle(forCheckModuleName string) error {
	forCheckModuleName = strings.TrimSpace(forCheckModuleName)
	if forCheckModuleName == "" {
		return errors.New("module name is empty")
	}
	if jstring.StringStartWith(forCheckModuleName, ".") {
		return errors.New("module name not start with '.'")
	}
	if forCheckModuleName == "build" ||
		forCheckModuleName == "gradle" ||
		forCheckModuleName == "test" ||
		forCheckModuleName == "app" ||
		forCheckModuleName == "keystore" ||
		forCheckModuleName == "scripts" ||
		forCheckModuleName == "node_module" {
		return errors.New("module name not name as [ build gradle test app keystore scripts node_module ]")
	}
	return nil
}

func isPathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if ! os.IsExist(err) {
			return false
		} else {
			return false
		}
	}
	return true
}
