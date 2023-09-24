package initialize

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/global"
	"time"
)

func NewMongo() *global.MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	username := global.ADTPL_CONFIG.Mongo.UserName
	password := global.ADTPL_CONFIG.Mongo.Password
	path := global.ADTPL_CONFIG.Mongo.Path
	port := global.ADTPL_CONFIG.Mongo.Port
	database := global.ADTPL_CONFIG.Mongo.DbName
	fmt.Println(username, password, path, port, database)
	// mongodb://myUserAdmin:wind4Mongo@119.3.155.86:27017/blc2
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+username+":"+password+"@"+path+":"+port))
	if err != nil {
		fmt.Println("failed to connect to mongodb")
		panic("connect mongodb error:%s" + err.Error())
	}
	fmt.Println("succeed to connect to mongodb")
	db := client.Database(database)
	session, errSession := db.Client().StartSession()
	if errSession != nil {
		panic("start mongodb session error:" + errSession.Error())
	}
	return &global.MongoDB{Client: client, Db: db, Session: session}
}
