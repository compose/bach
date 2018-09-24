// Copyright © 2017 Compose, an IBM Company
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

package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the backups subcommand
var logfilesCmd = &cobra.Command{
	Use:   "logfiles",
	Short: "Commands for Logfiles",
	Long:  `A selection of subcommands are available for listing and retrieving Logfiles `,
}

func init() {
	RootCmd.AddCommand(logfilesCmd)
}
