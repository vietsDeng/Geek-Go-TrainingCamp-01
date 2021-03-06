## 极客学院 - Go训练营一期作业

### 第二章作业

1.我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答: 遇到一个 sql error 的时候，应该利用errors.Is方法判断是否为sql.ErrNoRows。若是，则正常返回数据为空，因为该error抛出原因为查询行数为空，实际为正常业务逻辑，如：gorm已优化该处理；否则，包装error并返回给上层。

路径：
```
helloWorld
```

具体代码:
```
internal/dao/dao.go 的QueryUserList方法

# 关键代码
rows, err := db.Query(selectSql, level)
if err != nil {
    //sql.ErrNoRows，则返回没有数据
    if errors.Is(err, sql.ErrNoRows) {
        return &list, 0, nil
    } else {
        //其他错误，则包装返回
        return nil, 0, fmt.Errorf("Query User List Failed: %w", err)
    }
}
```

### 第三章作业

1.基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

答: 

启动三个协程：

g1 为http server，通过请求触发 g2 引导自身退出

g2 作用为引导 g1 退出

g3 作用为接收 linux signal 信号退出

路径：
```
ErrGroupDemo
```

具体代码:
```
// 退出路径
// 1、[g1] http请求 -> g1发送信号 -> [g2]继续 -> [g1]退出 -> g3[退出] -> [g2]退出 -> wait退出
// 2、[g3] os中断 -> [g3]退出 -> [g2]继续 -> [g1]退出  -> [g2]退出 -> wait退出
```

## 第四章作业

1.按照自己的构想，写一个项目满足基本的目录结构和工程，代码需包含对数据层、业务层、API注册，以及main函数对于服务的注册和启动，信号处理，使用Wire构建依赖。可以使用自己熟悉的框架。

答: 

```sh
- template
    - certs [证书过程及最终文件，满足x509]
    - cmd
        - myapp [单独一个应用服务进行打包]
            - main.go [<服务注册> 启动文件]
            - user.go [<API注册> 接口绑定]
    - config [配置文件]
        - db.toml [数据库配置]
    - internal
        - pkg [内部工具包]
        - myapp
            - biz [api -> biz，外部接口调用]
            - data [<数据层> service -> data，数据库操作]
            - model [数据结构]
                - biz [biz数据结构]
                - data [数据库数据结构]
            - service [<业务层> biz -> service，服务接口，供biz组装]
            - server.go [服务常用变量，含初始化]
            - wire.go [wire初始化依赖]
    - pkg [外部工具包]
        - util
    - proto [api定义]
        - google [grpc-gateway依赖文件]
        - myapp [grpc proto文件，支持grpc & http]
    - test [单元测试]
        - myapp
```

## 第八章作业

1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

## 第九章作业

1、总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用。

2、实现一个从 socket connection 中解码出 goim 协议的解码器。

## 毕业项目

> Auth

对当下自己项目中的业务，进行一个微服务改造，需要考虑如下技术点：

1）微服务架构（BFF、Service、Admin、Job、Task 分模块）

2）API 设计（包括 API 定义、错误码规范、Error 的使用）

3）gRPC 的使用

4）Go 项目工程化（项目结构、DI、代码分层、ORM 框架）

5）并发的使用（errgroup 的并行链路请求

6）微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）

7）缓存的使用优化（一致性处理、Pipeline 优化）