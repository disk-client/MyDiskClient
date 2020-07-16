/*
 * @Author: your name
 * @Date: 2020-06-20 11:35:38
 * @LastEditTime: 2020-07-13 21:36:13
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskClient/utils/init.go
 */

package utils

// Host 管理端ip
var Host = "47.92.149.230"

// Port 管理端打洞用的端口
var Port = "8087"

// Init 项目初始化
// func Init() {
// 	iniObj, err := ini.Load("../conf/admin.ini")
// 	PanicErr(err)
// 	var sec = iniObj.Section("platform")
// 	Host = sec.Key("HOST").String()
// 	var portStr = sec.Key("PORT").String()
// 	Port, err = strconv.Atoi(portStr)
// 	PanicErr(err)
// 	return
// }
