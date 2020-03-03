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
