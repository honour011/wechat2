// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechatv2 for the canonical source repository
// @license     https://github.com/chanxuehong/wechatv2/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package masstogroup

import (
	"errors"
	"net/http"

	"github.com/chanxuehong/wechatv2/mp"
)

type Client struct {
	mp.WechatClient
}

// 创建一个新的 Client.
//  如果 HttpClient == nil 则默认用 http.DefaultClient
func NewClient(TokenService mp.TokenService, HttpClient *http.Client) *Client {
	if TokenService == nil {
		panic("TokenService == nil")
	}
	if HttpClient == nil {
		HttpClient = http.DefaultClient
	}

	return &Client{
		WechatClient: mp.WechatClient{
			TokenService: TokenService,
			HttpClient:   HttpClient,
		},
	}
}

func (clt *Client) SendText(msg *Text) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return clt.send(msg)
}

func (clt *Client) SendImage(msg *Image) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return clt.send(msg)
}

func (clt *Client) SendVoice(msg *Voice) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return clt.send(msg)
}

func (clt *Client) SendVideo(msg *Video) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return clt.send(msg)
}

func (clt *Client) SendNews(msg *News) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}
	return clt.send(msg)
}

func (clt *Client) send(msg interface{}) (msgid int64, err error) {
	var result struct {
		mp.Error
		MsgId int64 `json:"msg_id"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token="
	if err = clt.PostJSON(incompleteURL, msg, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	msgid = result.MsgId
	return
}
