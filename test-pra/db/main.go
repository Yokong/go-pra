package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

type Profile struct {
	Uin    int
	Name   string
	Gender string
	Height int
	Age    int
}

func main() {
	session, err := mgo.Dial("")
	if err != nil {
		fmt.Println(err)
	}
	defer session.Clone()

	profile := session.DB("account").C("profile")
	// user1 := Profile{
	// 	Uin:    1001,
	// 	Name:   "KangKang",
	// 	Gender: "F",
	// }
	// profile.Insert(user1)

	// profile.Update(
	// 	bson.M{"uin": 1000},
	// 	bson.M{"$set": bson.M{"height": 178}},
	// )

	var results []Profile
	err = profile.Find(nil).All(&results)
	if err != nil {
		fmt.Println("----")
		fmt.Println(err)
	}

	fmt.Println(results)
}
