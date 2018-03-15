package main

import (
  "fmt"
  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "./entity"
)

func main() {

    session, err := mgo.Dial("mongodb://timeadmin:admin123@localhost:28011,localhost:28012,localhost:28013/timecard")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    db := session.DB("timecard")

    coll := db.C("timecardUser")

    userList := []entity.TimecardUser{}

    err = coll.Find(bson.M{}).All(&userList)
    if err != nil {
        panic(err)
    }

    for _, row := range userList {
        fmt.Printf("%s %s %s %s\n", row.Id, row.LoginId, row.UserNmJa, row.CreatedAt)
    }


}
