package service

import (
	"context"
	"goframe-serverless-template/internal/model/entity"
	"goframe-serverless-template/internal/service/internal/dao"
)

type (
	// sTask is service struct of module Task.
	sTask struct{}
)

var (
	// insTask is the instance of service Task.
	insTask = sTask{}
)

// Task returns the interface of Task service.
func Task() *sTask {
	return &insTask
}

func (s *sTask) GetAll(ctx context.Context) (items []*entity.Task, err error) {
	// items := ([]*entity.Task)(nil)
	dao.Task.Ctx(ctx).Scan(&items)
	return items, nil
}

func (s *sTask) CreateOne(ctx context.Context,item *entity.Task) (err error) {
	// items := ([]*entity.Task)(nil)
	_, err = dao.Task.Ctx(ctx).Data(item).Insert()
	return err
}
