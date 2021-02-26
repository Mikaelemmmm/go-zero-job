package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"zeroblog/app/services/job/internal/svc"
	"fmt"
)

/**
* @Description TODO
* @Version 1.0
**/
type Job2 struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJob2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Job2 {
	return &Job2{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Job2)Start()  {
	fmt.Printf("start job2 \n")
}

func (l *Job2)Stop()  {
	fmt.Printf("stop job2 \n")
}