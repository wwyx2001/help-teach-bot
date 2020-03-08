package model

import(
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strconv"
	"time"
)

var (
	Collection     *mongo.Collection
	UserCollection *mongo.Collection
)

func init(){
	client , err:=mongo.NewClient(options.Client().ApplyURI("mongodb://175.24.41.84:27017"))
	if err !=nil{
		fmt.Println(err)
		return
	}
	ctx,_:=context.WithTimeout(context.Background(), time.Second*30)//超时取消
	err = client.Connect(ctx)
	if err !=nil{
		fmt.Println(err)
		return
	}
	if err = client.Ping(ctx,readpref.Primary()); err != nil{
		fmt.Println(err)
		return
	}
	Collection = client.Database("test").Collection("DDL")
	UserCollection = client.Database("test").Collection("User")
}

type NameList struct {
	UserID int `bson:"user_id"`
}

func FindName() ([]NameList, error){
	var Backmessage []NameList
	filter := bson.D{}
	res,err := UserCollection.Find(context.Background(), filter)
	if err !=nil{
		fmt.Println(err)
		return nil,err
	}
	for res.Next(context.Background()) {
		var cur NameList
		err := res.Decode(&cur)
		if err != nil {
			return nil,err
		}
		Backmessage = append(Backmessage, cur)
	}
	return Backmessage,nil
}

type DDL struct {
	ExpireTime string `bson:"expire_time"`
	UserID     int    `bson:"user_id"`
	Task       string `bson:"context"`
	Number     int    `bson:"number"`
	Ahead      int    `bson:"ahead"`
	Interval   int    `bson:"interval"`
	Intervalc  int    `bson:"intervalc"`
}

func Calc(UserID int,num int,id int){
	filter := bson.D{{"user_id", UserID},{"number", id}}
	updater := bson.M{"$set": bson.M{"intervalc": num}}
	_, err := Collection.UpdateOne(context.Background(), filter, updater)
	if err != nil{
		fmt.Println(err)
	}
	return
}

func FindDDL(UserID int) ([]DDL,error){
	var Backmessage []DDL
	filter := bson.D{{"user_id",UserID}}
	res,err := Collection.Find(context.Background(), filter)
	if err !=nil{
		fmt.Println(err)
		return nil,err
	}
	for res.Next(context.Background()) {
		var cur DDL
		err := res.Decode(&cur)
		if err != nil {
			return nil,err
		}
		Backmessage = append(Backmessage, cur)
	}
	return Backmessage,nil
}

type Usermessage struct {
	UserID  int `bson:"user_id"`
	UserCnt int `bson:"user_cnt"`
}

func CheckExist(name int) bool {
	var User Usermessage
	filter := bson.D{{"user_id", name}}
	err := UserCollection.FindOne(context.Background(), filter).Decode(&User)
	if err != nil {
		return false
	}
	return true
}

func DelDDL(UserID int, Number string) error {
	number, err := strconv.Atoi(Number)//字符串转数字
	if err != nil {
		return err
	}
	_, err = Collection.DeleteMany(context.Background(), bson.D{{"number", number}, {"user_id", UserID}})
	if err != nil {
		return err
	}
	return nil
}

func AddDDL(UserID int, Context string, Time string,Hour string,Minute string) error{
	local, err := time.LoadLocation("Asia/Shanghai")
	IntHour,err := strconv.Atoi(Hour)
	IntMinute,err := strconv.Atoi(Minute)
	if err != nil {return err}
	RealTime, err := time.ParseInLocation("2006-01-02 15:04:05", Time, local)
	if err != nil {
		return errors.New("日期格式不符合标准qwq")
	}
	curtime := time.Now().Unix()
	diff := RealTime.Unix() - curtime
	if diff < 0 {
		return errors.New("宁的DDL时间已经过了呀QAQ")
	}
	var cnt int
	filter := bson.D{{"user_id", UserID}}
	if CheckExist(UserID) {//查查该用户是否存在
		var tmp Usermessage
		err := UserCollection.FindOne(context.Background(), filter).Decode(&tmp)
		if err != nil {
			fmt.Println(err)
			return err
		}
		cnt = tmp.UserCnt + 1
	} else {
		tmp := Usermessage{
			UserID:  UserID,
			UserCnt: 0,
		}
		_, err := UserCollection.InsertOne(context.Background(), tmp)
		if err != nil {
			return err
		}
		cnt = 1
	}
	updater := bson.M{"$set": bson.M{"user_cnt": cnt}}
	_, err = UserCollection.UpdateOne(context.Background(), filter, updater)
	if err != nil {
		return err
	}
	cur := DDL{
		ExpireTime:Time,
		UserID:UserID,
		Task:Context,
		Number:cnt,
		Ahead:IntHour,
		Interval:IntMinute-1,
		Intervalc:IntMinute-1,
	}
	_, err = Collection.InsertOne(context.Background(), cur)
	if err != nil {
		return err
	}
	return nil
}
