package controller

import (
	"context"

	"goframe-serverless-template/app/svc-template/apiv1"
)

var (
	Echo = cEcho{}
)

type cEcho struct{}

func (h *cEcho) Say(ctx context.Context, req *apiv1.EchoSayReq) (res *apiv1.EchoSayRes, err error) {
	return &apiv1.EchoSayRes{Content: req.Content}, nil
}
