package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"image"
)

type Record struct {
	Id                    int
	VId                   int
	VisitorName           string
	VisitorEmail          string
	VisitorPhone          string
	VisitorAddress        string
	VisitorImagePath      string
	HostName              string
	HostEmail             string
	CompanyRepresentating string
	TrackVisitors         []TrackVisitor `gorm:"ForeignKey:VId"`
}
