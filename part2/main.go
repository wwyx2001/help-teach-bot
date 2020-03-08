package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"helpteachbot/model"
	"helpteachbot/controller"
	"time"
)

func through(){
	ticker := time.Tick(time.Minute)//循环执行不需要清理
	sc := 0
	for {
		ti := time.Now()
		if ti.Second() !=0 {
			sc = 0
			continue
		}
		if sc==1 {continue}
		sc=1

		Namelist, err := model.FindName()
		if err != nil{
			fmt.Println(err)
			return
		}
		for _, cur :=range Namelist{
			controller.Calluser(cur.UserID)
		}
		<-ticker
	}
}

func main(){
	go through()
	r := gin.Default()
	r.POST("/",controller.ReceivePost)
	_ = r.Run()
}
