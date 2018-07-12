## 测试环境

MacBook Air (11-inch, Early 2015)   
处理器：1.6 GHz Intel Core i5   
内存：4 GB 1600 MHz DDR3   

## 测试结果

type            |0      |10ms   |1000ms
----------------|-------|-------|------
native(req/s)   |44975  |7274   |92
iris(req/s)     |43370  |7476   |92
gin(req/s)      |41818  |7754   |92

## 测试方法

分别启动iris和native，依次指定sleep参数（模拟业务逻辑耗时）：0->10ms->1000ms，并利用`wrk`
工具并发请求测试，16线程+100并发，持续30s：

```
############## 0ms
➜  wrk git:(master) ./wrk -t16 -c100 -d30s http://localhost:8080/
Running 30s test @ http://localhost:8080/
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.21ms    1.61ms  55.28ms   91.94%
    Req/Sec     2.83k   439.81    11.35k    81.17%
  1353902 requests in 30.10s, 165.27MB read
Requests/sec:  44975.45
Transfer/sec:      5.49MB

➜  wrk git:(master) ./wrk -t16 -c100 -d30s http://localhost:8000/
Running 30s test @ http://localhost:8000/
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.38ms    2.18ms  73.21ms   92.96%
    Req/Sec     2.73k   490.43    19.03k    79.21%
  1305213 requests in 30.09s, 159.33MB read
Requests/sec:  43370.81
Transfer/sec:      5.29MB

############## 10ms
➜  wrk git:(master) ./wrk -t16 -c100 -d30s http://localhost:8080/
Running 30s test @ http://localhost:8080/
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    13.14ms    2.13ms  57.03ms   70.45%
    Req/Sec   457.54     48.30   590.00     74.49%
  218967 requests in 30.10s, 26.73MB read
Requests/sec:   7274.18
Transfer/sec:      0.89MB

➜  wrk git:(master) ./wrk -t16 -c100 -d30s http://localhost:8000/
Running 30s test @ http://localhost:8000/
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    12.78ms    1.87ms  23.38ms   64.39%
    Req/Sec   469.84     49.90   595.00     63.57%
  225046 requests in 30.10s, 27.47MB read
Requests/sec:   7476.12
Transfer/sec:      0.91MB

############## 1000ms
➜  wrk git:(master) ./wrk -t16 -c100 -d30s http://localhost:8080/
Running 30s test @ http://localhost:8080/
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.01s     2.19ms   1.01s    61.42%
    Req/Sec     5.27      0.44     6.00     73.06%
  2784 requests in 30.09s, 348.00KB read
Requests/sec:     92.53
Transfer/sec:     11.57KB

➜  wrk git:(master) ./wrk -t16 -c100 -d30s http://localhost:8000/
Running 30s test @ http://localhost:8000/
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.01s     2.64ms   1.01s    65.19%
    Req/Sec     5.96      4.96    49.00     96.47%
  2784 requests in 30.08s, 348.00KB read
Requests/sec:     92.55
Transfer/sec:     11.57KB

```
