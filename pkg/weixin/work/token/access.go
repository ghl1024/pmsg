package token

import (
	"fmt"
	"time"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin"
)

type AccessTokenMeta struct {
	AccessToken string    `json:"access_token"`        // 微信接口调用凭证
	ExpireIn    int64     `json:"expires_in"`          // 微信接口调用凭证有效时间，单位：秒
	ExpireAt    time.Time `json:"expire_at,omitempty"` // 微信接口调用凭证到期时间
}

func (t AccessTokenMeta) String() string {
	if t.ExpireAt.IsZero() {
		return fmt.Sprintf("access_token: %q, expires_in: %v", t.AccessToken, t.ExpireIn)
	}
	return fmt.Sprintf("access_token: %q, expires_in: %v, expire_at: %q", t.AccessToken, t.ExpireIn, t.ExpireAt.Format(time.RFC3339))
}

// AccessTokenResponse 响应
type AccessTokenResponse struct {
	weixin.ResponseMeta
	AccessTokenMeta
}

const accessTokenURL = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"

// GetAccessToken 获取微信接口调用凭证
//
//	{
//	 "errcode": 0,
//	 "errmsg": "ok",
//	 "access_token": "accesstoken000001",
//	 "expires_in": 7200
//	}
func GetAccessToken(corpID, corpSecret string) (*AccessTokenMeta, error) {
	url := fmt.Sprintf(accessTokenURL, corpID, corpSecret)
	var resp AccessTokenResponse
	_, err := client.GetJSON(url, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Succeed() {
		return nil, fmt.Errorf("%w; %v", weixin.ErrWeiXinRequest, resp.ResponseMeta)
	}

	resp.AccessTokenMeta.ExpireAt = time.Now().Add(time.Second * time.Duration(resp.AccessTokenMeta.ExpireIn))

	return &resp.AccessTokenMeta, nil
}
