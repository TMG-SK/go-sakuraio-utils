// Copyright tsumug, inc.

/*
CheckSignature は、sakura.ioのoutgoing webhookで設定されたリクエストが正当なものかどうか判断します。
*/
package sakuraio

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

// sakura.ioのシグネチャを検証する
func CheckSignature(secret string, body []byte, signature string) bool {
	h := hmac.New(sha1.New, []byte(secret))
	h.Write(body)
	return signature == hex.EncodeToString(h.Sum(nil))
}
