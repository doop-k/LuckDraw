package models

import (
	"context"
	"fmt"
)

const (
	WaitTake = iota
	Taken
)

type Paper struct{
	Id int64
	Content string
	Status int64
	TakenUser int64
}
var Papers map[int64]*Paper


func init(){
	Papers = make(map[int64]*Paper,100)
	for i := 1; i < 101; i++ {
		id := int64(i)
		content := fmt.Sprintf("这是抽取的第:[%d]幸运纸条",id)
		Papers[id]=&Paper{
			Id:id,
			Content: content,
		}
	}
}

// 填充数据到redis
func FillUserOnRedis(){
	for _,paper := range Papers{
		var cmd = Client.LPush(context.Background(), "papers", paper.Id)
		if err:=cmd.Err();err!=nil{
			fmt.Println("err:",err)
			return
		}
	}
}