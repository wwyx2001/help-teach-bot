package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"helpteachbot/model"
	"net/http"
	"strings"
	"time"
)

const (
	Help = "1.回复\"add,离散数学作业,2020-03-01 20:26:20,x,y\"即可添加一个名叫\"离散数学作业\"的任务，并且在提前x小时每隔y分钟提醒一次\n2.回复\"tasks\"即可列出自己的任务列表以及编号\n3.回复\"del,233\"即可删除编号为233的任务\n"
	Homework = "今天的作业是练习一"
)

type QQprivatemessage struct{
	UserID    int     `json:"user_id"`
	Message   string  `json:"message"`
}

type QQgroupmessage struct{
	GroupID   int     `json:"group_id"`
	Message   string  `json:"message"`
}

func Sendmessage(message string,number int){
	client := &http.Client{}
	msg := QQprivatemessage{
		UserID:  number,
		Message: message,
	}
	Requestbody := new(bytes.Buffer)
	err := json.NewEncoder(Requestbody).Encode(msg)
	if err !=nil {
		fmt.Println(err)
		return
	}
	Request, err :=http.NewRequest("POST" , "http://175.24.41.84:5700/send_msg" , Requestbody)
	if err !=nil {
		fmt.Println(err)
		return
	}
	Request.Header.Set("Content-Type","application/json")
	_, err = client.Do(Request)
	if err !=nil {
		fmt.Println(err)
		return
	}
}

func Sendgroupmessage(message string,number int){
	client := &http.Client{}
	msg := QQgroupmessage{
		GroupID:  number,
		Message: message,
	}
	Requestbody := new(bytes.Buffer)
	err := json.NewEncoder(Requestbody).Encode(msg)
	if err !=nil {
		fmt.Println(err)
		return
	}
	Request, err :=http.NewRequest("POST" , "http://175.24.41.84:5700/send_msg" , Requestbody)
	if err !=nil {
		fmt.Println(err)
		return
	}
	Request.Header.Set("Content-Type","application/json")
	_, err = client.Do(Request)
	if err !=nil {
		fmt.Println(err)
		return
	}
}

type QQRequest struct{
	PostType    string  `json:"post_type"`
	RequestType string  `json:"request_type"`
	MessageType string  `json:"message_type"`
	UserID      int     `json:"user_id"`
	GroupID     int     `json:"group_id"`
	Message     string  `json:"message"`
	RawMessage  string  `json:"raw_message"`
	Flag        string  `json:"flag"`
	File        string  `json:"file"`
}

func ReceivePost(ctx *gin.Context){
	var Request QQRequest
	err := ctx.BindJSON(&Request)
	if err !=nil {
		fmt.Println(err)
		return
	}
	if Request.RequestType == "friend" {
		AddFriend(Request,ctx)
	}
	if Request.MessageType == "group" {
		Handldgrouprequest(Request)
	}
	if Request.MessageType == "private" {
		Handldprivaterequest(Request)
	}
}

type AddFriendRequest struct{
	Flag        string `json:"flag"`
	Approve     bool   `json:"approve"`
}

func AddFriend(Request QQRequest, ctx *gin.Context){
	client := &http.Client{}
	msg := AddFriendRequest{
		Flag: Request.Flag ,
		Approve: true,
	}
	Requestbody := new(bytes.Buffer)
	err := json.NewEncoder(Requestbody).Encode(msg)
	if err !=nil {
		fmt.Println(err)
		return
	}
	rRequest, err :=http.NewRequest("POST" , "http://175.24.41.84:5700/set_friend_add_request" , Requestbody)
	if err !=nil {
		fmt.Println(err)
		return
	}
	rRequest.Header.Set("Content-Type","application/json")
	_, err = client.Do(rRequest)
	if err !=nil {
		fmt.Println(err)
		return
	}
	Sendmessage("初次见面，请回复\"help\"获取DDL提醒功能体验", Request.UserID)
}

func Handldgrouprequest(Request QQRequest){
	cur := Request.Message
	if cur == "作业" {
		Sendgroupmessage(Homework, Request.GroupID)
		return
	}
}

func Handldprivaterequest(Request QQRequest){
	cur := Request.Message
	if cur == "help" {
		Sendmessage(Help, Request.UserID)
		return
	}
	if cur == "tasks" {
		DDLlist , err :=model.FindDDL(Request.UserID)
		if err != nil{
			fmt.Println(err)
			return
		}
		if len(DDLlist)==0 {
			Sendmessage("亲，你还没有DDL呢", Request.UserID)
			return
		}
		send := ""
		for _,u:=range DDLlist {
			send = send +  u.Task + ",截止日期："+ u.ExpireTime + "编号为" + fmt.Sprintf("%d", u.Number) + "\n"
		}
		Sendmessage(send, Request.UserID)
		return
	}

	res := strings.Split(cur, ",")
	if res[0] == "add"{
		if len(res) != 5 {
			Sendmessage("格式错误", Request.UserID)
			return
		}
		err := model.AddDDL(Request.UserID,res[1],res[2],res[3],res[4])
		if err != nil {
			Sendmessage(err.Error(), Request.UserID)
			return
		}
		Sendmessage("添加成功", Request.UserID)
	}
	if res[0] == "del"{
		if len(res) != 2 {
			Sendmessage("格式错误", Request.UserID)
			return
		}
		DDLlist , err :=model.FindDDL(Request.UserID)
		if err != nil{
			fmt.Println(err)
			return
		}
		if len(DDLlist)==0 {
			Sendmessage("亲，你还没有DDL呢", Request.UserID)
			return
		}

		err = model.DelDDL(Request.UserID, res[1])
		if err != nil{
			Sendmessage(err.Error(), Request.UserID)
			return
		}
		Sendmessage("删除成功", Request.UserID)
	}
}

func Calluser(UserID int){
	DDLlist,err := model.FindDDL(UserID)
	if err != nil{
		fmt.Println(err)
		return
	}
	local,err := time.LoadLocation("Asia/Shanghai")
	if err != nil{
		fmt.Println(err)
		return
	}
	cnt :=0
	send := "您的下列任务即将到期，现在的时间是：\n"
	send = send + time.Now().Format("2006-01-02 15:04:05") + "\n"
	for _,u := range DDLlist{
		DDL,err:=time.ParseInLocation("2006-01-02 15:04:05", u.ExpireTime, local)
		if err != nil{
			fmt.Println(err)
			return
		}
		curtime := time.Now().Unix()
		diff := DDL.Unix() - curtime
		if diff<0 {
			err = model.DelDDL(UserID, fmt.Sprintf("%d", u.Number))
		}
		if (diff/3600 >= int64(u.Ahead)) || (diff<0) {continue}

		model.Calc(UserID,u.Intervalc-1,u.Number)
		if u.Intervalc == 0 {
			model.Calc(UserID,u.Interval,u.Number)
		}else {
			continue
		}
		send = send +  u.Task + ",Stop At："+ u.ExpireTime + "\n"
		cnt++
	}
	if cnt==0 {return}
	Sendmessage(send,UserID)
}
