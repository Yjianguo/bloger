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

func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{
		ArticleInfo: model.ArticleInfo{
			CategoryId:   6,
			Summary:      "kkkk",
			Title:        "html",
			ViewCount:    1111,
			CommentCount: 1111,
			Username:     "xiaowang",
			CreateTime:   time.Now(),
		},
		Content: "KDFSKFSJHFSKJHSKJKKSJDFHSJKKSDJKFSKDFKKSFKSKJIU3HIUHSNKJFNJSNJKFBBS",
		Category: model.Category{
			CategoryId:   6,
			CategoryName: "Html",
			CategoryNo:   106,
		},
	}
	articleId, err := dao.InsertArticle(article)
	if err != nil {
		return
	}
	t.Logf("ArticleId: %d\n", articleId)

}
func TestGetArticleList(t *testing.T) {
	articleList, err := dao.GetArticleRecordList(1, 2)
	if err != nil {
		panic(err)
	}
	for _, v := range articleList {
		t.Logf("ArticleInfo: %#v", v)
	}
}

func TestGetArticleDetailById(t *testing.T) {
	articleDetail, err := dao.GetArticleDetailById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("ArticleDetail: %#v", articleDetail)
}

func TestGetArticleListByCategoryId(t *testing.T) {
	articleInfos, err := dao.GetArticleRecordListByCategoryId(1)
	if err != nil {
		panic(err)
	}
	for _, v := range articleInfos {
		t.Logf("ArticleList: %#v", v)
	}
}
