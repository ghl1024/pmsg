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
	"github.com/lenye/pmsg/pkg/weixin/offiaccount/message"
)

// weiXinOfficialAccountTplSubCmd 微信公众号一次性订阅消息
var weiXinOfficialAccountTplSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "publish weixin official account template subscribe message (onetime)",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdMpSendTemplateSubscribeParams{
			UserAgent:   userAgent,
			AccessToken: accessToken,
			AppID:       appID,
			AppSecret:   appSecret,
			ToUser:      toUser,
			TemplateID:  templateID,
			Scene:       scene,
			Title:       title,
			Url:         url,
			Mini:        mini,
			Data:        args[0],
		}
		if err := message.CmdMpSendTemplateSubscribe(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
	Example: "pmsg weixin offiaccount template subscribe -i app_id -s app_secret --scene scene --title title -p template_id -o open_id '{\"first\":{\"value\":\"test\"}}'",
}

func init() {
	weiXinSetAccessTokenFlags(weiXinOfficialAccountTplSubCmd)

	weiXinOfficialAccountTplSubCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "weixin user open id (required)")
	weiXinOfficialAccountTplSubCmd.MarkFlagRequired(flags.ToUser)

	weiXinOfficialAccountTplSubCmd.Flags().StringVarP(&templateID, flags.TemplateID, "p", "", "weixin template id (required)")
	weiXinOfficialAccountTplSubCmd.MarkFlagRequired(flags.TemplateID)

	weiXinOfficialAccountTplSubCmd.Flags().StringVar(&scene, flags.Scene, "", "weixin subscribe scene (required)")
	weiXinOfficialAccountTplSubCmd.MarkFlagRequired(flags.Scene)

	weiXinOfficialAccountTplSubCmd.Flags().StringVar(&title, flags.Title, "", "message title (required)")
	weiXinOfficialAccountTplSubCmd.MarkFlagRequired(flags.Title)

	weiXinOfficialAccountTplSubCmd.Flags().StringVar(&url, flags.Url, "", "url")
	weiXinOfficialAccountTplSubCmd.Flags().StringToStringVar(&mini, flags.Mini, nil, "weixin mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")
}
