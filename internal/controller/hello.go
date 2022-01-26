package controller

import (
	"context"

	"goframe-serverless-template/apiv1"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/genv"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (h *cHello) Hello(ctx context.Context, req *apiv1.HelloReq) (res *apiv1.HelloRes, err error) {
	g.Log().Line(true).Debug(ctx, "Hello")
	dblink := genv.GetWithCmd("dblink").String()
	g.Log().Line().Debug(ctx, "dblink = ", dblink)
	res = &apiv1.HelloRes{
		Version: "1.1",
	}
	return
}
