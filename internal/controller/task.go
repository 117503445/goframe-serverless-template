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
	items, err := service.Task().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if err := gconv.Structs(items, &res); err != nil {
		panic(err)
	}
	if res == nil {
		res = make([]*apiv1.TaskGetAllRes, 0) // TODO process in middleware
	}
	return
}

func (h *cTask) TaskCreate(ctx context.Context, req *apiv1.TaskCreateReq) (res []*apiv1.TaskCreateRes, err error) {
	err = service.Task().CreateOne(ctx, &entity.Task{
		Title:  req.Title,
		IsDone: 0,
	})
	if err != nil {
		return nil, err
	}
	return
}
