## 代码建构

### hertz-build

初始化当前/TinyMarket下的的文件结构

.
├── cmd
│   └── hz-gateway
├── idl
│   ├── auth.proto
│   ├── cart.proto
│   ├── checkout.proto
│   ├── order.proto
│   ├── payment.proto
│   ├── product.proto
│   └── user.proto
├── instruction.md
└── readme.md

4 directories, 9 files

在当前/TinyMarket路径下执行，初始化go.mod

```
go mod init TinyMarket
```

使用hertz构建一个http服务器。在cmd/hz-gateway路径下执行

```
hz new -module TinyMarket --out_dir cmd/hz-gateway
```

执行命令后，当前/TinyMarket/cmd/hz-gateway下的的文件树

.
├── biz
│   ├── handler
│   │   └── ping.go
│   └── router
│       └── register.go
├── build.sh
├── main.go
├── router.go
├── router_gen.go
└── script
    └── bootstrap.sh

5 directories, 7 files

在当前/TinyMarket/cmd/hz-gateway路径下

```
cd biz
mkdir errors
mkdir middleware
```

下面新建文件。

在当前/TinyMarket/cmd/hz-gateway/errors路径下新建errors.go，这是处理网关错误的实现。

在当前/TinyMarket/cmd/hz-gateway/middleware路径下新建gateway_sign.go，这是分发给网关层的请求签名功能的实现。

在当前/TinyMarket/cmd/hz-gateway/types路径下新建types.go，这是响应信息的定义。

在当前/TinyMarket/idl路径下新建common.thrift，这是定义通用信息的idl文件。

在当前/TinyMarket/cmd/hz-gateway/handler路径下将ping.go删除，新建handler.go，网关功能的实现。

### hertz-code

在完成了文件结构的设计后，开始进入到hertz-gateway层的代码设计。

1.定义common.thrift中网关错误部分，见[common.thrift](/idl/common.thrift)。

2.在当前/TinyMarket路径下执行以下命令

```
kitex -module TinyMarket idl/common.thrift
```

3.完成/TinyMarket/cmd/hz-gateway/errors/errors.go，/TinyMarket/cmd/hz-gateway/types/types.go，/TinyMarket/cmd/hz-gateway/handler/handler.go，/TinyMarket/cmd/hz-gateway/route.go代码
