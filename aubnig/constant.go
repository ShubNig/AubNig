package aubnig

const (
	// mode has dev test prd
	MODE string = "dev"
	DEV_PROJ_PATH string = "/Users/sinlov/goPath/src/github.com/ShubNig/AubNig"

	DEFAULT_GIT_URL string = "https://github.com/MDL-Sinlov/MDL_Android-Temp.git"
	DEFAULT_GIT_TAG string = ""
	DEFAULT_VERSION_NAME string = "0.0.1"
	DEFAULT_VERSION_CODE int = 1
	DEFAULT_GROUP string = "com.sinlov.android"

	CLI_CHILD_MAKER_NAME = "maker"
	CLI_CHILD_MAKER_DESC string = "[ maker ] is AubNig make tools, " +
		"it can make android project for single project"

	KEY_NODE_GIT string = "GitSet"
	KEY_GIT_URL string = "GitTempURL"
	KEY_GIT_TAG string = "GitTag"

	KEY_NODE_SET string = "KeySet"
	KEY_SET_BASE_KEY = "base_key"
)
