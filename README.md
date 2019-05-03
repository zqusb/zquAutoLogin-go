# zquAutoLogin-go
### 肇庆学院校园网wifi自动登录程序
电信和移动网络均适用<br>
#### 立项目的
将登录校园网的繁琐操作集成到路由器中完成<br>
所以程序将采用命令行方式运行<br>
#### 运行
```
Usage: zquAutoLogin-go -u [studentId] -p [password]
Options:
 -p string
        设置校园网登录密码(身份证后8位)
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
