/*
 * @Author: your name
 * @Date: 2020-07-11 16:46:29
 * @LastEditTime: 2020-07-13 21:57:48
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskClient/core/fileClient.go
 */

package core

import (
	"MyDiskClient/conf"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendPostRequest(url string, params interface{}) (resp *http.Response) {
	jsonstr, err := json.Marshal(params)
	if err != nil {
		return nil
	}
	buffer := bytes.NewBuffer(jsonstr)
	req, err := http.NewRequest("POST", "http://"+conf.TargetAddr+url, buffer)
	if err != nil {
		fmt.Println(err)
		resp.Body.Close()
		fmt.Println("connect error")
		return nil
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8") //添加请求头
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		resp.Body.Close()
		fmt.Println("connect error2")
		return nil
	}
	return
}

// SimpleResponse 普通的返回值
type SimpleResponse struct {
	Msg  string `json:"msg"`
	Succ bool   `json:"succ"`
}

// DirPerm 权限
type DirPerm struct {
	Own   string `json:"own"`
	Group string `json:"group"`
	Other string `json:"owner"`
}

// DirInfo 详情
type DirInfo struct {
	FileName string  `json:"filename"`
	FileType int     `json:"filetype"`
	FilePerm DirPerm `json:"fileperm"`
}

// NullResponse 空请求
type NullResponse struct {
	Valid    bool
	Response interface{}
}

// GetDirList 获取文件列表
func GetDirList(path string) (res NullResponse) {
	res.Valid = false
	var urlPath = "/FileMenu"
	var postData = map[string]string{
		"path": path,
	}
	var resp = sendPostRequest(urlPath, postData)
	fmt.Println(1111)
	if resp == nil {
		fmt.Println("no……")
		return
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(respBytes))
	var result DirInfo
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		fmt.Println("Unmarshal err")
		return
	}
	res.Response = result
	res.Valid = true
	return
}

// CreateDir 创建文件夹
func CreateDir(path string) (res NullResponse) {
	res.Valid = false
	var urlPath = "/NewDir"
	var postData = map[string]string{
		"path": path,
	}
	var resp = sendPostRequest(urlPath, postData)
	if resp == nil {
		return
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	var result SimpleResponse
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return
	}
	res.Response = result
	res.Valid = true
	return
}

// RenameDir 重命名文件或文件夹
func RenameDir(path, newname string) (res NullResponse) {
	res.Valid = false
	var urlPath = "/Rename"
	var postData = map[string]string{
		"oldpath": path,
		"newfile": newname,
	}
	var resp = sendPostRequest(urlPath, postData)
	if resp == nil {
		return
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	var result SimpleResponse
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return
	}
	res.Response = result
	res.Valid = true
	return
}
