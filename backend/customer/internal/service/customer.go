package service

import (
	"context"
	"customer/api/verifyCode"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"regexp"
	"time"

	pb "customer/api/customer"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeReq) (*pb.GetVerifyCodeResp, error) {
	// 写一个正则校验手机号
	pattern := `^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`
	// 编译正则表达式
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(req.Telephone) {
		return &pb.GetVerifyCodeResp{
			Code:    1,
			Message: "手机号格式错误",
		}, nil
	}
	// 创建一个grpc的链接
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		//grpc.WithMiddleware(
		//	recovery.Recovery(),
		//),
	)
	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code:    1,
			Message: "验证码服务不可用",
		}, nil
	}
	// 函数执行完的回调函数，关闭连接
	defer func() {
		_ = conn.Close()
	}()

	// 调用验证码服务
	client := verifyCode.NewVerifyCodeClient(conn)
	reply, err := client.GetVerifyCode(
		context.Background(),
		&verifyCode.GetVerifyCodeRequest{
			Length: 10,
			Type:   3,
		},
	)
	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code:    1,
			Message: "验证码获取失败",
		}, nil
	}

	// 返回当前的响应
	return &pb.GetVerifyCodeResp{
		Code:           0,
		Message:        "返回成功",
		VerifyCode:     reply.GetCode(),
		VerifyCodeTime: time.Now().Unix(),
		VerifyCodeLife: 60,
	}, nil
}
