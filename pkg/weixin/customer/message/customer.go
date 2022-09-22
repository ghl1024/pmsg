package message

import (
	"fmt"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
)

// 小程序 msgtype 的合法值
const (
	MiniProgramMsgTypeText            = "text"            // 文本消息
	MiniProgramMsgTypeImage           = "image"           // 图片消息
	MiniProgramMsgTypeLink            = "link"            // 图文链接
	MiniProgramMsgTypeMiniProgramPage = "miniprogrampage" // 小程序卡片
)

// ValidateMiniProgramMsgType 验证
func ValidateMiniProgramMsgType(v string) error {
	switch v {
	case MiniProgramMsgTypeText, MiniProgramMsgTypeImage, MiniProgramMsgTypeLink, MiniProgramMsgTypeMiniProgramPage:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q]", v, MiniProgramMsgTypeText, MiniProgramMsgTypeImage, MiniProgramMsgTypeLink, MiniProgramMsgTypeMiniProgramPage)
	}
	return nil
}

// 公众号 msgtype 的合法值
const (
	MpMsgTypeText            = "text"            // 文本消息
	MpMsgTypeImage           = "image"           // 图片消息
	MpMsgTypeVoice           = "voice"           // 语音消息
	MpMsgTypeVideo           = "video"           // 视频消息
	MpMsgTypeMusic           = "music"           // 音乐消息
	MpMsgTypeNews            = "news"            // 图文消息（点击跳转到外链）
	MpMsgTypeMpNews          = "mpnews"          // 图文消息（点击跳转到图文消息页面）
	MpMsgTypeMpNewsArticle   = "mpnewsarticle"   // 图文消息（点击跳转到图文消息页面）使用通过 “发布” 系列接口得到的 article_id
	MpMsgTypeMsgMenu         = "msgmenu"         // 菜单消息
	MpMsgTypeWxCard          = "wxcard"          // 卡券
	MpMsgTypeMiniProgramPage = "miniprogrampage" // 小程序卡片（要求小程序与公众号已关联）
)

// ValidateMpMsgType 验证
func ValidateMpMsgType(v string) error {
	switch v {
	case MpMsgTypeText, MpMsgTypeImage, MpMsgTypeVoice, MpMsgTypeVideo, MpMsgTypeMusic, MpMsgTypeNews, MpMsgTypeMpNews, MpMsgTypeMpNewsArticle, MpMsgTypeMsgMenu, MpMsgTypeWxCard, MpMsgTypeMiniProgramPage:
	default:
		return fmt.Errorf("%s not in [%q %q %q %q %q %q %q %q %q %q %q]", v, MpMsgTypeText, MpMsgTypeImage, MpMsgTypeVoice, MpMsgTypeVideo, MpMsgTypeMusic, MpMsgTypeNews, MpMsgTypeMpNews, MpMsgTypeMpNewsArticle, MpMsgTypeMsgMenu, MpMsgTypeWxCard, MpMsgTypeMiniProgramPage)
	}
	return nil
}

// CustomerMessage 微信客服消息
type CustomerMessage struct {
	ToUser          string               `json:"touser"`
	MsgType         string               `json:"msgtype"`
	CustomService   *ServiceMeta         `json:"customservice,omitempty"`
	Text            *TextMeta            `json:"text,omitempty"`
	Image           *ImageMeta           `json:"image,omitempty"`
	Voice           *VoiceMeta           `json:"voice,omitempty"`
	Video           *VideoMeta           `json:"video,omitempty"`
	Music           *MusicMeta           `json:"music,omitempty"`
	News            *NewsMeta            `json:"news,omitempty"`
	MpNews          *MpNewsMeta          `json:"mpnews,omitempty"`
	MpNewsArticle   *MpNewsArticleMeta   `json:"mpNewsArticle,omitempty"`
	MsgMenu         *MsgMenuMeta         `json:"msgmenu,omitempty"`
	WxCard          *WxCardMeta          `json:"wxcard,omitempty"`
	MiniProgramPage *MiniProgramPageMeta `json:"miniprogrampage,omitempty"`
	Link            *LinkMeta            `json:"link,omitempty"`
}

// ServiceMeta 客服帐号
// 如果需要以某个客服帐号来发消息（在微信6.0.2及以上版本中显示自定义头像），则需在 JSON 数据包的后半部分加入 customservice 参数
type ServiceMeta struct {
	KfAccount string `json:"kf_account"`
}

// TextMeta 文本
type TextMeta struct {
	Content string `json:"content"`
}

// ImageMeta 图片
type ImageMeta struct {
	MediaID string `json:"media_id"`
}

// VoiceMeta 语音
type VoiceMeta struct {
	MediaID string `json:"media_id"`
}

// VideoMeta 视频
type VideoMeta struct {
	MediaID      string `json:"media_id"`
	ThumbMediaID string `json:"thumb_media_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

// MusicMeta 音乐
type MusicMeta struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	MusicUrl     string `json:"musicurl"`
	HqmusicUrl   string `json:"hqmusicurl"`
	ThumbMediaID string `json:"thumb_media_id"`
}

// NewsMeta 图文（点击跳转到外链）
type NewsMeta struct {
	Articles []Article `json:"articles"`
}

// Article 图文内容（点击跳转到外链）
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}

// MpNewsMeta 图文（点击跳转到图文消息页面）
type MpNewsMeta struct {
	MediaID string `json:"media_id"`
}

// MpNewsArticleMeta 图文消息（点击跳转到图文消息页面）使用通过 “发布” 系列接口得到的 article_id
type MpNewsArticleMeta struct {
	ArticleID string `json:"article_id"`
}

// MsgMenuMeta 菜单消息
type MsgMenuMeta struct {
	HeadContent string        `json:"head_content"`
	List        []MsgMenuItem `json:"list"`
	TailContent string        `json:"tail_content"`
}

// MsgMenuItem 菜单内容
type MsgMenuItem struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// WxCardMeta 卡券消息 特别注意客服消息接口投放卡券仅支持非自定义 Code 码和导入 code 模式的卡券的卡券
type WxCardMeta struct {
	CardID string `json:"card_id"`
}

// MiniProgramPageMeta 小程序卡片（要求小程序与公众号已关联）
type MiniProgramPageMeta struct {
	Title        string `json:"title"`
	AppID        string `json:"appid,omitempty"` // 小程序发送不需要填写
	PagePath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}

// LinkMeta 图文链接
type LinkMeta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	ThumbUrl    string `json:"thumb_url"`
}

const customerSendURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"

// SendCustomer 发送微信客服消息
func SendCustomer(accessToken string, msg *CustomerMessage) error {
	url := fmt.Sprintf(customerSendURL, accessToken)
	var resp weixin.ResponseMeta
	_, err := client.PostJSON(url, msg, &resp)
	if err != nil {
		return err
	}
	if !resp.Succeed() {
		return fmt.Errorf("%w; %v", weixin.ErrWeiXinRequest, resp)
	}
	return nil
}