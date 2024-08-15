package service

import (
	"context"
	"fmt"
	"math/rand"

	pb "verifyCode/api/verifyCode"
)

type VerifyCodeService struct {
	pb.UnimplementedVerifyCodeServer
}

func NewVerifyCodeService() *VerifyCodeService {
	return &VerifyCodeService{}
}

func (s *VerifyCodeService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	return &pb.GetVerifyCodeReply{
		Code: RandCode(req.Length, req.Type),
	}, nil
}

// RandCode 开放调用的随机生成验证码
// TODO 实现验证码的生成
// length - 验证码长度
// type - 验证码类型：数字(digit), 字母(letter), 混合(mixed)
func RandCode(length int32, p pb.TYPE) (code string) {
	switch p {
	case pb.TYPE_DEFAULT:
		fallthrough
	case pb.TYPE_DIGIT:
		return randCode("0123456789", length, 4)
	case pb.TYPE_LETTER:
		return randCode("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", length, 5)
	case pb.TYPE_MIXED:
		return randCode("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", length, 6)
	}
	return ""
}

// randCode 生成指定长度的随机数字字符串
// chars - 包含生成随机数的字符范围
// lenght - 要生成的随机数的长度
// TODO 实现一个简单, 直接的随机数生成器
//func randCode(chars string, lenght int32) string {
//	result := make([]byte, lenght)
//	for i := 0; i < int(lenght); i++ {
//		result[i] = chars[rand.Intn(len(chars))]
//	}
//	return string(result)
//}

// randCode 生成指定长度的随机数
// chars - 包含生成随机数的字符范围
// lenght - 要生成的随机数的长度
// indexBit - 每个随机数的位数
// TODO 实现一个高效的随机数生成器
func randCode(chars string, lenght, indexBit int32) string {
	// 实现一个掩码
	idxMask := ^(-1 << indexBit)
	// 63可以用多少次
	idxMax := 63 / indexBit
	// 返回字符串
	result := make([]byte, lenght)
	// 实现一个随机数生成器
	for i, cache, remain := 0, rand.Int63(), idxMax; i < int(lenght); {
		if 0 == remain {
			fmt.Println("test")
			cache, remain = rand.Int63(), idxMax
		}
		if randIdx := int(cache & int64(idxMask)); randIdx < len(chars) {
			result[i] = chars[randIdx]
			i++
		}
		//	使用下一组
		cache >>= indexBit
		remain--
	}
	return string(result)
}
