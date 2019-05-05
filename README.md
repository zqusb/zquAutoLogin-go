# zquAutoLogin-go
### 肇庆学院校园网wifi自动登录程序
电信和移动网络均适用<br>
#### 立项目的
将登录校园网的繁琐操作集成到路由器中完成<br>
所以程序将采用命令行方式运行<br>
#### 使用方法
```
Usage: zquAutoLogin-go -u [studentId] -p [password]
Options:
 -p string
        设置校园网登录密码(身份证后8位)
 -t     循环运行(按Ctrl+C结束程序)
 -u string
        设置校园网登录学号
```
#### Windows 
在程序所在文件夹打开cmd或PowerShell<br>
```
// 64位
./zquAutoLogin-go_windows_amd64.exe -u 学号 -p 密码
// 32位
./zquAutoLogin-go_windows_i386.exe -u 学号 -p 密码
```
**[便捷启动方式]**<br>
在同目录下新建start.bat文件<br>
编辑以下内容(对应位置填写你的学号密码)并保存<br>
```
:: 语句前加 rem 起注释作用，该行语句不会执行
:: 执行一次
zquAutoLogin-go_windows_amd64 -u 2016241314xx -p 身份证后8位
:: 循环执行
:: zquAutoLogin-go_windows_amd64 -u 2016241314xx -p 身份证后8位
pause
exit
```
双击start.bat即可运行程序<br>
可以创建start.bat的快捷方式到任意位置运行<br>
#### Linux/路由器
```
chmod 777 zquAutoLogin-go_xxx
./zquAutoLogin-go_xxx
```
后台运行<br>
```
nohup ./zquAutoLogin-go_xxx &
```
程序启动后，将会循环检查网络状态<br>
程序写的比较弱智，欢迎提交修改意见和bug<br>

####更新内容
[2019.5.5] 修改程序运行规则为默认**单次执行**，加-t参数可循环执行
