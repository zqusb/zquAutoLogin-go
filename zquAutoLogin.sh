#! /bin/bash
if [ ! -f "/tmp/zquAutoLogin-go" ];then
	echo [zquAutoLogin] 程序不存在,开始下载自动登录程序...  >> /tmp/syslog.log
	wget -O /tmp/zquAutoLogin-go https://zqu-auto-login-go-1252708919.cos.ap-guangzhou.myqcloud.com/zquAutoLogin-go_linux_mipsle
	if [ $? -eq 0 ];then
		echo [zquAutoLogin] 程序下载完成,开始运行程序 >> /tmp/syslog.log
	else 
		echo [zquAutoLogin] 程序下载失败 >> /tmp/syslog.log
		exit 1
	fi
else
	echo [zquAutoLogin] 程序存在,开始运行程序 >> /tmp/syslog.log
fi
chmod 777 /tmp/zquAutoLogin-go
/tmp/zquAutoLogin-go -u $1 -p $2 >> /tmp/syslog.log
v=$(/tmp/zquAutoLogin-go -v | grep "0")
if [[ "$v" != "" ]];then
	rm /tmp/zquAutoLogin-go
	echo [zquAutoLogin] 检查到新版本，开始更新 >> /tmp/syslog.log
fi
echo [zquAutoLogin] 程序运行完毕