/*
 * @Author: your name
 * @Date: 2020-06-20 17:37:54
 * @LastEditTime: 2020-07-24 16:38:26
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskClient/conf/params.go
 */

package conf

// DecryptSegLen 解密时使用的一次循环下的解密的长度
const DecryptSegLen = 128

// CryptSegLen 加密时使用的一次循环下的解密的长度
const CryptSegLen = 100

// DecryptB64SegLen 解密base64时一次的长度
const DecryptB64SegLen = 172

// LogLev 日志等级
var LogLev = map[string]int{
	"DEBUG": 0,
	"INFO":  1,
	"WARN":  2,
	"ERROR": 3,
	"CRIT":  4,
}

// TargetAddr 目标地址
const TargetAddr = "47.92.149.230:8087"

// AesKey 对称加密密钥
var AesKey = []byte("0f90023fc9ae101e")
