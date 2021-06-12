# redis benchmark 测试

1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

```shell
redis-benchmark -h 127.0.0.1 -q -t set,get -r 10000 -n 100000 -d 10
redis-benchmark -h 127.0.0.1 -q -t set,get -r 10000 -n 100000 -d 20
redis-benchmark -h 127.0.0.1 -q -t set,get -r 10000 -n 100000 -d 50
redis-benchmark -h 127.0.0.1 -q -t set,get -r 10000 -n 100000 -d 100
redis-benchmark -h 127.0.0.1 -q -t set,get -r 10000 -n 100000 -d 200
redis-benchmark -h 127.0.0.1 -q -t set,get -r 10000 -n 100000 -d 1024
redis-benchmark -h 127.0.0.1 -q -t set,get -r 10000 -n 100000 -d 5120
```

| value字节大小/B | get 请求/秒 | set 请求/秒 |
| --- | --- | --- |
| 10 | 97943.19 | 102986.61 |
| 20 | 99502.48 | 94696.97 |
| 50 | 100200.40 | 97087.38 |
| 100 | 93109.87 | 88967.98 |
| 200 | 91324.20 | 88261.25 |
| 1k | 82440.23 | 78554.59 |
| 5k | 84317.03 | 77279.75 |

- 结论

随着value增大，get/set吞吐降低，但在测试集中，降低不明显。

2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
> info memory<br>
> 每次写入10W个左右key<br>
> key的格式，key:000000043017

```shell
redis-benchmark -h 127.0.0.1 -q -t set -r 100000 -n 200000 -d 10

flushdb
info memory
key *
```

| value字节大小/B | 写入前used_memory_dataset/B | 写入后used_memory_dataset/B | key数量 | 平均每个key的占用内存空间/B |
| --- | --- | --- | --- | --- |
| 10 | 6626 | 4851914 | 86519 | 56.00 |
| 20 | 7042 | 6233610 | 86477 | 72.00 |
| 50 | 7458 | 9007322 | 86535 | 104.00 |
| 100 | 8066 | 13168906 | 86583 | 152.00 |
| 200 | 8482 | 21448306 | 86450 | 248.00 |
| 1k | 8898 | 136271658 | 86461 | 1576.00 |
| 5k | 9314 | 489792410 | 86351 | 5672.00 |