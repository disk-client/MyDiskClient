/*
 * @Author: your name
 * @Date: 2020-06-20 11:35:38
 * @LastEditTime: 2020-07-03 21:33:29
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskClient/utils/init.go
 */

package utils

import (
	"strconv"

	"gopkg.in/ini.v1"
)

// Host 管理端ip
var Host = ""

// Port 管理端打洞用的端口
var Port = 0

// Init 项目初始化
func Init() {
	iniObj, err := ini.Load("../conf/admin.ini")
	PanicErr(err)
	var sec = iniObj.Section("platform")
	Host = sec.Key("HOST").String()
	var portStr = sec.Key("PORT").String()
	Port, err = strconv.Atoi(portStr)
	PanicErr(err)
	return
}
