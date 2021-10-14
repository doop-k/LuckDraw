package models

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
)

var wg *sync.WaitGroup
var lock sync.Mutex

func RunLuckDraw(w *sync.WaitGroup,m sync.Mutex){
	wg = w
	lock = m
	fmt.Println("抽奖前剩余:", len(Papers))
	for _,user := range Users{
		wg.Add(1)
		go luckDraw(user) //抽取一次奖
	}
	wg.Wait()
	last := 0
	for _,v := range Papers{
		if v.Status == WaitTake{
			last+=1
		}
	}
	fmt.Println("抽奖后剩余:",last)
	users := make([]*User,0,100)
	for _,user := range Users{
		if user.DrawPaper != nil{
			users = append(users,user)
		}
	}
	fmt.Println("-------------用户抽取情况-------------")
	for _,user := range users{
		fmt.Printf("user=>%+v --- paper=>%+v\n",user,user.DrawPaper)
	}
	fmt.Println("------------------------------------")
}


func luckDraw(user *User){
	lock.Lock()
	defer wg.Done()
	defer lock.Unlock()
	fmt.Printf("用户[%s]开始抽奖\n",user.Username)
	cmd := Client.LLen(context.Background(),"papers")
	if cmd.Err() !=nil {
		fmt.Println("err:",cmd.Err())
		return
	}
	paper_len := int(cmd.Val())
	if paper_len <= 0 {
		fmt.Printf("err:用户[%s]来晚了,已经没有纸条了.\n",user.Username)
		return
	}

	index := rand.Intn(paper_len)
	valueCmd := Client.LIndex(context.Background(),"papers",int64(index))
	if valueCmd.Err() !=nil{
		fmt.Println("err:",valueCmd.Err())
		return
	}
	paper_id,err:=valueCmd.Int64()
	if err!=nil{
		fmt.Println("err:",err)
		return
	}
	paper,has:=Papers[paper_id]
	if !has{
		fmt.Println("err:",err)
		return
	}
	if paper.Status == Taken{
		fmt.Println("err:",err)
		return
	}
	Papers[paper_id].Status = Taken
	user.DrawPaper = paper
	remCmd := Client.LRem(context.Background(),"papers",1,paper_id)
	if remCmd.Err() !=nil {
		fmt.Println("err:",cmd.Err())
		return
	}
	fmt.Printf("用户【%s】抽奖完毕,抽奖结果:【%s】\n",user.Username,paper.Content)
}
