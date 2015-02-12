// @description wechat2 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/philsong/wechat2 for the canonical source repository
// @license     https://github.com/philsong/wechat2/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"net/url"
)

// 扫码原生支付模式1的地址
func NativeURL1(appId, mchId, productId, timestamp, nonceStr, apiKey string) string {
	m := make(map[string]string, 5)
	m["appid"] = appId
	m["mch_id"] = mchId
	m["product_id"] = productId
	m["time_stamp"] = timestamp
	m["nonce_str"] = nonceStr

	signature := Sign(m, apiKey, nil)

	return "weixin://wxpay/bizpayurl?sign=" + signature +
		"&appid=" + url.QueryEscape(appId) +
		"&mch_id=" + url.QueryEscape(mchId) +
		"&product_id=" + url.QueryEscape(productId) +
		"&time_stamp=" + url.QueryEscape(timestamp) +
		"&nonce_str=" + url.QueryEscape(nonceStr)
}

// 扫码原生支付模式2的地址
func NativeURL2(codeURL string) string {
	return "weixin://wxpay/bizpayurl?sr=" + url.QueryEscape(codeURL)
}
