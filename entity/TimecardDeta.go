package entity
import (
)
 
type TimecardDeta struct {
    DateDd     string `bson:"dateDd"`
    BegTm      string `bson:"begTm"`
    EndTm      string `bson:"endTm"`
    RstTm      string `bson:"rstTm"`
    RemarksTxt string `bson:"remarksTxt"`
    Pj1Tm      string `bson:"pj1Tm"`
    Pj2Tm      string `bson:"pj2Tm"`
    Pj3Tm      string `bson:"pj3Tm"`
    PjATm      string `bson:"pjATm"`
    PjKTm      string `bson:"pjKTm"`
}
