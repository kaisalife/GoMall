package dal

import (
	"github.com/cloudwego/gomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/cloudwego/gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
