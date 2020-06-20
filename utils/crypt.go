/*
 * @Author: xiaoboya
 * @Date: 2020-06-17 16:15:46
 * @LastEditTime: 2020-06-20 17:40:34
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /SelfDisk/utils/crypt.go
 */

package utils

import (
	"MyDiskClient/conf"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// PriKey 私钥，与前端加解密相关
var PriKey = new(rsa.PrivateKey)

// PubKey 公钥，与前端加解密相关
var PubKey = new(rsa.PublicKey)

// CryptMsg 加密数据
func CryptMsg(msgIn string) string {
	if PubKey == nil {
		return ""
	}
	var msg = url.QueryEscape(msgIn)
	var enCryptMsg = ""
	var i = 0
	var tLen = len(msg)
	for i < tLen {
		var content []byte
		if tLen > i+conf.CryptSegLen {
			content = []byte(msg[i : i+conf.CryptSegLen])
		} else {
			content = []byte(msg[i:tLen])
		}
		encryptBytes, err := rsa.EncryptPKCS1v15(rand.Reader, PubKey, content)
		PanicErr(err)
		enCryptMsg += base64.StdEncoding.EncodeToString(encryptBytes)
		i += conf.CryptSegLen
	}
	return enCryptMsg
}

// DecryptMsg 解密数据
func DecryptMsg(msgIn string, key int) string {
	var k = new(rsa.PrivateKey)
	if PriKey == nil {
		return ""
	}
	k = PriKey
	var quoteMsg = ""
	var tLen = len(msgIn)
	var i = 0
	msgIn = strings.Replace(msgIn, " ", "", -1)
	for i < tLen {
		enc, err := base64.StdEncoding.DecodeString(msgIn[i : i+conf.DecryptB64SegLen])
		PanicErr(err)
		res, err := rsa.DecryptPKCS1v15(rand.Reader, k, enc)
		PanicErr(err)
		quoteMsg += string(res)
		i += conf.DecryptB64SegLen
	}
	result, err := url.QueryUnescape(quoteMsg)
	PanicErr(err)
	return result
}

// ReqParse 接收请求，解密request
func ReqParse(c *gin.Context, obj interface{}) {
	// body, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	c.String(400, err.Error())
	// 	c.Abort()
	// }
	// c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	var msg = c.PostForm("msg")
	// var msg = string(body)
	msg = DecryptMsg(msg, 1)
	var err = json.Unmarshal([]byte(msg), obj)
	PanicErr(err)
}

// InitKeys 初始化公钥和私钥
func InitKeys() {
	var f []byte
	var err error

	// 生成解密用的私钥
	if CheckFileIsExist("../cert/fontPrivate.pem") {
		f, err = ioutil.ReadFile("../cert/fontPrivate.pem")
		PanicErr(err)
	} else {
		return
	}
	block, _ := pem.Decode(f)
	if block == nil {
		return
	}
	PriKey, _ = x509.ParsePKCS1PrivateKey(block.Bytes)
	if PriKey == nil {
		return
	}

	// 生成加密用的公钥
	if CheckFileIsExist("../cert/backendPublic.pem") {
		f, err = ioutil.ReadFile("../cert/backendPublic.pem")
		PanicErr(err)
	} else {
		return
	}
	block, _ = pem.Decode(f)
	if block == nil {
		return
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if pubKey == nil {
		PanicErr(err)
		return
	}
	value, ok := pubKey.(*rsa.PublicKey)
	if ok == true {
		PubKey = value
	} else {
		return
	}
}
