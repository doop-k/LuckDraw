package models

import "fmt"

type User struct {
	Id int64
	Username string
	DrawPaper *Paper
}

var Users []*User

func init(){
	for i := 1; i < 101; i++ {
		username := fmt.Sprintf("第[%d]个用户",i)
		Users = append(Users,&User{
			Id: int64(i),
			Username: username,
		})
	}
}