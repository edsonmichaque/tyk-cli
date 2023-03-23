// Copyright 2023 Edson Michaque
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
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"runtime"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/edsonmichaque/template-cli/internal/build"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func cmdVersion(opts *Options) *Cmd {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Check version",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			tpl := heredoc.Doc(`
				Template CLI version:  %v
				Template API endpoint: %v
				Template API version:  %v
				OS/Arch:               %v/%v
			`)

			cmd.Printf(
				tpl,
				build.Version,
				prodBaseURL,
				"v1",
				runtime.GOOS,
				runtime.GOARCH,
			)

			return nil
		},
	}

	return newCmd(cmd, withOptions(opts))
}
