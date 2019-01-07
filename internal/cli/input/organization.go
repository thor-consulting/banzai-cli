// Copyright © 2018 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package input

import (
	"context"
	"log"

	"github.com/banzaicloud/banzai-cli/internal/cli"
	"gopkg.in/AlecAivazis/survey.v1"
)

// AskOrganization asks for an organization.
func AskOrganization(banzaiCli cli.Cli) int32 {
	orgs, _, err := banzaiCli.Client().OrganizationsApi.ListOrgs(context.Background())
	if err != nil {
		log.Fatalf("could not list organizations: %v", err)
	}

	orgSelection := make([]string, len(orgs))
	orgResultMap := make(map[string]int32, len(orgs))
	for i, org := range orgs {
		orgSelection[i] = org.Name
		orgResultMap[org.Name] = org.Id
	}

	var name string
	err = survey.AskOne(&survey.Select{Message: "Organization:", Options: orgSelection}, &name, survey.Required)
	if err != nil {
		log.Fatalf("could choose an organization organizations: %v", err)
	}

	return orgResultMap[name]
}

// GetOrganization returns the current organization.
func GetOrganization(banzaiCli cli.Cli) int32 {
	id := banzaiCli.Context().OrganizationID()

	if id == 0 {
		id = AskOrganization(banzaiCli)
	}

	return id
}