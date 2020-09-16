package model

import "time"

type Comment struct {
	Id         int64     `db:"id"`
	Content    string    `db:"content"`
	Username   string    `db:"username"`
	ArticleID  int64     `db:"article_id"`
	CreateTime time.Time `db:"create_time"`
}
