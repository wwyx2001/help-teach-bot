## help-teach-bot

## Day1
使用以前Linux上的服务器并[使用docker部署CoolQ HTTP API](https://cqhttp.cc/docs/4.10/#/)

但无法访问noVNC，questioning

## Day2
终于发现本机ip指的是服务器的ip！！！

从腾讯云get云服务器，操作系统Ubuntu

登录云服务器并安装docker、部署coolq、登上noVNC，此时我有了一个什么也不会的小bot

通过基础的学习[API](https://cqhttp.cc/docs/4.14/#/API?id=%E5%93%8D%E5%BA%94%E8%AF%B4%E6%98%8E)，我让这个一无所知的小bot给自己发了个消息

一开始我send出去，收获了fail，查明原因发现我的前端需要[打开开发模式](https://docs.cqp.im/dev/v9/devmode/)，然后再send，我收获了一条来自小bot的消息
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200303234415798.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0JsdWVfQ3VTTzQ=,size_16,color_FFFFFF,t_70)

## Day3
1、我在网上找了一个简单的用json接受应答消息的.go程序，发现用的是MVC架构，于是乎对MVC框架进行了解；然后可以通过postman给coolq发送指令使其send message

2、打算使用gin框架所以对gin框架的用法进行学习

3、搜索资料打算把go project跑在服务器上

## Day4
1、打算把go project跑在服务器上，故在服务器上[安装go语言环境](https://www.runoob.com/go/go-environment.html)，还有[这个资料](https://golang.google.cn/doc/install?download=go1.14.linux-amd64.tar.gz)，然后是[把golang项目部署到Linux服务器上](https://blog.csdn.net/qq_33230584/article/details/81536572)，这其中有一步是将该文件放入linux系统某个文件夹下

查询应该使用的命令是: scp local_file username@userip:remote_folder

我在本地cmd使用了命令：scp F:\helpteachbot\main ubuntu@175.24.41.84:/usr/local

但这样给我报无权限错误，我查询资料发现可以转移到别的文件夹下，然后再mv过去

所以我最后使用的命令是：scp F:\helpteachbot\main ubuntu@175.24.41.84:/tmp

至此我解决了go语言的环境

2、但这样我发现他还是无法监听到消息，明明事件上报的端口就是8080，我listen的也是8080，为什么听不到呢？questioning

3、一些笔记

postform貌似是后面的无法通过命令进行改变
而query则是可以通过post命令改变的值
如果在前面加一个Default，那么你的语句里要有一个“，”，如果是postform则一直是后面那个值恒定
而如果是query，那么没读入就是“，”后面那个，读入了就是读入的

一个非常好的[学习gin框架](https://www.cnblogs.com/-beyond/p/9391892.html)的网页

get是从浏览器上访问就行
而post需要发一个post

ShouldBindQuery(&json)//GET请求
ShouldBind(&json)//POST请求

## Day5
非常zz的一天，在找究竟是什么导致QQ不能给我回复，多亏了学长的指导。。。
[这个blog](https://blog.csdn.net/weixin_33889665/article/details/93166512)的步骤五
最后成功的配置结果
![在这里插入图片描述](https://img-blog.csdnimg.cn/20200307104733493.png)

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200307104845648.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0JsdWVfQ3VTTzQ=,size_16,color_FFFFFF,t_70)

不过也不算是闲着了，大概码了一下层次二的代码，数据库打算用MongoDB，但还没下载下来

## Day6
今天本来打算在Linux上安装MongoDB，可是出bug，用尽所有方法也解决不了

于是，我用[docker安装了MongoDB](https://www.runoob.com/docker/docker-install-mongodb.html)，不出几分钟就安好了（所以我真是ZZ啊！！！！

今晚码完层二！

贴一些很好的学习链接

[mongodb官方的golang驱动基础使用](https://www.jianshu.com/p/0344a21e8040)

[MongoDB Go Driver使用帮助文档](http://www.mongoing.com/archives/27257)

[go语言time用法总结](https://blog.csdn.net/yzf279533105/article/details/92797674?depth_1-utm_source=distribute.pc_relevant.none-task&utm_source=distribute.pc_relevant.none-task)

[Golang MongoDB bson.M查询&修改](https://blog.csdn.net/LightUpHeaven/article/details/82663146)

代码还差一个进阶，那就明天发好了。。。

## Day7
昨晚想了一下进阶定时提醒的做法：每隔1min就扫一下所有的仓库，如果达到了提前x小时的提前时间内，当时保存的时候多加一个参数表示每隔多少时间，时间过1min参数+1，当加到原数字的两倍的时候再提醒一次并变成原数字

开工
