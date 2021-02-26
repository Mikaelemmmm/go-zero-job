package handler

import (
	"context"
	"github.com/tal-tech/go-zero/core/service"
	"zeroblog/app/services/job/internal/logic"
	"zeroblog/app/services/job/internal/svc"
)

/**
* @Description 注册job
* @Version 1.0
**/

func RegisterJob(serverCtx *svc.ServiceContext,group *service.ServiceGroup)  {

	group.Add(logic.NewJob1Logic(context.Background(),serverCtx))
	group.Add(logic.NewJob2Logic(context.Background(),serverCtx))

	group.Start()
}