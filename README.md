# Apache benchmark

```
ab -n 5000 -c 100 http://localhost:3001/
```

# Running servers

## Ruby

```
bundle exec thin start -R app.ru -p 3001
```

## Go

```
GOMAXPROCS=4 go run app.go
```

## Node.js

```
node app.js
```


# Results

## Go

```
✗ ab -n 5000 -c 100 http://localhost:3000/
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 500 requests
Completed 1000 requests
Completed 1500 requests
Completed 2000 requests
Completed 2500 requests
Completed 3000 requests
Completed 3500 requests
Completed 4000 requests
Completed 4500 requests
Completed 5000 requests
Finished 5000 requests


Server Software:
Server Hostname:        localhost
Server Port:            3000

Document Path:          /
Document Length:        5832 bytes

Concurrency Level:      100
Time taken for tests:   2.957 seconds
Complete requests:      5000
Failed requests:        0
Write errors:           0
Total transferred:      29600000 bytes
HTML transferred:       29160000 bytes
Requests per second:    1690.95 [#/sec] (mean)
Time per request:       59.138 [ms] (mean)
Time per request:       0.591 [ms] (mean, across all concurrent requests)
Transfer rate:          9775.80 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.6      0       8
Processing:     2   58  44.0     54     437
Waiting:        2   58  44.0     54     437
Total:          2   59  44.0     54     437

Percentage of the requests served within a certain time (ms)
  50%     54
  66%     61
  75%     69
  80%     88
  90%    111
  95%    146
  98%    182
  99%    208
 100%    437 (longest request)
```

## Ruby

```
✗ ab -n 5000 -c 100 http://localhost:3001/
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 500 requests
Completed 1000 requests
Completed 1500 requests
Completed 2000 requests
Completed 2500 requests
Completed 3000 requests
Completed 3500 requests
Completed 4000 requests
Completed 4500 requests
Completed 5000 requests
Finished 5000 requests


Server Software:        thin
Server Hostname:        localhost
Server Port:            3001

Document Path:          /
Document Length:        6002 bytes

Concurrency Level:      100
Time taken for tests:   7.890 seconds
Complete requests:      5000
Failed requests:        0
Write errors:           0
Total transferred:      30675000 bytes
HTML transferred:       30010000 bytes
Requests per second:    633.72 [#/sec] (mean)
Time per request:       157.799 [ms] (mean)
Time per request:       1.578 [ms] (mean, across all concurrent requests)
Transfer rate:          3796.73 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.6      0       8
Processing:    60  156  27.5    161     236
Waiting:       12  142  26.5    145     217
Total:         68  157  27.3    162     236

Percentage of the requests served within a certain time (ms)
  50%    162
  66%    171
  75%    176
  80%    181
  90%    188
  95%    199
  98%    204
  99%    211
 100%    236 (longest request)

```


## Node.js

```
✗ ab -n 5000 -c 100 http://localhost:3002/
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 500 requests
Completed 1000 requests
Completed 1500 requests
Completed 2000 requests
Completed 2500 requests
Completed 3000 requests
Completed 3500 requests
Completed 4000 requests
Completed 4500 requests
Completed 5000 requests
Finished 5000 requests


Server Software:
Server Hostname:        localhost
Server Port:            3002

Document Path:          /
Document Length:        5802 bytes

Concurrency Level:      100
Time taken for tests:   3.081 seconds
Complete requests:      5000
Failed requests:        0
Write errors:           0
Total transferred:      29545000 bytes
HTML transferred:       29010000 bytes
Requests per second:    1622.82 [#/sec] (mean)
Time per request:       61.621 [ms] (mean)
Time per request:       0.616 [ms] (mean, across all concurrent requests)
Transfer rate:          9364.51 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       9
Processing:     2   61  33.6     55     255
Waiting:        2   61  33.6     55     254
Total:          2   61  33.8     55     256

Percentage of the requests served within a certain time (ms)
  50%     55
  66%     68
  75%     77
  80%     83
  90%    101
  95%    121
  98%    154
  99%    189
 100%    256 (longest request)
 ```
