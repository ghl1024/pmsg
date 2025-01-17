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
	"github.com/lenye/pmsg/pkg/weixin/token"
)

// weiXinAccessTokenCmd 获取微信接口调用凭证
var weiXinAccessTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "get weixin access token",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		arg := token.CmdTokenParams{
			UserAgent: userAgent,
			AppID:     appID,
			AppSecret: appSecret,
		}
		if err := token.CmdGetAccessToken(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
	Example: "pmsg weixin token -i app_id -s app_secret",
}

func init() {
	weiXinAccessTokenCmd.Flags().StringVarP(&appID, flags.AppID, "i", "", "weixin app id (required)")
	weiXinAccessTokenCmd.MarkFlagRequired(flags.AppID)

	weiXinAccessTokenCmd.Flags().StringVarP(&appSecret, flags.AppSecret, "s", "", "weixin app secret (required)")
	weiXinAccessTokenCmd.MarkFlagRequired(flags.AppSecret)
}
