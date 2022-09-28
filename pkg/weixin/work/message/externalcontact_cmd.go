package message

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
	"github.com/lenye/pmsg/pkg/weixin/work/token"
)

type CmdWorkSendExternalContactParams struct {
	UserAgent              string
	AccessToken            string
	CorpID                 string
	CorpSecret             string
	RecvScope              int
	ToParentUserID         []string
	ToStudentUserID        []string
	ToParty                []string
	ToAll                  int
	MsgType                string
	AgentID                int64
	EnableIDTrans          int
	EnableDuplicateCheck   int
	DuplicateCheckInterval int
	Data                   string
}

func (t *CmdWorkSendExternalContactParams) Validate() error {
	if t.AccessToken == "" && t.CorpID == "" {
		return flags.ErrWeixinWorkAccessToken
	}

	if t.ToParentUserID != nil && len(t.ToParentUserID) > 1000 {
		return fmt.Errorf("%v supports up to 1000", flags.ToParentUserID)
	}
	if t.ToStudentUserID != nil && len(t.ToStudentUserID) > 1000 {
		return fmt.Errorf("%v supports up to 1000", flags.ToStudentUserID)
	}
	if t.ToParty != nil && len(t.ToParty) > 100 {
		return fmt.Errorf("%v supports up to 100", flags.ToParty)
	}

	if t.RecvScope != 0 && t.RecvScope != 1 && t.RecvScope != 2 {
		return fmt.Errorf("invalid %v", flags.RecvScope)
	}
	if t.ToAll != 0 && t.ToAll != 1 {
		return fmt.Errorf("invalid %v", flags.ToAll)
	}
	if t.EnableIDTrans != 0 && t.EnableIDTrans != 1 {
		return fmt.Errorf("invalid %v", flags.EnableIDTrans)
	}
	if t.EnableDuplicateCheck != 0 && t.EnableDuplicateCheck != 1 {
		return fmt.Errorf("invalid %v", flags.EnableDuplicateCheck)
	}
	if t.DuplicateCheckInterval <= 0 || t.DuplicateCheckInterval > 3600*4 {
		return fmt.Errorf("invalid %v", flags.DuplicateCheckInterval)
	}

	if err := ValidateExternalContactMsgType(t.MsgType); err != nil {
		return fmt.Errorf("invalid flags %s: %v", flags.MsgType, err)
	}

	return nil
}

// CmdWorkSendExternalContact 发送企业微信互联企业消息
func CmdWorkSendExternalContact(arg *CmdWorkSendExternalContactParams) error {

	if err := arg.Validate(); err != nil {
		return err
	}

	msg := ExternalContactMessage{
		RecvScope:              arg.RecvScope,
		ToParentUserID:         arg.ToParentUserID,
		ToStudentUserID:        arg.ToStudentUserID,
		ToParty:                arg.ToParty,
		ToAll:                  arg.ToAll,
		MsgType:                arg.MsgType,
		AgentID:                arg.AgentID,
		EnableIDTrans:          arg.EnableIDTrans,
		EnableDuplicateCheck:   arg.EnableDuplicateCheck,
		DuplicateCheckInterval: arg.DuplicateCheckInterval,
	}

	buf := bytes.NewBufferString("")
	buf.WriteString(arg.Data)
	switch arg.MsgType {
	case ExternalContactMsgTypeText:
		var msgMeta TextMeta
		msgMeta.Content = buf.String()
		msg.Text = &msgMeta
	case ExternalContactMsgTypeImage:
		var msgMeta ImageMeta
		msgMeta.MediaID = buf.String()
		msg.Image = &msgMeta
	case ExternalContactMsgTypeVoice:
		var msgMeta VoiceMeta
		msgMeta.MediaID = buf.String()
		msg.Voice = &msgMeta
	case ExternalContactMsgTypeVideo:
		var msgMeta VideoMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		if msgMeta.MediaID == "" {
			return errors.New("media_id is empty")
		}
		msg.Video = &msgMeta
	case ExternalContactMsgTypeFile:
		var msgMeta FileMeta
		msgMeta.MediaID = buf.String()
		msg.File = &msgMeta
	case ExternalContactMsgTypeNews:
		var msgMeta NewsMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		lenArticles := len(msgMeta.Articles)
		if lenArticles == 0 || lenArticles > 8 {
			return errors.New("length of articles is 1-8")
		}
		msg.News = &msgMeta
	case ExternalContactMsgTypeMpNews:
		var msgMeta MpNewsMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		lenArticles := len(msgMeta.Articles)
		if lenArticles == 0 || lenArticles > 8 {
			return errors.New("length of articles is 1-8")
		}
		msg.MpNews = &msgMeta
	case ExternalContactMsgTypeMiniProgramNotice:
		var msgMeta MiniProgramNoticeMeta
		if err := json.Unmarshal(buf.Bytes(), &msgMeta); err != nil {
			return fmt.Errorf("invalid json format, %v", err)
		}
		lenArticles := len(msgMeta.ContentItem)
		if lenArticles > 10 {
			return errors.New("content_item up to 10")
		}
		msg.MiniProgramNotice = &msgMeta
	}

	client.UserAgent = arg.UserAgent

	if arg.AccessToken == "" {
		accessTokenResp, err := token.GetAccessToken(arg.CorpID, arg.CorpSecret)
		if err != nil {
			return err
		}
		arg.AccessToken = accessTokenResp.AccessToken
	}

	if resp, err := SendExternalContact(arg.AccessToken, &msg); err != nil {
		return err
	} else {
		fmt.Println(fmt.Sprintf("%v; %v", weixin.MessageOK, resp))
	}

	return nil
}