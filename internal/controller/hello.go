package controller

import (
	"context"

	"goframe-serverless-template/apiv1"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (h *cHello) Hello(ctx context.Context, req *apiv1.HelloReq) (res *apiv1.HelloRes, err error) {
	g.Log().Line(true).Debug(ctx, "Hello")
	res = &apiv1.HelloRes{
		Version: "1.4",
	}
	return
}
