<!--
 * @Author: SUIBING112233
 * @Date: 2019-11-24 02:00:18
 * @LastEditTime : 2020-01-19 12:46:50
 * @WebSite: https://blog.icedtech.xyz
 -->
ncmdump
===

[![CircleCI](https://circleci.com/gh/SUIBING112233/ncmdump.svg?style=svg)](https://circleci.com/gh/SUIBING112233/ncmdump)

介绍
==== 
此软件遵守 GPLv3 进行分发，与依赖所用 Apache 协议互相兼容。

该软件是使用 Golang 实现的 .NCM 格式文件转换工具，依赖了 "github.com/yoki123/ncmdump" 项目。

特性
====
- 编译与转换都快速、简单。
- 其支持超多文件队列转换。
- 自动跳过已转换文件。
- 静默转换。

安装&使用方法
====
Linux/macOS/Other UNIX
=====
首先，您只需安装 Golang 环境，随后键入（若处中国大陆，则需网络代理或者配置GOPROXY）

```shell
go get -v github.com/SUIBING112233/ncmdump
```

而后切换到本项目文件夹键入

```shell
go build -v
```
即可在当前目录输出编译完成的程序。之后，您可将其置于系统 PATH 中，以备不时之需。

用法：

```shell
ncmdump [文件名...]
```
Windows
=====
相较于 Linux，Windows 下使用ncmdump显得更快捷。您只需选择 .ncm 文件，拖拽至 ncmdump.exe，即可转换文件至同目录下。

特别声明
====
本软件仅供学习交流，严禁用于任何商业用途。
