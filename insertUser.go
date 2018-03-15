package main

import (
  "os"
  "io"
  "fmt"
  "time"
  "encoding/csv"
  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "./entity"
)

func main() {

    fp, err := os.Open("userlist.tsv")
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    session, err := mgo.Dial("mongodb://timeadmin:admin123@localhost:28011,localhost:28012,localhost:28013/timecard")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    db := session.DB("timecard")
    coll := db.C("timecardUser")


    reader := csv.NewReader(fp)
    reader.Comma = '\t'
    for {
        row, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }
        fmt.Printf("[%s] [%s] [%s]\n", row[0], row[3], row[4])

        user := &entity.TimecardUser{
            Id: bson.NewObjectId(),
             LoginId: row[0],
             Password: row[1],
             PasswordLstChgDt: row[2],
             UserNmJa: row[3],
             UserNmKn: row[4],
             UserNmEn: row[5],
             JoinDt: row[6],
             QuitDt:row[7], 
             VersionNo: 1,
             CreatedAt: time.Now(),
             CreatedBy: "init",
             UpdatedAt: time.Now(),
             UpdatedBy: "init",
             IsActived: "1",
        }

        err = coll.Insert(user)
        if err != nil {
             panic(err)
        }
    }

}
