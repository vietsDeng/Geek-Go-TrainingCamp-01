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
