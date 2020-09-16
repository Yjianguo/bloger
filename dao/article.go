package dao

import (
	"code.lssj.me/bloger/dao/db"
	"code.lssj.me/bloger/model"
	"log"
)

// 插入文章-投稿
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	//参数验证
	if article == nil {
		log.Fatal("Params error.", err)
		return
	}
	sqlStr := `insert into article(category_id,content,summary,title,view_count,comment_count,username) 
values (?,?,?,?,?,?,?)`
	result, err := db.DB.Exec(sqlStr, article.ArticleInfo.CategoryId,
		article.Content, article.ArticleInfo.Summary,
		article.ArticleInfo.Title, article.ArticleInfo.ViewCount,
		article.ArticleInfo.CommentCount, article.ArticleInfo.Username)
	articleId, err = result.LastInsertId()
	if err != nil {
		log.Fatal("SQL exec failed.", err)
		return
	}
	return
}

// 查询首页文章列表
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//验证参数
	if pageNum < 0 || pageSize < 0 {
		log.Fatal("Params error.", err)
		return
	}
	pageNum = (pageNum - 1) * pageSize
	sqlStr := `select a.id,a.category_id,a.title,a.view_count,a.comment_count,a.username,a.summary,a.create_time,c.category_name from article a,category c
where a.status=1 and c.id = a.category_id order by a.create_time desc limit ?,?`
	err = db.DB.Select(&articleRecordList, sqlStr, pageNum, pageSize)
	return
}

// 查询文章详情
func GetArticleDetailById(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	// 检验参数
	if articleId < 0 {
		log.Fatal("Params error.", err)
		return
	}
	articleDetail = &model.ArticleDetail{}
	sqlStr := `select a.id,a.category_id,a.title,a.view_count,a.comment_count,a.username,a.summary,a.create_time,a.content,c.category_name from article a, category c
where a.id=? and a.category_id=c.id`
	err = db.DB.Get(articleDetail, sqlStr, articleId)

	return
}

// 按照分类查询文章列表
func GetArticleRecordListByCategoryId(categoryId int64) (articleRecordList []*model.ArticleRecord, err error) {
	// 参数校验
	if categoryId < 0 {
		log.Fatal("Params error.", err)
		return
	}
	sqlStr := `select a.id,a.category_id,a.title,a.view_count,a.comment_count,a.username,a.summary,a.create_time, c.category_name from article a, category c
where a.status=1 and a.category_id=c.id and a.category_id =? order by a.create_time desc limit 10`
	err = db.DB.Select(&articleRecordList, sqlStr, categoryId)

	return
}
