package response

// 业务响应码枚举
const (
	CodeSuccess       = 200  // 成功
	CodeParamError    = 400  // 参数错误
	CodeUnauthorized  = 401  // 未授权
	CodeForbidden     = 403  // 禁止访问
	CodeNotFound      = 404  // 资源不存在
	CodeServerError   = 500  // 服务器内部错误
)