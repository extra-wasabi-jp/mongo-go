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

    fp, err := os.Open("timecard.tsv")
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
    coll := db.C("timecard")


    reader := csv.NewReader(fp)
    reader.Comma = '\t'

    var ownerId = "direct.k"
    var dateYm = "201803"
    var noteTxt = "ノート"
    var pj1Code = "AAA18001"
    var pj2Code = "BBB18002"
    var pj3Code = "CCC18003"

    dateList := []entity.TimecardDeta{}
    for {
        row, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }
        day := entity.TimecardDeta{
             DateDd: row[0],
             BegTm: row[1],
             EndTm: row[2],
             RstTm: row[3],
             RemarksTxt: row[4],
             Pj1Tm: row[5],
             Pj2Tm: row[6],
             Pj3Tm:row[7], 
             PjATm:row[8], 
             PjKTm:row[9], 
        }
        dateList = append(dateList, day)
        fmt.Printf("%s\n", day)

    }

    timecard := &entity.Timecard{
        Id:        bson.NewObjectId(),
        OwnerId:   ownerId,
        DateYm:    dateYm,
        DateList:  dateList,
        NoteTxt:   noteTxt,
        Pj1Code:   pj1Code,
        Pj2Code:   pj2Code,
        Pj3Code:   pj3Code,

        VersionNo: 1,
        CreatedAt: time.Now(),
        CreatedBy: "init",
        UpdatedAt: time.Now(),
        UpdatedBy: "init",
        IsActived: "1",
    }
    err = coll.Insert(timecard)
    if err != nil {
         panic(err)
    }

}
