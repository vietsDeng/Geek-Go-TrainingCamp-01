# README

## protoc
```sh
cd proto
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=template/proto/google/api:. ./myapp/user.proto

protoc --grpc-gateway_out=logtostderr=true:. ./myapp/user.proto
```

## 证书

https://blog.csdn.net/weixin_40280629/article/details/113563351

## 打包命令
```sh
go build -o ./cmd/myapp/main ./cmd/myapp/
```

## 项目结构
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
