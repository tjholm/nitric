package embeds

import (
	_ "embed"
	"html/template"
	"strings"
)

func templateFactory[T any](name string, tmpl string) func(T) (string, error) {
	temp, parseErr := template.New(name).Parse(tmpl)
	return func(args T) (string, error) {
		if parseErr != nil {
			return "", parseErr
		}

		var output strings.Builder

		err := temp.Execute(&output, args)
		if err != nil {
			return "", err
		}

		return output.String(), nil
	}
}

//go:embed create-site-version.sh
var createSiteVersionScript string

type CreateSiteVersionScriptArgs struct {
	Config any
	SiteId string
}

var CreateSiteVersion = templateFactory[CreateSiteVersionScriptArgs]("createSiteVersionScript", createSiteVersionScript)

//go:embed delete-site-version.sh
var deleteSiteVersionScript string

type DeleteSiteVersionScriptArgs struct {
	// We don't need args here we can pull the values from the last known stdout
	// of the create/update commands run by pulumi
}

var DeleteSiteVersion = templateFactory[DeleteSiteVersionScriptArgs]("deleteSiteVersionScript", deleteSiteVersionScript)

//go:embed populate-site-files.sh
var populateSiteFilesScript string

type PopulateSiteFilesScriptArgs struct {
	SiteId    string
	VersionId string
	Files     map[string]string
}

var PopulateSiteFilesScript = templateFactory[PopulateSiteFilesScriptArgs]("populateSiteFilesScript", populateSiteFilesScript)

//go:embed upload-file.sh
var uploadFileScript string

type UploadFileScriptArgs struct {
	FilePath  string
	SiteId    string
	VersionId string
	FileHash  string
}

var UploadFilesScript = templateFactory[UploadFileScriptArgs]("uploadFileScript", uploadFileScript)

//go:embed finalize-site-version.sh
var finalizeSiteVersionScript string

type FinalizeSiteVersionScriptArgs struct {
	SiteId    string
	VersionId string
}

var FinalizeSiteVersion = templateFactory[FinalizeSiteVersionScriptArgs]("finalizeSiteVersionScript", finalizeSiteVersionScript)

//go:embed release-site-version.sh
var releaseSiteVersionScript string

type ReleaseSiteVersionScriptArgs struct {
	SiteId    string
	VersionId string
}

var ReleaseSiteVersion = templateFactory[ReleaseSiteVersionScriptArgs]("releaseSiteVersionScript", releaseSiteVersionScript)
