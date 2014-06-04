package main

import (
	"fmt"
	"github.com/astaxie/goredis"
)

var (
	client goredis.Client
)

type User struct {
	Name string
}

func setupDB() {
	client.Addr = "127.0.0.1:6379"
//	client, err := DialURL("tcp://auth:password@127.0.0.1:6379/0?timeout=10s&maxidle=1")
}

func insertData(){
	fmt.Println("In instert data.")
	vals := []string{"pravin", "ankit", "abhay", "aditya", "avinaw"}
	for _, v := range vals {
		client.Rpush("username", []byte(v))
	}
}

func getData() {
	fmt.Println("In get data.")
	username,_ := client.Lrange("username", 0, 10000)
	for i, v := range username {
		println("Name", i,":",string(v))
	}
}

func main() {
	setupDB()
	fmt.Println("Client: ", client)
	insertData()
	getData()
}
