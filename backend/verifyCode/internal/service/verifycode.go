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
		Code: RandCode(int(req.Length), req.Type),
	}, nil
}
func RandCode(l int, t pb.TYPE) string {
	switch t {
	case pb.TYPE_DEFAULT:
		fallthrough
	case pb.TYPE_DIGIT:
		chars := "1234567890"
		return randCode(chars, l)
	case pb.TYPE_LETTER:
		chars := "abcdefghijklmnopqrstuvwsyz"
		return randCode(chars, l)
	case pb.TYPE_MIXED:
		chars := "1234567890abcdefghijklmnopqrstuvwsyz"
		return randCode(chars, l)
	default:

	}
	return ""
}
func randCode(chars string, l int) string {
	charsLen := len(chars)
	result := make([]byte, l)
	for i := 0; i < l; i++ {
		randIndex := rand.Intn(charsLen)
		result[i] = chars[randIndex]
	}
    fmt.Println(string(result))
	return string(result)
}
