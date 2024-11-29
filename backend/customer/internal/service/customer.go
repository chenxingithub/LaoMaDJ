package service

import (
	"context"
	"regexp"
	"time"

	pb "customer/api/customer"
	"customer/api/verifyCode"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeReq) (*pb.GetVerifyCodeResp, error) {

	// 校验手机号
	pattern := `^1[3-9]\d{9}$`
	regexpPattern := regexp.MustCompile(pattern)
	if !regexpPattern.MatchString(req.Telephone){
		return &pb.GetVerifyCodeResp{
			Code: 1,
			Message: "手机号码格式错误",
		}, nil
	}

	// 连接验证码生成服务(grpc)
    conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("localhost:9000"),
	)

	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code: 1,
			Message: "验证码服务不可用",
		}, nil
	}

	// 关闭连接
	defer func ()  {
		_ = conn.Close()
	}()

	//请求获取验证码(grpc)
	client := verifyCode.NewVerifyCodeClient(conn)
	reply, err := client.GetVerifyCode(context.Background(), &verifyCode.GetVerifyCodeRequest{
		Length: 6,
		Type: 1,
	})

	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code: 1,
			Message: "验证码获取失败",
		}, nil
	}

	return &pb.GetVerifyCodeResp{
		Code:  0,
		Message: "",
		VerifyCode: reply.Code,
		VerifyCodeTime: time.Now().Unix(),
	}, nil
}
