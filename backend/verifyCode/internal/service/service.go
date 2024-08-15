package service

import "github.com/google/wire"

// ProviderSet is service providers.
// 将服务代码注册到grpc服务中
var ProviderSet = wire.NewSet(NewGreeterService, NewVerifyCodeService)
