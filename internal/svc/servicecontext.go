/**
* @Description 配置
* @Author Mikael
* @Email 13247629622@163.com
* @Date 2021/1/18 12:05
* @Version 1.0
**/
package svc

import (
	"job/internal/config"
	"github.com/tal-tech/go-queue/dq"
)

type ServiceContext struct {
	Config config.Config
	Consumer      dq.Consumer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Consumer: dq.NewConsumer(c.DqConf),
	}
}
