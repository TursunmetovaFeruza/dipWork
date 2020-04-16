package model

import (
	"time"

	"github.com/lib/pq"
)

//Event
type Event struct {
	ID        int       `json:"id" bson:"id"`
	LectionId int       `json:"lectionId" bson:"lectionid"`
	MasterId  int       `json:"masterid" bson:"masterid"`
	StartTime time.Time `json:"starttime" bson:"starttime"`
	EndTime   time.Time `json:"endtime" bson:"endtime"`
	StudentId []int     `json:"students" bson:"students"`
}

type PostEvent struct {
	ID        int            `json:"id" bson:"id"`
	LectionId string         `json:"lectionId" bson:"lectionid"`
	MasterId  string         `json:"masterid" bson:"masterid"`
	StartTime map[string]int `json:"starttime" bson:"starttime"`
	EndTime   map[string]int `json:"endtime" bson:"endtime"`
	StudentId []string       `json:"studentid" bson:"studentid"`
}

type GetEvent struct {
	ID        int           `json:"id" bson:"id"`
	Lection   string        `json:"lection" bson:"lection"`
	Master    string        `json:"master" bson:"master"`
	StartTime string        `json:"starttime" bson:"starttime"`
	EndTime   string        `json:"endtime" bson:"endtime"`
	StudentId pq.Int64Array `json:"students" bson:"students"`
	Students  []string      `json:"studentlist" bson:"studentlist"`
	Attend    []string      `json:"attend" bson:"attend"`
	Apsent    []string      `json:"apsent" bson:"apsent"`
}

//items of Event
type Events []Event
