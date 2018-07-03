package childcliaubnig

import (
	"github.com/mkideal/cli"
	"github.com/ShubNig/AubNig/aubnig"
	"strings"
	"github.com/kataras/go-errors"
	sSli "github.com/sinlov/golang_utils/cli"
	"fmt"
	"os"
	"path/filepath"
)

type makerT struct {
	cli.Helper
	TempURL       string `cli:"tempUrl" usage:"Choose Temple git URL, this has default read by aubnig.json" dft:""`
	TempLocal     string `cli:"tempLocal" usage:"Choose Temple git local, this has default read by aubnig.json" dft:""`
	TempTAG       string `cli:"t,tempTag" usage:"Choose Temple git tag, this has default read by aubnig.json" dft:""`
	ProjectName   string `cli:"p,projectName" usage:"maker new project name" prompt:"Input want build project name"`
	Group         string `cli:"g,group" usage:"maker group" prompt:"Input group code(default aubnig.json: config.temp.group)"`
	ModuleName    string `cli:"m,moduleName" usage:"maker new out module name" prompt:"Input want build module name"`
	ModulePackage string `cli:"modulePackage" usage:"maker new out module package name" prompt:"Input want build module package name"`
	ArtifactId    string `cli:"i,artifactId" usage:"maker group" prompt:"Input artifact id (default aubnig.json: moduleName)"`
	DeveloperName string `cli:"d,developerName" usage:"maker developer name" prompt:"Input developer name (default aubnig.json: config.temp.developer_name)"`
	VersionName   string `cli:"n,versionName" usage:"maker version name" prompt:"Input version name(default aubnig.json: config.temp.version_name)"`
	VersionCode   int    `cli:"c,versionCode" usage:"maker version code" prompt:"Input version code(default aubnig.json: config.temp.version_code)"`
}

type Maker struct {
	Config        aubnig.ConfAubNig
	RunMode       string
	CodePath      string
	ProjectPath   string
	ModuleName    string
	ModulePackage string
	ArtifactId    string
	VersionName   string
	VersionCode   int
}

// runMode dev or prod
// codeCatchPath codeCatchPath
// projectPath project path
func (m *Maker) MakeCliDef() *cli.Command {
	var maker = &cli.Command{
		// child cli must has Name
		Name: aubnig.CLI_CHILD_MAKER_NAME,
		Desc: aubnig.CLI_CHILD_MAKER_DESC,
		Argv: func() interface{} {
			return new(makerT)
		},
		Fn: func(ctx *cli.Context) error {
			cfgTemp := m.Config.Temp
			argv := ctx.Argv().(*makerT)

			// projectName
			projectName := argv.ProjectName
			err := checkCliInputStringParams(projectName, "projectName")
			if err != nil {
				return err
			} else {
				// TODO sinlov 2018/7/3 ?do project name
			}

			// group
			group := argv.Group
			if group == "" {
				group = cfgTemp.Group
				if group == "" {
					group = aubnig.DEFAULT_GROUP
				}
			} else {
				err = checkPackageNameAsJava(group)
				if err != nil {
					return err
				}
			}
			cfgTemp.Group = group

			// moduleName
			moduleName := argv.ModuleName
			err = checkCliInputStringParams(moduleName, "moduleName")
			if err != nil {
				return err
			} else {
				err = checkModuleNameAsGradle(moduleName)
				if err != nil {
					return err
				}
				moduleName = strings.ToLower(moduleName)
			}
			m.ModuleName = moduleName

			// modulePackageName
			modulePackageName := argv.ModulePackage
			err = checkCliInputStringParams(modulePackageName, "modulePackageName")
			if err != nil {
				return err
			} else {
				err = checkPackageNameAsJava(modulePackageName)
				if err != nil {
					return err
				}
			}
			m.ModulePackage = modulePackageName

			// developerName
			developerName := argv.DeveloperName
			if developerName == "" {
				developerName = cfgTemp.DeveloperName
			}
			err = checkCliInputStringParams(developerName, "developerName")
			if err != nil {
				return err
			}
			cfgTemp.DeveloperName = developerName

			// artifactId
			artifactId := argv.ArtifactId
			if artifactId == "" {
				artifactId = moduleName
			}
			m.ArtifactId = artifactId

			// versionName
			versionName := argv.VersionName
			if versionName == "" {
				versionName = cfgTemp.VersionName
				if versionName == "" {
					versionName = aubnig.DEFAULT_VERSION_NAME
				}
			}
			m.VersionName = versionName

			versionCode := argv.VersionCode
			if versionCode == 0 {
				versionCode = cfgTemp.VersionCode
				if versionCode == 0 {
					versionCode = aubnig.DEFAULT_VERSION_CODE
				}
			}
			m.VersionCode = versionCode

			// tempLocal
			tempLocal := argv.TempLocal
			if tempLocal == "" {
				tempLocal = cfgTemp.Git.GitLocal
			}
			cfgTemp.Git.GitLocal = tempLocal

			// tempUrl
			tempUrl := argv.TempURL
			if tempUrl == "" {
				tempUrl = cfgTemp.Git.GitURL
			}
			// for dev mode URL
			if m.RunMode == "dev" {
				tempUrl = aubnig.DEFAULT_GIT_URL
			}
			cfgTemp.Git.GitURL = tempUrl

			// refresh Temp
			m.Config.Temp = cfgTemp

			ctx.String("\n=== Your setting start ===\n")
			ctx.String("temp codeCatchPath: %v\n", m.CodePath)
			ctx.String("temp Url: %v\n", tempUrl)
			ctx.String("group : %v\n", group)
			ctx.String("project name: %v\n", projectName)
			ctx.String("project Path: %v\n", m.ProjectPath)
			ctx.String("module name: %v\n", moduleName)
			ctx.String("artifact_id : %v\n", artifactId)
			ctx.String("developer Name : %v\n", developerName)
			ctx.String("version Name : %v\n", versionName)
			ctx.String("version Code : %v\n", versionCode)
			ctx.String("=== Your setting end ===\n")
			err = gitDownloadLastCommitByURL(ctx, m)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return maker
}

func gitDownloadLastCommitByURL(ctx *cli.Context, m *Maker) error {
	if m.Config.Temp.GitURL == "" {
		return errors.New("download git url is empty")
	}
	if m.ProjectPath == "" {
		return errors.New("project path is empty")
	}

	if ! isPathExist(m.ProjectPath) {
		os.MkdirAll(m.ProjectPath, os.ModePerm)
	}

	codeCatchAbsPath := filepath.Join(m.CodePath, m.Config.Git.GitLocal)
	if isPathExist(codeCatchAbsPath) {
		ctx.String("just has download code at %v\n", codeCatchAbsPath)
		updateCmd := fmt.Sprintf("cd %v && git pull && cd %v\n", codeCatchAbsPath, m.ProjectPath)
		ctx.String("updateCmd -> %v\n", updateCmd)
		b, code, strOut, strInfo := sSli.CmdExec("", updateCmd)
		if ! b {
			return errors.New(fmt.Sprintf("code %v, run error %v\n", code, strOut))
		} else {
			ctx.String("cmd out %v\n", strInfo)
		}
	} else {
		ctx.String("just start download by git as %v\n", m.Config.Temp.GitURL)
		cloneCmd := fmt.Sprintf("git clone %v %v", m.Config.Temp.GitURL, codeCatchAbsPath)
		ctx.String("clone -> %v\n", cloneCmd)
		b, code, strOut, strInfo := sSli.CmdExec("", cloneCmd)
		if ! b {
			return errors.New(fmt.Sprintf("code %v, run error %v\n", code, strOut))
		}
		ctx.String("cmd out %v\n", strInfo)
		ctx.String("download by git as %v at %v\n", m.Config.Temp.GitURL, codeCatchAbsPath)
	}

	return nil
}
