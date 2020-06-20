/*
 * @Author: your name
 * @Date: 2020-06-20 11:38:01
 * @LastEditTime: 2020-06-20 11:38:27
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /MyDiskClient/utils/publicFunc.go
 */

package utils

// PanicErr 异常打断程序
func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
