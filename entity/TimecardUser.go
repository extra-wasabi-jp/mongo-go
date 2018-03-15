package entity
import (
    "time"
    "gopkg.in/mgo.v2/bson"
)
 

type TimecardUser struct {
    Id bson.ObjectId `bson:"_id"`
    LoginId string   `bson:"loginId"`
    Password string  `bson:"password"`
    PasswordLstChgDt  string `bson:"passwordLstChgDt"`
    UserNmJa string  `bson:"userNmJa"`
    UserNmKn string  `bson:"userNmKn"`
    UserNmEn string  `bson:"userNmEn"`
    JoinDt string    `bson:"joinDt"`
    QuitDt string    `bson:"quitDt"`
    VersionNo int64  `bson:"versionNo"`
    CreatedAt time.Time   `bson:"createdAt"`
    CreatedBy string `bson:"createdBy"`
    UpdatedAt time.Time   `bson:"updatedAt"`
    UpdatedBy string `bson:"updatedBy"`
    IsActived string `bson:"isActived"`
}
