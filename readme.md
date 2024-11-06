## 技术选型

| 功能/技术栈  | 详情               |
| ------------ | ------------------ |
| HTTP 框架    | hertz              |
| RPC 框架     | kitex              |
| 数据库       | MySQL, redis       |
| 代码生成     | cwgo(hertz, kitex) |
| 服务发现中心 | Consul             |
| 数据库映射   | gorm               |
| 限流中间件   | hertz-limiter      |
| 部署工具     | docker-compose     |
| 对象存储     | Tencent cloud      |
| 日志监测     | OpenTelemetry      |

## 需求设计

| 需求             | 详情                                                         |
| ---------------- | ------------------------------------------------------------ |
| 认证中心/auth    | 分发身份令牌, 续期身份令牌（高级）, 校验身份令牌             |
| 用户服务/user    | 创建用户, 登录, 用户登出（可选）, 删除用户（可选）, 更新用户（可选）, 获取用户身份信息 |
| 商品服务/product | 创建商品（可选）, 修改商品信息（可选）, 删除商品（可选）, 查询商品信息（单个商品、批量商品） |
| 购物车服务/cart  | 创建购物车, 清空购物车, 获取购物车信息                       |
| 订单服务/order   | 创建订单, 修改订单信息（可选）, 订单定时取消（高级）         |
| 结算/checkout    | 订单结算                                                     |
| 支付/payment     | 取消支付（高级）, 定时取消支付（高级）, 支付                 |

## 设计思路

hertz作为http服务器担任网关功能，kitex作为rpc框架实现微服务

## 代码建构

[代码建构](instruction.md)

## 参考文档

1. [CloudWeGo](https://cloudwego.cn/zh/) 
2. [Gomall](https://github.com/cloudwego/biz-demo/tree/main/gomall)
