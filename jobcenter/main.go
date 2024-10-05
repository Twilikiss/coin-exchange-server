// Package jobcenter
// @Author twilikiss 2024/5/4 22:50:50
package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"jobcenter/internal/config"
	"jobcenter/internal/svc"
	"jobcenter/internal/task"
	"os"
	"os/signal"
	"syscall"
)

var configFile = flag.String("f", "etc/conf.yaml", "the config file")

// TODO 改造项目为通过已验证身份的http请求的形式开启和关闭任务
func main() {
	flag.Parse()

	// 为了调试方便，我们先将log日志设置一下
	logx.MustSetup(logx.LogConf{
		Stat:     false,   // 关闭定时的stat检测
		Encoding: "plain", // 限两种：json和plain，默认json，plain表示直接打印到控制台上
	})

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 创建上下文
	ctx := svc.NewServiceContext(c)
	// 创建并启动任务
	t := task.NewTask(ctx)
	t.Run()

	//优雅退出
	go func() {
		exit := make(chan os.Signal)
		signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-exit:
			logx.Info("任务中心中断执行，开始clear资源")
			t.Stop()
			ctx.MongoClient.Disconnect()
		}
	}()

	t.StartBlocking()
}
