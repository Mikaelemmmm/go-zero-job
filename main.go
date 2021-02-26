/**
* @Description 启动文件
* @Author Mikael
* @Email 13247629622@163.com
* @Date 2021/1/18 12:05
* @Version 1.0
**/
package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"job/internal/config"
	"job/internal/handler"
	"job/internal/svc"
	"os"
	"os/signal"
	"syscall"
	"time"
)


var configFile = flag.String("f", "etc/job.yaml", "the config file")

func main() {
	flag.Parse()

	//配置
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	fmt.Printf("ctx : %v",ctx)

	//注册job
	group := service.NewServiceGroup()
	handler.RegisterJob(ctx,group)

	//捕捉信号
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-ch
		logx.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Printf("stop group")
			//group.Stop()
			logx.Info("job exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}