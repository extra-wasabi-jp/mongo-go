package entity
import (
    "time"
    "gopkg.in/mgo.v2/bson"
)
 
type Timecard struct {
    Id bson.ObjectId `bson:"_id"`

    OwnerId   string   `bson:"ownerId"`
    DateYm    string   `bson:"dateYm"`
    DateList  [] TimecardDeta   `bson:"dateList"`
    NoteTxt   string   `bson:"noteTxt"`
    Pj1Code   string   `bson:"pj1Code"`
    Pj2Code   string   `bson:"pj2Code"`
    Pj3Code   string   `bson:"pj3Code"`

    VersionNo int64  `bson:"versionNo"`
    CreatedAt time.Time   `bson:"createdAt"`
    CreatedBy string `bson:"createdBy"`
    UpdatedAt time.Time   `bson:"updatedAt"`
    UpdatedBy string `bson:"updatedBy"`
    IsActived string `bson:"isActived"`
}

