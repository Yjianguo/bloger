package service

import (
	"code.lssj.me/bloger/dao"
	"code.lssj.me/bloger/model"
)

// 查询所有文章和对应的分类信息
func GetAllArticleList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	articleRecordList, err = dao.GetArticleRecordList(pageNum, pageSize)
	if err != nil {
		panic(err)
		return
	}
	return
}

//根据分类id,获取文章和他们对应的分类信息
func GetArticleRecordByCategoryId(categoryId int64) (articleRecordList []*model.ArticleRecord, err error) {
	articleRecordList, err = dao.GetArticleRecordListByCategoryId(categoryId)
	if err != nil {
		panic(err)
		return
	}
	return
}

// 根据文章id 获取文章详情
func GetArticleDetailById(articleId int64) (articleDetailList *model.ArticleDetail, err error) {
	articleDetailList, err = dao.GetArticleDetailById(articleId)
	if err != nil {
		panic(err)
		return
	}
	return
}

// 文章投稿
func WriteArticle(article *model.ArticleDetail) (articleId int64, err error) {
	articleId, err = dao.InsertArticle(article)
	if err != nil {
		panic(err)
		return
	}
	return
}
