package service

import (
	"code.lssj.me/bloger/dao"
	"code.lssj.me/bloger/model"
)

// 发表评论
func PostComment(comment *model.Comment) (commentId int64, err error) {
	commentId, err = dao.InsertComment(comment)
	if err != nil {
		panic(err)
		return
	}
	return
}

// 获取评论列表
func GetAllComments(articleId int64) (commentList []*model.Comment, err error) {
	commentList, err = dao.GetAllCommentListByArticleId(articleId)
	if err != nil {
		panic(err)
		return
	}
	return
}
