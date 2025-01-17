// Copyright 2022-2023 The pmsg Authors. All rights reserved.
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
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/slack/bot"
)

// slackBotCmd slack bot
var slackBotCmd = &cobra.Command{
	Use:   "bot",
	Short: "publish slack bot message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := bot.CmdSendParams{
			UserAgent: userAgent,
			URL:       url,
			Data:      args[0],
		}
		if err := bot.CmdSend(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
	Example: "pmsg slack bot --url webhook_url '{\"text\": \"Hello, World!\"}'",
}

func init() {
	slackBotCmd.Flags().StringVar(&url, flags.Url, "", "slack webhook url")
	slackBotCmd.MarkFlagRequired(flags.Url)
}
