package model

import "time"

// 定义Article 结构体
type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
	Summary      string    `db:"summary"`
	CreateTime   time.Time `db:"create_time"`
}

// 文章详情页实体
type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`
	Category
}

// 文章分页实体
type ArticleRecord struct {
	Category
	ArticleInfo
}
