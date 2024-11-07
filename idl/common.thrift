namespace go common

// 错误类的定义
enum Err {
    // 通用错误
    UnknownError            = 10000, // 未知错误
    BadRequest              = 10001, // 无效的请求参数
    Unauthorized            = 10002, // 未授权访问
    Forbidden               = 10003, // 无权限访问
    NotFound                = 10004, // 未找到资源
    InternalServerError     = 10005, // 内部服务器错误
    ServiceUnavailable      = 10006, // 服务不可用
	ResponseUnableParse     = 10007, // 无法解析响应
	
    // 网关错误
    GatewayBadRequest       = 20001, // 网关收到无效请求
    GatewayTimeout          = 20002, // 网关请求超时
    GatewayServiceUnavailable = 20003, // 网关服务不可用

    // 用户服务错误
    UserNotFound            = 30001, // 用户未找到
    InvalidUserCredentials  = 30002, // 无效的用户凭据
    UserAlreadyExists       = 30003, // 用户已存在
    UserForbidden           = 30004, // 用户无权限

    // 产品服务错误
    ProductNotFound         = 40001, // 产品未找到
    ProductOutOfStock       = 40002, // 产品库存不足
    ProductInvalidRequest   = 40003, // 产品请求无效

    // 购物车服务错误
    CartItemNotFound        = 50001, // 购物车项未找到
    CartInvalidOperation    = 50002, // 购物车操作无效
    CartUpdateFailed        = 50003, // 更新购物车失败

    // 订单服务错误
    OrderNotFound           = 60001, // 订单未找到
    OrderInvalidStatus      = 60002, // 订单状态无效
    OrderCreationFailed     = 60003, // 创建订单失败
    OrderCancellationFailed = 60004, // 取消订单失败

    // 支付服务错误
    PaymentFailed           = 70001, // 支付失败
    PaymentNotFound         = 70002, // 支付记录未找到
    PaymentAuthorizationFailed = 70003, // 支付授权失败

    // 库存服务错误
    InventoryNotAvailable   = 80001, // 库存不足
    InventoryUpdateFailed   = 80002, // 更新库存失败
    InventoryNotFound       = 80003  // 库存未找到
}