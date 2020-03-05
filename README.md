## help-teach-bot

## Day1
使用以前Linux上的服务器并[使用docker部署CoolQ HTTP API](https://cqhttp.cc/docs/4.10/#/)

但无法访问noVNC，questioning

## Day2
终于发现本机ip指的是服务器的ip！！！

从腾讯云get云服务器，操作系统Ubuntu

登录云服务器并安装docker、部署coolq、登上noVNC，此时我有了一个什么也不会的小bot

通过基础的[API与postman](https://cqhttp.cc/docs/4.14/#/API?id=%E5%93%8D%E5%BA%94%E8%AF%B4%E6%98%8E)，我让这个一无所知的小bot给自己发了个消息

一开始我send出去，收获了fail，查明原因发现我的前端需要[打开开发模式](https://docs.cqp.im/dev/v9/devmode/)，然后再send，我收获了一条来自小bot的消息
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200303234415798.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0JsdWVfQ3VTTzQ=,size_16,color_FFFFFF,t_70)

## Day3
1、我在网上找了一个简单的用json接受应答消息的.go程序，发现用的是MVC架构，于是乎对MVC框架进行了解；然后可以通过goland给coolq发送指令使其send message
2、打算使用gin框架所以对gin框架的用法进行学习
3、搜索资料打算把go project跑在服务器上

## Day4
1、打算把go project跑在服务器上，故在服务器上[安装go语言环境](https://www.runoob.com/go/go-environment.html)，还有[这个资料](https://golang.google.cn/doc/install?download=go1.14.linux-amd64.tar.gz)，然后是[把golang项目部署到Linux服务器上](https://blog.csdn.net/qq_33230584/article/details/81536572)，这其中有一步是将该文件放入linux系统某个文件夹下
查询应该使用的命令是: scp local_file username@userip:remote_folder
我在本地cmd使用了命令：scp F:\helpteachbot\main ubuntu:175.24.41.84:/usr/local
但这样给我报无权限错误，我查询资料发现可以转移到别的文件夹下，然后再mv过去
所以我最后使用的命令是：scp F:\helpteachbot\main ubuntu:175.24.41.84:/tmp
至此我解决了go语言的环境
2、但这样我发现他还是无法监听到消息，明明事件上报的端口就是8080，我listen的也是8080，为什么听不到呢？questioning
