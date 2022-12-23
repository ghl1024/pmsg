// Copyright 2022 The pmsg Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
)

// slackCmd slack
var slackCmd = &cobra.Command{
	Use:     "slack",
	Aliases: []string{"sk"},
	Short:   "slack",
}

func init() {
	slackCmd.PersistentFlags().StringVarP(&userAgent, flags.UserAgent, "a", "", "http user agent")

	slackCmd.AddCommand(slackBotCmd)
}
