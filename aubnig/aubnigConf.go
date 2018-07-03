package aubnig

type ConfAubNig struct {
	RunMode string `json:"runMode"`
	Temp           `json:"temp"`
}

type Temp struct {
	DeveloperName string `json:"developer_name"`
	Git                  `json:"git"`
	GradleVersion string `json:"gradle_version"`
	Group         string `json:"group"`
	VersionCode   int    `json:"version_code"`
	VersionName   string `json:"version_name"`
}

type Git struct {
	GitBranch string `json:"git_branch"`
	GitTag    string `json:"git_tag"`
	GitURL    string `json:"git_url"`
	GitLocal  string `json:"git_local"`
}
