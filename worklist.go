/*
 * @Author: SUIBING112233
 * @Date: 2019-12-15 09:53:15
 * @LastEditTime: 2019-12-15 09:57:32
 * @WebSite: https://blog.icedtech.xyz
 */
package main

import (
	"runtime"
)

type PipeWork struct {
	*filePointer
	Done bool
	Err  error
}

var cpuNum int

func Prepare() {
	cpuNum = runtime.NumCPU()
}
