package main

import (
  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type Person struct {
    ID       bson.ObjectId `bson: "_id"`
    LoginId  string      `bson: "loginId"`
    UserJaNm string      `bson: "userJaNm"`
}

func main() {

    session, _ := mgo.Dial("mongodb://192.168.1.201:28001/webdisk")
    defer session.Close()

    db := session.DB("webdisk")
    user := &Person{bson.NewObjectId(), "hogehoge", "Hoge, Hoge"}
        
    db.C("hogehoge").Insert(user)

}
