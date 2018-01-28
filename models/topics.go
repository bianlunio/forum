package models

import (
	"fmt"
	"net/url"
	"time"

	"forum/utils"

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

func (t Topic) List(values url.Values) gin.H {
	var topics []Topic
	query := db.Model(&topics)
	total, err := query.Count()
	handleDBError(err)
	err = db.Model(&topics).
		Apply(orm.Pagination(values)).
		Select()
	handleDBError(err)
	page := utils.String2Int(values.Get("page"))
	limit := utils.String2Int(values.Get("limit"))
	return gin.H{
		"pagination": Pagination{
			Page:  page,
			Limit: limit,
			Total: total,
		},
		"list": topics,
	}
}

func (t Topic) Create() Topic {
	err := db.Insert(&t)
	handleDBError(err)
	return t
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
	t.UpdatedAt = time.Now()
	return nil
}

type Tab struct {
	Id   int    `sql:"id"`
	Name string `sql:"name"`
}

func (tab Tab) String() string {
	return fmt.Sprintf("Tab<Id=%d Name=%q>", tab.Id, tab.Name)
}
