# Kratos Project Template

## 极客学院 - Go训练营一期作业

### 第二章作业
1.我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答: 遇到一个 sql error 的时候，应该利用errors.Is方法判断是否为sql.ErrNoRows。若是，则正常返回数据为空，因为该error抛出原因为查询行数为空，实际为正常业务逻辑，如：gorm已优化该处理；否则，包装error并返回给上层。

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


## Install Kratos
```
go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# create a template project
kratos new helloworld

cd helloworld
# Add a proto template
kratos proto add api/helloworld/helloworld.proto
# Generate the source code of service by proto file
kratos proto server api/helloworld/helloworld.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/helloworld -conf ./configs
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```
