package handler

import (
	"context"
	"github.com/tal-tech/go-zero/core/service"
	"job/internal/logic"
	"job/internal/svc"
)

/**
* @Description 注册job
* @Version 1.0
**/

func RegisterJob(serverCtx *svc.ServiceContext,group *service.ServiceGroup)  {

	group.Add(logic.NewProducerLogic(context.Background(),serverCtx))
	group.Add(logic.NewConsumerLogic(context.Background(),serverCtx))

	group.Start()
}