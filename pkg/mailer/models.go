package mailer

import (
	"database/sql"
	"time"
)

type Notificationattacments struct {
	Id             int32
	Notificationid int32
	Filename       string
	Summary        string
	Fileext        string
	Physicalpath   string
	Filecontent    string
}

type Notificationbody struct {
	Id             int32
	Notificationid int32
	Title          string
	Body           string
}

type Notifications struct {
	Id                 int32
	Receipents         string
	Senddate           time.Time
	Sentdate           sql.NullTime
	Issent             bool
	Notificationtypeid int32
}

type Notificationtype struct {
	Id   int32
	Name string
}

type Notificationsdto struct {
	Id                     int32
	Receipents             []Receipents
	Senddate               time.Time
	Sentdate               sql.NullTime
	Issent                 bool
	Notificationtypeid     int32
	Notificationtype       Notificationtype
	Notificationbody       Notificationbody
	Notificationattacments []Notificationattacments
}

type Receipents struct {
	Fullname string
	Email    string
	Phone    string
}
