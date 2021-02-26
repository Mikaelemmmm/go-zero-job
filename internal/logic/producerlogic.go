package logic

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-queue/dq"
	"github.com/tal-tech/go-zero/core/logx"
	"job/internal/svc"
	"strconv"
	"time"
)

/**
* @Description 生产者
* @Version 1.0
**/

type Producer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProducerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Producer {
	return &Producer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Producer)Start()  {
	fmt.Printf("start  Producer \n")

	producer := dq.NewProducer([]dq.Beanstalk{
		{
			Endpoint: "localhost:7771",
			Tube:     "tube1",
		},
		{
			Endpoint: "localhost:7772",
			Tube:     "tube2",
		},
	})
	for i := 1000; i < 1005; i++ {
		_, err := producer.Delay([]byte(strconv.Itoa(i)), time.Second * 1)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (l *Producer)Stop()  {
	fmt.Printf("stop Producer \n")
}

