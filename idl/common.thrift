namespace go common

// Definition of error types
enum Err {
    // General errors
    UnknownError            = 10000, // Unknown error
    BadRequest              = 10001, // Invalid request parameters
    Unauthorized            = 10002, // Unauthorized access
    ServerNotFound          = 10003, // Server not found 
    ServerMethodNotFound    = 10004, // Server method not found 
    RequestServerFail       = 10005, // Request to server failed
    ServerHandleFail        = 10006, // Server handling request failed 
    ResponseUnableParse     = 10007, // Unable to parse server response
	
    // Gateway errors
    GatewayBadRequest       = 20001, // Gateway received invalid request
    GatewayTimeout          = 20002, // Gateway request timeout
    GatewayServiceUnavailable = 20003, // Gateway service unavailable

    // User service errors
    UserNotFound            = 30001, // User not found
    InvalidUserCredentials  = 30002, // Invalid user credentials
    UserAlreadyExists       = 30003, // User already exists
    UserForbidden           = 30004, // User forbidden

    // Product service errors
    ProductNotFound         = 40001, // Product not found
    ProductOutOfStock       = 40002, // Product out of stock
    ProductInvalidRequest   = 40003, // Invalid product request

    // Cart service errors
    CartItemNotFound        = 50001, // Cart item not found
    CartInvalidOperation    = 50002, // Invalid cart operation
    CartUpdateFailed        = 50003, // Failed to update cart

    // Order service errors
    OrderNotFound           = 60001, // Order not found
    OrderInvalidStatus      = 60002, // Invalid order status
    OrderCreationFailed     = 60003, // Failed to create order
    OrderCancellationFailed = 60004, // Failed to cancel order

    // Payment service errors
    PaymentFailed           = 70001, // Payment failed
    PaymentNotFound         = 70002, // Payment record not found
    PaymentAuthorizationFailed = 70003, // Payment authorization failed

    // Inventory service errors
    InventoryNotAvailable   = 80001, // Inventory not available
    InventoryUpdateFailed   = 80002, // Failed to update inventory
    InventoryNotFound       = 80003  // Inventory not found
}