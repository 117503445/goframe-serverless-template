package controller

import (
	"context"

	"goframe-serverless-template/apiv1"
)

var (
	Task = cTask{}
)

type cTask struct{}

func (h *cTask) TaskGetAll(ctx context.Context, req *apiv1.TaskGetAllReq) (res []*apiv1.TaskGetAllRes, err error) {
	res = []*apiv1.TaskGetAllRes{
		{
			Id:     1,
			Title:  "task1",
			IsDone: true,
		}, {
			Id:     2,
			Title:  "task2",
			IsDone: false,
		},
	}

	return
}

func (h *cTask) TaskCreate(ctx context.Context, req *apiv1.TaskCreateReq) (res []*apiv1.TaskCreateRes, err error) {

	return
}
