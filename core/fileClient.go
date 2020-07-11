/*
 * @Author: your name
 * @Date: 2020-07-11 16:46:29
 * @LastEditTime: 2020-07-11 17:27:22
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskClient/core/fileClient.go
 */

package core

import (
	"MyDiskClient/conf"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetDirList 获取文件列表
func GetDirList(path string) {
	var urlPath = "/FileMenu"
	var postData = map[string]string{
		"path": path,
	}
	jsonstr, err := json.Marshal(postData)
	if err != nil {
	}
	buffer := bytes.NewBuffer(jsonstr)
	req, err := http.NewRequest("POST", conf.TargetAddr+urlPath, buffer)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8") //添加请求头
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		resp.Body.Close()
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
	}
}
