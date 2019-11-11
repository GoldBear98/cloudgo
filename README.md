**1、概述**
开发简单web服务程序cloudgo，了解web服务器工作原理。
任务目标：
（1）熟悉go服务器工作原理

（2）基于现有web库，编写一个简单web应用类似cloudgo

（3）使用curl工具访问web程序

（4）对web执行压力测试

**2、任务要求**
基本要求：
（1）编程web服务程序 类似cloudgo应用。
	要求有详细的注释
	是否使用框架、选哪个框架自己决定请在README.md 说明你决策的依据
	
（2）使用 curl 测试，将测试结果写入 README.md

（3）使用 ab 测试，将测试结果写入 README.md。并解释重要参数。

**3.实验过程**
这次实验我选择了使用框架，选用的框架是martini。之所以选择 martini框架是因为martini框架是使用Go语言作为开发语言的一个强力的快速构建模块化web应用与服务的开发框架。martini 是一个新锐的框架，概念非常不错。虽然martini只是一个微型框架，只带有简单的核心、路由功能和依赖注入容器inject，但是它还是有很好的发展前景的。

选好了框架之后，就要开始实现main.go和service.go这两部分了：
main.go

```
package main

import (
    "os"
    "service"
     flag "github.com/spf13/pflag"
)

const (
    PORT string = "8080" /*设置默认的端口为8080*/
)

func main() {
    port := os.Getenv("PORT") 
    if len(port) == 0 {
        port = PORT
    }
    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")/*设置端口*/
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }
     service.NewServer(port)/*启动服务器*/
}
```

service.go

```
package service

import (
    "github.com/go-martini/martini" 
)

func NewServer(port string) {   
    app := martini.Classic()
    app.Get("/hello/:name", func(params martini.Params) string {
        return "Hello " + params["name"] + " !"
    })
    app.RunOnAddr(":"+port)   
}
```

注意：在编译main.go时会报错：缺少包go-martini/martini 
解决办法就是在GitHub上下载缺少的包，输入指令：

```
git clone https://github.com/go-martini/martini
```
下载好martini包后还报错：缺少包codegangsta/inject
解决办法就是在GitHub上下载缺少的包，输入指令：
```
git clone https://github.com/codegangsta/inject
```

终于不再有报错出现，可以运行main.go了：![在这里插入图片描述](https://img-blog.csdnimg.cn/20191111232059936.png)![在这里插入图片描述](https://img-blog.csdnimg.cn/20191111230935455.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0dvbGRCZWFyOTg=,size_16,color_FFFFFF,t_70)
可以看到实验中我选取的端口是5050，在main文件夹下输入测试指令`go run main.go -p5050`后，打开Google Chrome浏览器，输入网址`http://localhost:5050/hello/goldbear98`可以看到出现了想要的结果——Hello goldbear98 !

接下来使用curl测试：
打开第二个终端，输入指令：

```
curl -v http://localhost:5050/hello/goldbear98
```
得到的结果如下所示：![在这里插入图片描述](https://img-blog.csdnimg.cn/20191112000858536.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0dvbGRCZWFyOTg=,size_16,color_FFFFFF,t_70)
在第一个终端下的结果如下所示：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191112001050526.png)
说明可以监听到，test测试成功。

最后进行ab测试：
打开第三个终端，输入指令：

```
ab -n 10000 -c 100 http://localhost:5050/hello/goldbear98
```
发现出现报错：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191112001427218.png)
在搜索解决办法后找到如下解决办法：
输入指令`yum -y install httpd-tools`即可
![在这里插入图片描述](https://img-blog.csdnimg.cn/2019111122134663.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0dvbGRCZWFyOTg=,size_16,color_FFFFFF,t_70)![在这里插入图片描述](https://img-blog.csdnimg.cn/20191111221550520.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0dvbGRCZWFyOTg=,size_16,color_FFFFFF,t_70)
可以看到安装完成了，接下来我们再试一次刚才的指令：
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191112001933425.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0dvbGRCZWFyOTg=,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/2019111200200822.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L0dvbGRCZWFyOTg=,size_16,color_FFFFFF,t_70)
经过ab测试，可以看到发送了10000个请求，每一个请求花费时间为0.281ms，50%的请求需要26ms，100%的请求需要77ms。

最后对重要参数进行解释：
Server Hostname    服务器主机名
Server Port	服务器端口
Document Path	文件路径
Document Length	文件大小
Concurrency Level	并发等级
Requst per second	平均每秒的请求个数。
Time per request	用户平均的等待时间。
Connection Times	表内描述了所有的过程中所消耗的最小、中位、最长时间。
Percentage of the requests served within a certain time	每个百分段的请求完成所需的时间

[我的博客](https://github.com/GoldBear98/cloudgo)
