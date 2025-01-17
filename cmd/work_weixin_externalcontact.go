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
	"strings"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/work/message"
)

// workWeiXinExternalContactCmd 企业微信家校消息
var workWeiXinExternalContactCmd = &cobra.Command{
	Use:     "externalcontact",
	Aliases: []string{"ec"},
	Short:   "publish work weixin external contact message",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		arg := message.CmdWorkSendExternalContactParams{
			UserAgent:              userAgent,
			AccessToken:            accessToken,
			CorpID:                 corpID,
			CorpSecret:             corpSecret,
			RecvScope:              recvScope,
			ToAll:                  toAll,
			MsgType:                msgType,
			AgentID:                agentID,
			EnableIDTrans:          enableIDTrans,
			EnableDuplicateCheck:   enableDuplicateCheck,
			DuplicateCheckInterval: duplicateCheckInterval,
			Data:                   args[0],
		}

		if toParentUserID != "" {
			arg.ToParentUserID = strings.Split(toParentUserID, "|")
		}
		if toStudentUserID != "" {
			arg.ToStudentUserID = strings.Split(toStudentUserID, "|")
		}
		if toParty != "" {
			arg.ToParty = strings.Split(toParty, "|")
		}

		if err := message.CmdWorkSendExternalContact(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
	Example: "pmsg workweixin externalcontact -i corp_id -s corp_secret -e agent_id -n 'parentuserid1|parentuserid2' -m text 'hello world'",
}

func init() {
	workWeiXinSetAccessTokenFlags(workWeiXinExternalContactCmd)

	workWeiXinExternalContactCmd.Flags().IntVarP(&recvScope, flags.RecvScope, "o", 0, "receive scope")

	workWeiXinExternalContactCmd.Flags().StringVarP(&toParentUserID, flags.ToParentUserID, "n", "", "work weixin parent user id list")
	workWeiXinExternalContactCmd.Flags().StringVarP(&toStudentUserID, flags.ToStudentUserID, "u", "", "work weixin student user id list")
	workWeiXinExternalContactCmd.Flags().StringVarP(&toParty, flags.ToParty, "p", "", "work weixin party id list")
	workWeiXinExternalContactCmd.Flags().IntVarP(&toAll, flags.ToAll, "l", 0, "send to all user")

	workWeiXinExternalContactCmd.Flags().Int64VarP(&agentID, flags.AgentID, "e", 0, "work weixin agent id (required)")
	workWeiXinExternalContactCmd.MarkFlagRequired(flags.AgentID)

	workWeiXinExternalContactCmd.Flags().StringVarP(&msgType, flags.MsgType, "m", "", "message type (required)")
	workWeiXinExternalContactCmd.MarkFlagRequired(flags.MsgType)

	workWeiXinExternalContactCmd.Flags().IntVarP(&enableIDTrans, flags.EnableIDTrans, "r", 0, "enable id translated")
	workWeiXinExternalContactCmd.Flags().IntVarP(&enableDuplicateCheck, flags.EnableDuplicateCheck, "c", 0, "enable duplicate check")
	workWeiXinExternalContactCmd.Flags().IntVarP(&duplicateCheckInterval, flags.DuplicateCheckInterval, "d", 1800, "duplicate check interval")

}
