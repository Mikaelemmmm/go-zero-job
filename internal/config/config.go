/**
* @Description 配置文件
* @Author Mikael
* @Email 13247629622@163.com
* @Date 2021/1/18 12:05
* @Version 1.0
**/

package config

import (
	"github.com/tal-tech/go-queue/dq"
	"github.com/tal-tech/go-zero/core/service"

)

type Config struct {
	service.ServiceConf
	DqConf dq.DqConf
}
