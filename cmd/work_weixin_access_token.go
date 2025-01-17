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
	"github.com/lenye/pmsg/pkg/weixin/work/token"
)

// workWeiXinAccessTokenCmd 获取企业微信接口调用凭证
var workWeiXinAccessTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "get work weixin access token",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		arg := token.CmdWorkTokenParams{
			UserAgent:  userAgent,
			CorpID:     corpID,
			CorpSecret: corpSecret,
		}
		if err := token.CmdWorkGetAccessToken(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
	Example: "pmsg workweixin token -i corp_id -s corp_secret",
}

func init() {
	workWeiXinAccessTokenCmd.Flags().StringVarP(&corpID, flags.CorpID, "i", "", "work weixin corp id (required)")
	workWeiXinAccessTokenCmd.MarkFlagRequired(flags.CorpID)

	workWeiXinAccessTokenCmd.Flags().StringVarP(&corpSecret, flags.CorpSecret, "s", "", "work weixin corp secret (required)")
	workWeiXinAccessTokenCmd.MarkFlagRequired(flags.CorpSecret)
}
