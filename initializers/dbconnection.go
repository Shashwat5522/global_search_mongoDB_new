package initializers

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Manager struct {
	Connection *mongo.Client
	Context    context.Context
	Cancel     context.CancelFunc
}

var (
	Mgr  *Manager
	once sync.Once
)

func GetInstance()*Manager {
	if Mgr == nil {
		once.Do(
			func() {
				Mgr = DBConnect()

			})
	} else {
		fmt.Println("single instance already created")

	}
	return Mgr
}

func DBConnect() *Manager {
	uri := "localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("%s%s", "mongodb://", uri)))
	if err != nil {
		log.Fatal(err)
	}
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(context)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected!!!")
	mgr:=Manager{Connection:client,Context:context,Cancel:cancel}
	return &mgr

}
