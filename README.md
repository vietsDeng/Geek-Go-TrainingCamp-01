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
- Template
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
