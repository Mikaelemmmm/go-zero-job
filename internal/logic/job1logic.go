package logic

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"job/internal/svc"
)

/**
* @Description TODO
* @Version 1.0
**/

type Job1 struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJob1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Job1 {
	return &Job1{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Job1)Start()  {
	fmt.Printf("start job1 \n")
}

func (l *Job1)Stop()  {
	fmt.Printf("stop job1 \n")
}

