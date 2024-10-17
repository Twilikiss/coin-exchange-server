// Package logic
// @Author twilikiss 2024/4/29 17:04:04
package logic

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"log"
	"testing"
	"time"
	"ucenter/internal/config"
	"ucenter/internal/svc"
)

var configFile = flag.String("f", "../../etc/conf.yaml", "the config file")

func TestRegisterLogic_SendCode(t *testing.T) {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx1 := svc.NewServiceContext(c)
	country := "中国"
	phone := "19820754340"
	key := RegisterCacheKey + country + "::" + phone
	ctx2, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	err := ctx1.Cache.SetexCtx(ctx2, key, "114514", 5*60)
	if err != nil {
		log.Fatal(err)
	}
}

func TestRegisterLogic_CheckCode(t *testing.T) {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx1 := svc.NewServiceContext(c)
	country := "中国"
	phone := "19820754340"
	key := RegisterCacheKey + country + "::" + phone

	var temp string
	temp, _ = ctx1.Cache.Get(key)
	if temp == "" {
		log.Println("当前key不存在")
	} else {
		ttl, _ := ctx1.Cache.Ttl(key)
		log.Printf("找到key，value=%s，过期时间=%d\n", temp, ttl)
	}
}
