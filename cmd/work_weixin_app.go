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
	"github.com/lenye/pmsg/pkg/weixin/work/message"
)

// workWeiXinAppCmd 企业微信应用消息
var workWeiXinAppCmd = &cobra.Command{
	Use:   "app",
	Short: "publish work weixin app message",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkSendAppParams{
			UserAgent:              userAgent,
			AccessToken:            accessToken,
			CorpID:                 corpID,
			CorpSecret:             corpSecret,
			ToUser:                 toUser,
			ToParty:                toParty,
			ToTag:                  toTag,
			AgentID:                agentID,
			MsgType:                msgType,
			Safe:                   safe,
			EnableIDTrans:          enableIDTrans,
			EnableDuplicateCheck:   enableDuplicateCheck,
			DuplicateCheckInterval: duplicateCheckInterval,
			Data:                   args[0],
		}
		if err := message.CmdWorkSendApp(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
	Example: "pmsg workweixin app -i corp_id -s corp_secret -e agent_id -o '@all' -m text 'hello world'",
}

func init() {
	workWeiXinAppCmd.AddCommand(workWeiXinUndoAppCmd)

	workWeiXinSetAccessTokenFlags(workWeiXinAppCmd)

	workWeiXinAppCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "work weixin user id list")
	workWeiXinAppCmd.Flags().StringVarP(&toParty, flags.ToParty, "p", "", "work weixin party id list")
	workWeiXinAppCmd.Flags().StringVarP(&toTag, flags.ToTag, "g", "", "work weixin tag id list")

	workWeiXinAppCmd.Flags().Int64VarP(&agentID, flags.AgentID, "e", 0, "work weixin agent id (required)")
	workWeiXinAppCmd.MarkFlagRequired(flags.AgentID)

	workWeiXinAppCmd.Flags().StringVarP(&msgType, flags.MsgType, "m", "", "message type (required)")
	workWeiXinAppCmd.MarkFlagRequired(flags.MsgType)

	workWeiXinAppCmd.Flags().IntVar(&safe, flags.Safe, 0, "safe")
	workWeiXinAppCmd.Flags().IntVarP(&enableIDTrans, flags.EnableIDTrans, "r", 0, "enable id translated")
	workWeiXinAppCmd.Flags().IntVarP(&enableDuplicateCheck, flags.EnableDuplicateCheck, "c", 0, "enable duplicate check")
	workWeiXinAppCmd.Flags().IntVarP(&duplicateCheckInterval, flags.DuplicateCheckInterval, "d", 1800, "duplicate check interval")
}
