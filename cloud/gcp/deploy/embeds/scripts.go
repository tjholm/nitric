package embeds

import (
	_ "embed"
)

func populateTemplate[T any](template string, args T) string {

}

//go:embed create-site-version.sh
var createSiteVersionScript string

func CreateSiteVersionScript() string {

}

//go:embed populate-site-files.sh
var populateSiteFilesScript string

func PopulateSiteFilesScript() string {

}

//go:embed upload-file.sh
var uploadFileScript string

func UploadFileScript() string {

}
