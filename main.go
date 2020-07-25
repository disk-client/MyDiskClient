/*
 * @Author: xiaoboya
 * @Date: 2020-06-19 16:03:52
 * @LastEditTime: 2020-07-25 11:31:27
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskClient/main.go
 */

package main

import (
	"MyDiskClient/core"
	"MyDiskClient/utils"
	"fmt"
	"time"
)

func main() {
	utils.InitKeys()
	var msg = core.Lighthouse("xiaoboya", "root1234")
	fmt.Println(msg)
	time.Sleep(10 * time.Second)
	var v = core.GetDirList("./")
	fmt.Println(v)
}
