package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TaskGetAllReq struct {
	g.Meta `path:"/task" tags:"Task" method:"get"`
}
type TaskGetAllRes struct {
	g.Meta `mime:"application/json"`
	Id int
	Title string
	IsDone bool
}

type TaskCreateReq struct {
	g.Meta `path:"/task" tags:"Task" method:"post"`
}
type TaskCreateRes struct {
	g.Meta `mime:"application/json"`
}