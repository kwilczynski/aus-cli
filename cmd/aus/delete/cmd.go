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

package delete

import (
	"github.com/spf13/cobra"
	"gitlab.cee.redhat.com/service/aus-cli/cmd/aus/delete/policy"
)

var Cmd = &cobra.Command{
	Use:           "delete",
	Short:         "Delete AUS resources",
	Long:          "Delete AUS resources",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	// Register the subcommands:
	Cmd.AddCommand(policy.Cmd)
}
