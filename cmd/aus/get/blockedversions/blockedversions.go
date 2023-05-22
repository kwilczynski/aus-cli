/*
Copyright (c) 2023 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package blockedversions

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
	"gitlab.cee.redhat.com/service/aus-cli/pkg/backend"
	"gitlab.cee.redhat.com/service/aus-cli/pkg/output"
)

var args struct {
	organizationId string
}

var Cmd = &cobra.Command{
	Use:   "version-blocks",
	Short: "Lists the blocked versions for an organization",
	Long:  "Lists the blocked versions for an organization",
	RunE:  run,
}

func init() {
	cmdFlags := Cmd.Flags()
	cmdFlags.StringVarP(
		&args.organizationId,
		"org-id",
		"o",
		"",
		"The ID of the OCM organization to inspect",
	)
}

func run(cmd *cobra.Command, argv []string) error {
	backendType, err := cmd.Flags().GetString("backend")
	if err != nil {
		return err
	}
	be, err := backend.NewPolicyBackend(backendType)
	if err != nil {
		return err
	}
	blockedVersions, err := be.ListBlockedVersionExpressions(args.organizationId)
	if err != nil {
		return err
	}
	body, _ := json.Marshal(blockedVersions)
	return output.PrettyList(os.Stdout, body)
}
