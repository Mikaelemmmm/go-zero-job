package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"job/internal/svc"
	"fmt"
)

/**
* @Description 消费者
* @Version 1.0
**/
type Consumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConsumerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Consumer {
	return &Consumer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Consumer)Start()  {
	fmt.Printf("start consumer \n")

	l.svcCtx.Consumer.Consume(func(body []byte) {
		fmt.Printf("consumer job  %s \n" ,string(body))
	})

}

func (l *Consumer)Stop()  {
	fmt.Printf("stop consumer \n")
}