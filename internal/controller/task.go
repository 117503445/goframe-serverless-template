package controller

import (
	"context"

	"goframe-serverless-template/apiv1"
	"goframe-serverless-template/internal/model/entity"
	"goframe-serverless-template/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
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
	items, err := service.Task().GetAll(ctx)
	if err := gconv.Structs(items, &res); err != nil {
		panic(err)
	}
	return
}

func (h *cTask) TaskCreate(ctx context.Context, req *apiv1.TaskCreateReq) (res []*apiv1.TaskCreateRes, err error) {
	service.Task().CreateOne(ctx, &entity.Task{
		Title:  req.Title,
		IsDone: 0,
	})
	return
}
