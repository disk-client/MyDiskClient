/*
 * @Author: your name
 * @Date: 2020-06-20 10:24:29
 * @LastEditTime: 2020-06-20 17:25:25
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskClient/utils/connectPlatform.go
 */

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Login 连接远端登录接口
func Login(username, cred string) (msg string) {
	var params = map[string]string{
		"name":     username,
		"password": cred,
	}
	var jsonStr, _ = json.Marshal(params)
	var client = &http.Client{}
	var loginURL = fmt.Sprintf("http://%s:8080/login", Host)
	resp, err := client.Post(loginURL, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
		return "连接管理平台失败"
	}
	result, _ := ioutil.ReadAll(resp.Body)
	var data map[string]map[string]interface{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		return err.Error()
	}
	var content = data["data"]["info"]
	msg, _ = content.(string)
	return msg
}

// GetSelfServerAddr 获取自己私有网盘的服务端地址
func GetSelfServerAddr(username string) (msg string) {
	var params = map[string]string{
		"name": username,
		"cmd":  "ServerAddr",
	}
	var jsonStr, _ = json.Marshal(params)
	var client = &http.Client{}
	var reqURL = fmt.Sprintf("http://%s:8080/client/server/ipaddr", Host)
	resp, err := client.Post(reqURL, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
		return "连接管理平台失败"
	}
	result, _ := ioutil.ReadAll(resp.Body)
	var data map[string]map[string]interface{}
	err = json.Unmarshal(result, &data)
	if err != nil {
		return err.Error()
	}
	var content = data["data"]["info"]
	msg, _ = content.(string)
	return msg
}
