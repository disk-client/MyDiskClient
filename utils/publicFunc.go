/*
 * @Author: your name
 * @Date: 2020-06-20 11:38:01
 * @LastEditTime: 2020-06-20 17:41:27
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskClient/utils/publicFunc.go
 */

package utils

import "os"

// PanicErr 异常打断程序
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckFileIsExist 检查文件是否存在
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
