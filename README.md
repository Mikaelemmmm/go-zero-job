### go-zero 分布式定时任务



日常任务开发中，我们会有很多异步、批量、定时、延迟任务要处理，go-zero中有go-queue，推荐使用go-queue去处理，go-queue本身也是基于go-zero开发的，其本身是有两种模式

- dq : 依赖于beanstalkd，分布式，可存储，延迟、定时设置，关机重启可以重新执行，消息会丢失，使用非常简单，go-queue中使用了redis setnx保证了每个消息只被消费一次，使用场景主要是用来做日常任务使用
- kq：依赖于kafka，这个就不多介绍啦，大名鼎鼎的kafka，使用场景主要是做日志用

我们主要说一下dq，kq使用也一样的，只是依赖底层不同，如果没使用过beanstalkd，没接触过beanstalkd的可以先google一下，使用起来还是挺容易的。

etc/job.yaml

```yaml
Name: job

Log:
  ServiceName: job
  Level: info

#dq 配置
DqConf:
  Beanstalks:
    - Endpoint: 127.0.0.1:7771
      Tube: tube1
    - Endpoint: 127.0.0.1:7772
      Tube: tube2
  Redis:
    Host: 127.0.0.1:6379
    Type: node
```



Internal/config/config.go

```go
type Config struct {
	service.ServiceConf
	DqConf dq.DqConf
}
```



Handler/router.go : 负责注册多任务

```go
func RegisterJob(serverCtx *svc.ServiceContext,group *service.ServiceGroup)  {

	group.Add(logic.NewProducerLogic(context.Background(),serverCtx))
	group.Add(logic.NewConsumerLogic(context.Background(),serverCtx))

	group.Start()
}
```



Logic:

```go
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


```



```doc
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
```



svc/servicecontext.go

```go
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

```



main.go启动文件

```go
package main

import (
	"flag"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/core/threading"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"job/internal/config"
	"job/internal/handler"
	"job/internal/svc"
)

/**
* @Description TODO
* @Version 1.0
**/
var configFile = flag.String("f", "etc/job.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	group := service.NewServiceGroup()
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	threading.GoSafe(func() {
		for s := range ch {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				group.Stop()
			}
		}
	})

	handler.RegisterJob(ctx,group)

	//阻塞直至有信号传入
	s := <-ch
	fmt.Println("退出job..", s)
}
```

















