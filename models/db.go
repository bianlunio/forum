package models

import (
	"github.com/globalsign/mgo"
)

var Session *mgo.Session

func init() {
	session, err := mgo.Dial("localhost")
	Must(err)
	session.SetMode(mgo.Monotonic, true)
	Session = session
}

func Must(err error) {
	if err != nil {
		panic(err.Error())
	}
}
