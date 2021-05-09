# WafHelper

## 概述

VeryNginx能开启浏览器验证，开了之后抗打能力挺强。遇到cc光靠他的限流根本顶不住。而VeryNginx不支持自动化开启浏览器验证，只能手动打开，故本程序诞生。

本程序靠80 443端口tcp连接数判断是否正在被攻击，默认小于40个连接为正常，可编辑`modules/checker.go`修改。

## 依赖

+ 安装了VerNginx
+ VeryNginx必须安装在默认位置，否则需要修改`modules/veryNginx.go`中的路径
+ 命令`nginx`在shell中可用，用于重载配置
+ 浏览器验证规则已配置并开启，本程序控制的是浏览器验证总开关

## 运行/编译

安装好go环境后

### 运行：

```shell
go run main.go
```

### 编译

```shell
go build
```

挂进后台则在命令末尾加个`&`即可，但建议使用screen

## 开机启动

写个脚本，放进crontab

脚本示例（使用screen）：

```shell
#!/bin/sh
screen_name=$"WafHelper"
screen -dmS $screen_name
screen -x -S $screen_name -p 0 -X stuff "/root/WafHelper"
screen -x -S $screen_name -p 0 -X stuff $'\n'
```

crontab示例：

```
@reboot xxx.sh
```