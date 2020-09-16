package test

import (
	"code.lssj.me/bloger/dao"
	"code.lssj.me/bloger/dao/db"
	"code.lssj.me/bloger/model"
	"log"
	"testing"
	"time"
)

func init() {
	err := db.Init()
	if err != nil {
		log.Fatal("数据库初始化失败", err)
	}
}

func TestInsertComment(t *testing.T) {
	comment := &model.Comment{
		Content:    "jjssdskjncksksdkskds",
		Username:   "lnkkhj",
		CreateTime: time.Now(),
		ArticleID:  3,
	}
	commentId, err := dao.InsertComment(comment)
	if err != nil {
		panic(err)
		return
	}
	t.Logf("CommentId: %d\n", commentId)
}

func TestGetAllCommentListByArticleId(t *testing.T) {
	commentList, err := dao.GetAllCommentListByArticleId(1)
	if err != nil {
		panic(err)
		return
	}
	for _, v := range commentList {
		t.Logf("comment: %#v\n", v)
	}
}
