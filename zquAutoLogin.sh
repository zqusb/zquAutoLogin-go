#! /bin/bash
ping -c 1 qq.com > /dev/null 2>&1

if [ $? -ne 0 ];then
	echo 检测网络连接异常
    if [ ! -f "/tmp/zquAutoLogin-go" ];then
        echo [zquAutoLogin] 程序不存在,且无法联网
        exit 1
    fi
    chmod 777 /tmp/zquAutoLogin-go
    /tmp/zquAutoLogin-go -u $1 -p $2
fi

echo [zquAutoLogin] 网络状态正常
echo [zquAutoLogin] 学号:$1 密码:$2

if [ ! -f "/tmp/zquAutoLogin-go" ];then
	echo [zquAutoLogin] 程序不存在,开始下载自动登录程序...
	wget -O /tmp/zquAutoLogin-go https://zqu-auto-login-go-1252708919.cos.ap-guangzhou.myqcloud.com/zquAutoLogin-go-v1.1
	echo [zquAutoLogin] 程序下载完成
else
	echo [zquAutoLogin] 程序存在,开始运行程序
fi
    echo [zquAutoLogin] 程序运行完毕

chmod 777 /tmp/zquAutoLogin-go
/tmp/zquAutoLogin-go -u $1 -p $2
