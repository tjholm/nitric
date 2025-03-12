// Copyright Nitric Pty Ltd.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deploy

import (
	"encoding/json"
	"fmt"

	deploymentspb "github.com/nitrictech/nitric/core/pkg/proto/deployments/v1"
	"github.com/pulumi/pulumi-command/sdk/go/command/local"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/firebase"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func getAccessToken() {

}

func createHostingSiteVersion(ctx *pulumi.Context) error {
	// Use curl to create a new version of the firebase hosting website
	// TODO: This would be much better as a custom provider
	createVersionCmd, err := local.NewCommand(ctx, "createVersion", &local.CommandArgs{
		Dir:    pulumi.String("website"),
		Create: pulumi.String("curl -X POST -H \"Content-Type: application/json\" -d '{\"version\": {\"status\": \"CREATED\"}}' https://firebasehosting.googleapis.com/v1beta1/sites/$PROJECT_ID/versions"),
		Delete: pulumi.String("curl -X DELETE https://firebasehosting.googleapis.com/v1beta1/sites/$PROJECT_ID/versions/$VERSION_ID"),
	})
	if err != nil {
		return err
	}
	// Apply the output and convert to a pulumi string map
	outputMap := createVersionCmd.Stdout.ApplyT(func(stdout string) (map[string]interface{}, error) {
		output := map[string]interface{}{}
		err := json.Unmarshal([]byte(stdout), &output)
		if err != nil {
			return nil, err
		}

		return output, nil
	}).(pulumi.MapOutput)

	versionName, ok := outputMap.MapIndex(pulumi.String("name")).(pulumi.StringOutput)
	if !ok {
		return fmt.Errorf("failed to get version name from output")
	}

	// Use curl to upload the files to the firebase hosting website
	populateFilesCmd, err := local.NewCommand(ctx, "populateFiles", &local.CommandArgs{
		Dir:    pulumi.String("website"),
		Create: pulumi.String("curl"),
		// We may not need a delete command here
		// Delete: pulumi.String("curl -X DELETE https://firebasehosting.googleapis.com/v1beta1/sites/$PROJECT_ID/versions/$VERSION_ID/files"),
	})
	if err != nil {
		return err
	}

	// Walk the websites structure and create gzips of each file and a sha256 hash of each gzip
	// Then send the hashes as a map to the populateFiles endpoint for the created version

	finalizeVersionCmd, err := local.NewCommand(ctx, "finalizeVersion", &local.CommandArgs{
		Dir:    pulumi.String("website"),
		Create: pulumi.String("curl"),
	})

	return err
}

func (a *NitricGcpPulumiProvider) createFirebaseSite(ctx *pulumi.Context) error {
	webapp, err := firebase.NewWebApp(ctx, "webApp", &firebase.WebAppArgs{
		Project:        pulumi.String(a.GcpConfig.ProjectId),
		DisplayName:    pulumi.Sprintf("Web application hosting for Nitric stack: %s", a.StackName),
		DeletionPolicy: pulumi.String("delete"),
	})
	if err != nil {
		return err
	}

	// Create a new firebase hosting site
	site, err := firebase.NewHostingSite(ctx, "site", &firebase.HostingSiteArgs{
		Project: pulumi.String(a.GcpConfig.ProjectId),
		SiteId:  pulumi.String(a.StackName),
		AppId:   webapp.AppId,
	})

	// Create the hosting site version manually
	// we do this because we need to upload the files to the site
	// and existing provider implementations do not allow this
}

// Website - Implements the Website deployment method for the GCP provider
func (a *NitricGcpPulumiProvider) Website(ctx *pulumi.Context, parent pulumi.Resource, name string, config *deploymentspb.Website) error {
	// Create a new firebase hosting site

	//

	// Pull the web uri out of the firebase deploy output?
	// Can probably already get this from the firebase project itself
	// cmd.Stdout

	return fmt.Errorf("websites aren't yet supported by Nitric on Google Cloud, remove the websites from your project and try again or run `nitric down` to destroy this stack")
}
