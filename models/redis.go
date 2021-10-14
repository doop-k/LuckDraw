package models

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Client  *redis.Client
const redis_url = "redis://localhost:6379/1"
func init(){
	opt,err := redis.ParseURL(redis_url)
	if err!=nil{
		fmt.Println("连接redis失败,地址:",redis_url)
		return
	}
	Client = redis.NewClient(opt)
}