# autossh

go写的一个ssh远程客户端。可一键登录远程服务器，主要用来弥补Mac/Linux Terminal ssh无法保存密码的不足。

## 版本

v0.21

## 下载
[https://github.com/lorock/autossh/releases](https://github.com/lorock/autossh/releases)

## 配置
下载编译好的二进制包autossh，复制server.example.json到家目录下为.servers.json文件。

```
cp server.example.json ~/.server.json

```
编辑servers.json，内容可参考server.example.json

```json
[
  {
    "name": "LorockOPS", // 显示名称
    "ip": "192.168.33.10", // 服务器IP或域名
    "port": 22, // 端口
    "user": "root", // 用户名
    "password": "LorockOPS", // 密码
    "method": "password" // 认证方式，目前支持password和pem
  },
  {
    "name": "ssh-pem",
    "ip": "192.168.33.11",
    "port": 22,
    "user": "root",
    "password": "your pem file password or empty", // pem密钥密码，若无密码则留空
    "method": "pem", // pem密钥认证
    "key": "your pem file path" // pem密钥文件绝对路径
  }
  // ...可配置多个
]
```
保存.servers.json，执行autossh，即可看到相应服务器信息，输入对应序号，即可自动登录到服务器
![登录演示](doc/images/demo.gif)

## 高级用法
设置alias，可在任意目录下调用
```bash
[root@localhost ~]# vim /etc/profile
在行尾追加 alias autossh="~/autossh_path/autossh"
[root@localhost ~]# . /etc/profile
```
更多快捷操作，可调用 `--help` 查看
```bash
[root@localhost autossh]# autossh --help
go写的一个ssh远程客户端。可一键登录远程服务器，主要用来弥补Mac/Linux Terminal ssh无法保存密码的不足。
基本用法：
  直接输入autossh不带任何参数，列出所有服务器，输入对应编号登录。
参数：
  -v, --version 	 显示 Lorock Autossh 的版本信息。
  -h, --help    	 显示帮助信息。
操作：
  list          	 显示所有server。
  add <name>    	 添加一个 server。如：autossh add Lorockops。
  edit <name>   	 编辑一个 server。如：autossh edit Lorockops。
  remove <name> 	 删除一个 server。如：autossh remove Lorockops。
```

## Q&amp;A
- Q: Downloads中为什么没有Windows的包？
- A: Windows下有很多优秀的ssh工具，autossh主要面向Mac/Linux群体。

- Q: 为什么要设置alias而不将autossh放到/usr/bin/下？
- A: autossh核心文件有两个，autossh和~/.servers.json 通过alias调用,经过Lorock的努力目前已经可以将autossh放到/usr/bin/目录下直接调用。

## 编译
go build main.go

## 依赖包
- golang.org/x/crypto/ssh

## TODO
- [x] -v, --version 查看版本号
- [x] -h, --help 显示帮助
- [x] list 显示所有server
- [x] add 添加一个server
- [x] remove name 删除一个server
- [x] edit name 编辑一个server

