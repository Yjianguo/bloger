package dao

import (
	"code.lssj.me/bloger/dao/db"
	"code.lssj.me/bloger/model"
	"log"
	"time"
)

// 发表评论
func InsertComment(comment *model.Comment) (commentId int64, err error) {
	sqlStr := `insert into comment (content, username, create_time, article_id) values (?,?,?,?)`
	result, err := db.DB.Exec(sqlStr, comment.Content, comment.Username, time.Now(), comment.ArticleID)
	if err != nil {
		log.Fatal(err)
		return
	}
	commentId, err = result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

//查看文章下的评论列表
func GetAllCommentListByArticleId(articleId int64) (commentList []*model.Comment, err error) {
	sqlStr := `select content, username, create_time, article_id from comment where article_id=? order by create_time desc limit 20`
	err = db.DB.Select(&commentList, sqlStr, articleId)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

// 查看所有评论列表
func GetAllCommentList() (commentList []*model.Comment, err error) {
	return
}
