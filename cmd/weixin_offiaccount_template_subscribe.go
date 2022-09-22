package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
	"github.com/lenye/pmsg/pkg/weixin/offiaccount/message"
	"github.com/lenye/pmsg/pkg/weixin/token"
)

const (
	nameScene = "scene"
	nameTitle = "title"
)

var (
	scene string
	title string
)

// weiXinOfficialAccountTplSubCmd 微信公众号一次性订阅消息
var weiXinOfficialAccountTplSubCmd = &cobra.Command{
	Use:     "subscribe",
	Aliases: []string{"sub"},
	Short:   "publish weixin offiaccount template subscribe message (onetime)",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := WeiXinOfficialAccountSendTemplateSubscribe(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinOfficialAccountTplCmd.AddCommand(weiXinOfficialAccountTplSubCmd)

	weiXinSetAccessTokenFlags(weiXinOfficialAccountTplSubCmd)

	weiXinOfficialAccountTplSubCmd.Flags().StringVarP(&openID, nameOpenID, "o", "", "weixin user open id (required)")
	weiXinOfficialAccountTplSubCmd.MarkFlagRequired(nameOpenID)

	weiXinOfficialAccountTplSubCmd.Flags().StringVarP(&templateID, nameTemplateID, "p", "", "weixin template id (required)")
	weiXinOfficialAccountTplSubCmd.MarkFlagRequired(nameTemplateID)

	weiXinOfficialAccountTplSubCmd.Flags().StringVar(&scene, nameScene, "", "weixin subscribe scene (required)")
	weiXinOfficialAccountTplSubCmd.MarkFlagRequired(nameScene)

	weiXinOfficialAccountTplSubCmd.Flags().StringVar(&title, nameTitle, "", "weixin message title (required)")
	weiXinOfficialAccountTplSubCmd.MarkFlagRequired(nameTitle)

	weiXinOfficialAccountTplSubCmd.Flags().StringVarP(&url, nameUrl, "u", "", "url")
	weiXinOfficialAccountTplSubCmd.Flags().StringToStringVarP(&mini, nameMini, "m", nil, "weixin template mini program, example: app_id=XiaoChengXuAppId,page_path=index?foo=bar")
}

// WeiXinOfficialAccountSendTemplateSubscribe 发送微信公众号一次性订阅消息
func WeiXinOfficialAccountSendTemplateSubscribe(args []string) error {

	if accessToken == "" {
		if appID == "" {
			return ErrMultiRequiredOne
		}
	}

	var dataItem map[string]message.TemplateDataItem
	buf := bytes.NewBufferString("")
	buf.WriteString(args[0])
	if buf.String() != "" {
		if err := json.Unmarshal(buf.Bytes(), &dataItem); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		for k, v := range dataItem {
			if v.Value == "" {
				return fmt.Errorf("data %v.value not set", k)
			}
			if len(v.Value) > 200 {
				return fmt.Errorf("data %v.value maximum length is within 200", k)
			}
		}
	}

	if len(title) > 15 {
		return fmt.Errorf("flag %q maximum length is within 15", nameTitle)
	}

	msg := message.TemplateSubscribeMessage{
		ToUser:     openID,
		TemplateID: templateID,
		Data:       dataItem,
		URL:        url,
		Scene:      scene,
		Title:      title,
	}

	// 跳小程序
	if mini != nil {
		var ok bool
		miniAppID, ok = mini[nameMiniAppID]
		if !ok {
			return fmt.Errorf("mini flag %q not set", nameMiniAppID)
		}
		if miniAppID == "" {
			return fmt.Errorf("mini flag %q not set", nameMiniAppID)
		}

		miniPagePath, ok = mini[nameMiniPagePath]
		if !ok {
			return fmt.Errorf("mini flag %q not set", nameMiniPagePath)
		}
		if miniPagePath == "" {
			return fmt.Errorf("mini flag %q not set", nameMiniPagePath)
		}

		msg.MiniProgram = &message.MiniProgramMeta{
			AppID:    miniAppID,
			PagePath: miniPagePath,
		}
	}

	if userAgent != "" {
		client.UserAgent = userAgent
	}

	if accessToken == "" {
		accessTokenResp, err := token.GetAccessToken(appID, appSecret)
		if err != nil {
			return err
		}
		accessToken = accessTokenResp.AccessToken
	}

	if err := message.SendTemplateSubscribe(accessToken, &msg); err != nil {
		return err
	}
	fmt.Println(weixin.MessageOK)

	return nil
}