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
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/weixin/work/bot"
)

// weiXinWorkBotUploadCmd 企业微信群机器人上传文件
var weiXinWorkBotUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "work weixin group bot file upload",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := bot.CmdUploadParams{
			UserAgent: userAgent,
			Key:       secret,
			File:      args[0],
		}
		if err := bot.CmdUpload(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
	Example: "pmsg workweixin bot upload -k key /img/app.png",
}

func init() {
	weiXinWorkSetKeyFlags(weiXinWorkBotUploadCmd)
}
