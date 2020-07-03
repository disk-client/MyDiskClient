/*
 * @Author: xiaoboya
 * @Date: 2020-06-19 16:03:52
 * @LastEditTime: 2020-07-03 21:48:47
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskClient/main.go
 */

package main

import (
	"MyDiskClient/utils"
	"net"
)

// SRCADDR UDP服务路径
var SRCADDR = &net.UDPAddr{IP: net.IPv4zero, Port: 7611}

func main() {
	utils.Init()
	// var msg = core.Login("xiaoboya", "root1234")
	// fmt.Println(msg)
}
