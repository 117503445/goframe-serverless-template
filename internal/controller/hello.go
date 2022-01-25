package controller

import (
	"context"

	"goframe-serverless-template/apiv1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (h *cHello) Hello(ctx context.Context, req *apiv1.HelloReq) (res *apiv1.HelloRes, err error) {
	res = &apiv1.HelloRes{
		Version: "1.0",
	}
	return
}
