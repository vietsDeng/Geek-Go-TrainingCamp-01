# 毕业项目（旧项目优化）

## 背景

工作上主要还是以PHP为主，Go项目较少。

暂且找出一个工作上被弃置的小项目做学习使用。

## 日志

2021.07 个人事宜较忙，先占个坑

## 要点
> 对当下自己项目中的业务，进行一个微服务改造，需要考虑如下技术点：
>
> 1）微服务架构（BFF、Service、Admin、Job、Task 分模块）
>
> 2）API 设计（包括 API 定义、错误码规范、Error 的使用）
>
> 3）gRPC 的使用
>
> 4）Go 项目工程化（项目结构、DI、代码分层、ORM 框架）
>
> 5）并发的使用（errgroup 的并行链路请求
>
> 6）微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）
>
> 7）缓存的使用优化（一致性处理、Pipeline 优化）

## 现状

1. GIN + GORM/MONGO + CASBIN + WIRE 实现的RBAC权限管理脚手架
2. 使用了路由器router，未使用gRPC，先改造使用gRPC