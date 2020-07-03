/*
 * @Author: your name
 * @Date: 2020-07-03 21:31:58
 * @LastEditTime: 2020-07-03 21:53:57
 * @LastEditors: Please set LastEditors
 * @Description: 初始时建立链接
 * @FilePath: /MyDiskClient/core/initConnect.go
 */

package core

import (
	"MyDiskClient/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
)

// Login 登陆客户端
func login(username, cred string) (msg string) {
	var params = map[string]string{
		"name":     username,
		"password": cred,
	}
	var jsonStr, _ = json.Marshal(params)
	var client = &http.Client{}
	var loginURL = fmt.Sprintf("http://%s:8080/login", utils.Host)
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

func connectControl(username string) error {
	var tcpAddr *net.TCPAddr
	var addr = utils.Host + ":" + strconv.Itoa(utils.Port)
	//这里在一台机测试，所以没有连接到公网，可以修改到公网ip
	tcpAddr, _ = net.ResolveTCPAddr("tcp", addr)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return err
	}
	conn.Write([]byte(username))
	fmt.Println(conn.LocalAddr().String() + " : Client connected success!")
	return nil
}

// Lighthouse 登陆代码
func Lighthouse(username, cred string) (flag bool) {
	flag = false
	var msg = login(username, cred)
	if msg == "登录成功" {
		if err := connectControl(username); err == nil {
			flag = true
		}
	}
	return
}
