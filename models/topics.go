package models

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/orm"
)

type Topic struct {
	Id            int       `sql:"id" json:"id"`
	Title         string    `sql:"title" json:"title" binding:"required"`
	Content       string    `sql:"content" json:"content" binding:"required"`
	Top           bool      `sql:"top,notnull" json:"top"`
	Good          bool      `sql:"good,notnull" json:"good"`
	Lock          bool      `sql:"lock,notnull" json:"lock"`
	ReplyCount    int32     `sql:"replyCount,notnull" json:"replyCount"`
	VisitCount    int32     `sql:"visitCount,notnull" json:"visitCount"`
	CreatedAt     time.Time `sql:"createdAt" json:"createdAt"`
	UpdatedAt     time.Time `sql:"updatedAt" json:"updatedAt"`
	LastRepliedAt time.Time `sql:"lastRepliedAt" json:"lastRepliedAt"`
	TabId         int       `sql:"tabId" pg:",fk:Tab" json:"tabId"`
	Deleted       bool      `sql:"deleted,notnull" json:"deleted"`
	DeletedAt     time.Time `sql:"deletedAt" json:"deletedAt"`
}

func (t Topic) List(page int, limit int) (res gin.H, err error) {
	var topics []Topic
	query := db.Model(&topics)
	total, err := query.Count()
	if err != nil {
		return nil, err
	}
	offset := (page - 1) * limit
	err = db.Model(&topics).
		Limit(limit).
		Offset(offset).
		Select()
	if err != nil {
		return nil, err
	}
	res = gin.H{
		"pagination": Pagination{
			Page:  page,
			Limit: limit,
			Total: total,
		},
		"list": topics,
	}
	return res, nil
}

func (t Topic) Create(topic Topic) (Topic, error) {
	err := db.Insert(&topic)
	handleDBError(err)
	return topic, nil
}

func (t Topic) Detail(id int) (Topic, error) {
	topic := Topic{Id: id}
	err := db.Select(&topic)
	return topic, err
}

func (t Topic) Update(id int, topic Topic) (Topic, error) {
	topic.Id = id
	err := db.Update(&topic)
	return topic, err
}

func (t Topic) SoftDelete(id int) (int, error) {
	topic := Topic{
		Id: id,
		Deleted: true,
	}
	res, err := db.Model(&topic).Column("deleted", "deletedAt").Update()
	return res.RowsAffected(), err
}

func (t Topic) String() string {
	return fmt.Sprintf("Topic<Id=%d Title=%q>", t.Id, t.Title)
}

func (t *Topic) BeforeInsert(db orm.DB) error {
	if t.CreatedAt.IsZero() {
		now := time.Now()
		t.CreatedAt = now
		t.UpdatedAt = now
	}
	return nil
}

func (t *Topic) BeforeUpdate(db orm.DB) error {
	now := time.Now()
	t.UpdatedAt = now
	if t.Deleted {
		t.DeletedAt = now
	}
	return nil
}

type Tab struct {
	Id   int    `sql:"id"`
	Name string `sql:"name"`
}

func (tab Tab) String() string {
	return fmt.Sprintf("Tab<Id=%d Name=%q>", tab.Id, tab.Name)
}
