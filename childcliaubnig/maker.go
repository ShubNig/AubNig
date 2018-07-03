package childcliaubnig

import (
	"github.com/mkideal/cli"
	"github.com/ShubNig/AubNig/aubnig"
	"strings"
)

type makerT struct {
	cli.Helper
	TempURL       string `cli:"tempUrl" usage:"Choose Temple git URL, this has default read by config.conf" dft:""`
	TempTAG       string `cli:"t,tempTag" usage:"Choose Temple git tag, this has default read by config.conf" dft:""`
	ProjectName   string `cli:"p,projectName" usage:"maker new project name" prompt:"Input want build project name"`
	Group         string `cli:"g,group" usage:"maker group" prompt:"Input group code(default aubnig.json: config.temp.group)"`
	ModuleName    string `cli:"m,moduleName" usage:"maker new out module name" prompt:"Input want build module name"`
	ArtifactId    string `cli:"i,artifactId" usage:"maker group" prompt:"Input artifact id (default aubnig.json: moduleName)"`
	DeveloperName string `cli:"d,developerName" usage:"maker developer name" prompt:"Input developer name (default aubnig.json: config.temp.developer_name)"`
	VersionName   string `cli:"n,versionName" usage:"maker version name" prompt:"Input version name(default aubnig.json: config.temp.version_name)"`
	VersionCode   int    `cli:"c,versionCode" usage:"maker version code" prompt:"Input version code(default aubnig.json: config.temp.version_code)"`
}

type Maker struct {
	Config      aubnig.ConfAubNig
	RunMode     string
	CodePath    string
	ProjectPath string
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
			argv := ctx.Argv().(*makerT)
			tempUrl := argv.TempURL
			projectName := argv.ProjectName
			err := checkCliInputStringParams(projectName, "projectName")
			if err != nil {
				return err
			}
			group := argv.Group
			if group == "" {
				group = m.Config.Temp.Group
				if group == "" {
					group = aubnig.DEFAULT_GROUP
				}
			} else {
				err = checkPackageNameAsJava(group)
				if err != nil {
					return err
				}
			}
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
			cfgTemp := m.Config.Temp
			developerName := argv.DeveloperName
			if developerName == "" {
				developerName = cfgTemp.DeveloperName
			}
			err = checkCliInputStringParams(developerName, "developerName")
			if err != nil {
				return err
			}
			artifactId := argv.ArtifactId
			if artifactId == "" {
				artifactId = moduleName
			}
			versionName := argv.VersionName
			if versionName == "" {
				versionName = cfgTemp.VersionName
				if versionName == "" {
					versionName = aubnig.DEFAULT_VERSION_NAME
				}
			}
			versionCode := argv.VersionCode
			if versionCode == 0 {
				versionCode = cfgTemp.VersionCode
				if versionCode == 0 {
					versionCode = aubnig.DEFAULT_VERSION_CODE
				}
			}
			if tempUrl == "" {
				tempUrl = cfgTemp.Git.GitURL
			}

			// for dev mode URL
			if m.RunMode == "dev" {
				tempUrl = aubnig.DEFAULT_GIT_URL
			}

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

			return nil
		},
	}
	return maker
}
