# cloudgo

## 运行程序
> $ go run ./web/cloudgo/main.go -p9090
[negroni] listening on :9090
[negroni] 2017-11-13T08:20:43+08:00 | 200 | 	 56.151µs | localhost:9090 | GET /hello/Feather

## 测试
### 测试命令
> curl -v http://localhost:9090/hello/Feather
### 测试结果
*   Trying ::1...
* Connected to localhost (::1) port 9090 (#0)
> GET /hello/Feather HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Mon, 13 Nov 2017 00:19:02 GMT
< Content-Length: 31
< 
{
  "Test": "Hello Feather"
}
* Connection #0 to host localhost left intact


### 压力测试命令
> ab -n 1000 -c 100 http://localhost:9090/hello/Feather

### 测试结果
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            9090

Document Path:          /hello/Feather
Document Length:        30 bytes

Concurrency Level:      100
Time taken for tests:   0.066 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      153000 bytes
HTML transferred:       30000 bytes
Requests per second:    15066.82 [#/sec] (mean)
Time per request:       6.637 [ms] (mean)
Time per request:       0.066 [ms] (mean, across all concurrent requests)
Transfer rate:          2251.19 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.5      1       2
Processing:     0    6   5.7      3      22
Waiting:        0    5   5.7      3      22
Total:          0    6   5.8      4      23

Percentage of the requests served within a certain time (ms)
  50%      4
  66%      4
  75%      8
  80%     10
  90%     19
  95%     21
  98%     22
  99%     22
 100%     23 (longest request)

具体安装、命令参数与结果解释参考:![CentOS服务器Http压力测试之ab](http://linux.it.net.cn/CentOS/fast/2015/0715/16393.html)

