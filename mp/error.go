// @description wechat2 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/philsong/wechat2 for the canonical source repository
// @license     https://github.com/philsong/wechat2/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mp

import "fmt"

const (
	ErrCodeOK                = 0
	ErrCodeInvalidCredential = 40001 // access_token 过期（无效）返回这个错误
	ErrCodeTimeout           = 42001 // access_token 过期（无效）返回这个错误（maybe!!!）
)

type Error struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("errcode: %d, errmsg: %s", e.ErrCode, e.ErrMsg)
}
