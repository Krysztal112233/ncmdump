ncmdump
===

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
Linux
=====
首先，您只需安装 Golang 环境，随后键入（若处中国大陆，则需网络代理）

```shell
go get -v github.com/SUIBING112233/ncmdump
```

而后切换到本项目的 cmd 文件夹以键入

```shell
go build ncmdump.go
```
即可在当前目录输出编译完成的程序。之后，您可将其置于系统 PATH 中，以备不时之需。

用法：

```shell
ncmdump [文件名...]
```
Windows
=====
相较于 Linux，Windows 下安装显得更快捷。您只需选择 .NCM 文件，拖拽至 ncmdump.exe，即可转换文件至同目录下。

特别声明
====
本软件仅供学习交流，严禁用于商业用途，并请于下载时始的 24 小时内删除。
