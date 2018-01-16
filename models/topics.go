package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Topic struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Content     string        `bson:"content" json:"content"`
	Top         bool          `bson:"top" json:"top"`
	Good        bool          `bson:"good" json:"good"`
	Lock        bool          `bson:"lock" json:"lock"`
	ReplyCount  int           `bson:"replyCount" json:"replyCount"`
	VisitCount  int           `bson:"visitCount" json:"visitCount"`
	CreateAt    *time.Time    `bson:"createAt" json:"createAt"`
	UpdateAt    *time.Time    `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	LastReplyAt *time.Time    `bson:"lastReplyAt,omitempty" json:"lastReplyAt,omitempty"`
	Tab         string        `bson:"tab" json:"tab,omitempty"`
	Deleted     bool          `bson:"deleted" json:"deleted,omitempty"`
	DeleteAt    *time.Time    `bson:"deleteAt,omitempty" json:"deleteAt,omitempty"`
}

type Topics []Topic

func (t *Topic) Create(title string, content string, tab string) *Topic {
	now := time.Now()
	t.Id = bson.NewObjectId()
	t.Title = title
	t.Content = content
	t.Tab = tab
	t.Top = false
	t.Good = false
	t.Lock = false
	t.ReplyCount = 0
	t.VisitCount = 0
	t.CreateAt = &now
	t.Deleted = false
	return t
}
