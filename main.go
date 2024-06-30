package main

import (
	"fmt"
	"gin-web-app/dao/mysql"
	"gin-web-app/dao/redis"
	"gin-web-app/logger"
	"gin-web-app/routes"
	"gin-web-app/settings"

	"go.uber.org/zap"
)

func main() {
	// 1. 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync() // 缓冲区的日志追加到日志中
	// 3. 初始化 mysql
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	// 4. 初始化 redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	// 5. 注册路由
	r := routes.Setup()
	// 6. 启动服务
	r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
}
