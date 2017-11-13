# cloudgo


## 运行程序
> $ go run ./web/cloudgo/main.go -p9090

[negroni] listening on :9090

[negroni] 2017-11-13T08:20:43+08:00 | 200 | 	 56.151µs | localhost:9090 | GET /hello/Feather

## 使用框架
gin


## 测试
### 测试命令
> curl -v http://localhost:9090/hello/Feather


### 测试结果
![](https://github.com/FlyingFeather/cloudgo/blob/master/picture/c-2.png)


### 压力测试命令
> ab -n 1000 -c 100 http://localhost:9090/hello/Feather


### 测试结果
![](https://github.com/FlyingFeather/cloudgo/blob/master/picture/c-3.png)

具体安装、命令参数与结果解释参考:![CentOS服务器Http压力测试之ab](http://linux.it.net.cn/CentOS/fast/2015/0715/16393.html)

