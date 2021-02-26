package config

import (
	"github.com/tal-tech/go-queue/dq"
	"github.com/tal-tech/go-zero/core/service"

)

/**
* @Description TODO
* @Version 1.0
**/

type Config struct {
	service.ServiceConf
	DqConf dq.DqConf
}
