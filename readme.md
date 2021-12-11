# User System

## Introduction

This repo contains the final project for https://u.geekbang.org/subject/go. It implements a hypothetical user system
which considered the following simplified use cases:

- A user can create an account with a name and email.
- An admin can list all user accounts in detail, including name, email, status, registration date, etc.

## Tasks

对当下自己项目中的业务，进行一个微服务改造，需要考虑如下技术点：

-[x] 微服务架构（BFF、Service、Admin、Job、Task 分模块）
-[x] API 设计（包括 API 定义、错误码规范、Error 的使用）
-[x] gRPC 的使用
-[x] Go 项目工程化（项目结构、DI、代码分层、ORM 框架）
-[ ] 并发的使用（errgroup 的并行链路请求）
-[ ] 微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）
-[ ] 缓存的使用优化（一致性处理、Pipeline 优化）

## Developer Guides

### Prerequisites

- Golang 1.16 +
- A MySQL database call `user_system`

### Quick start

1. build go binaries

```
make mod
make build 
```

2. start the server: `./server`
3. make a request via the client

```
$ ./client
2021/12/11 14:24:04 user created!
```

### Making changes
To update the gRPC service, follow [gRPC and Protobuf](https://grpc.io/docs/languages/go/quickstart/) and run `make api` to regenerate gRPC stubs. 







