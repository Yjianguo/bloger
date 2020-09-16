package controller

import (
	"code.lssj.me/bloger/model"
	"code.lssj.me/bloger/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 抽象出err
func errMsg(c *gin.Context, err error) {
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
}

// 首页handler
func IndexHandler(c *gin.Context) {
	// 获取文章列表数据
	articleList, err := service.GetAllArticleList(1, 10)
	errMsg(c, err)

	//获取分类列表信息
	categoryList, err := service.GetAllCategoryList()
	errMsg(c, err)

	//页面返回
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleList,
		"category_list": categoryList,
	})

}

// 文章详情Handler
func ArticleDetailHandler(c *gin.Context) {
	//根据文章id 获取文章内容
	articleId, err := strconv.ParseInt(c.Query("article_id"), 10, 64)
	errMsg(c, err)

	articleDetail, err := service.GetArticleDetailById(articleId)
	errMsg(c, err)

	//articleRecordList, err := service.GetArticleRecordByCategoryId(articleDetail.Category.CategoryId)
	//errMsg(c, err)
	//
	//获取分类列表信息
	categoryList, err := service.GetAllCategoryList()
	errMsg(c, err)

	// 获取评论列表
	commentList, err := service.GetAllComments(articleId)
	errMsg(c, err)

	//页面返回
	c.HTML(http.StatusOK, "views/detail.html", gin.H{
		"detail":       articleDetail,
		"category":     categoryList,
		"prev":         articleDetail,
		"next":         articleDetail,
		"comment_list": commentList,
		"article_id":   articleId,
	})

}

// 根据分类展示文章列表
func ArticleListByCategory(c *gin.Context) {
	articleId, err := strconv.ParseInt(c.Query("category_id"), 10, 64)
	errMsg(c, err)

	articleRecordList, err := service.GetArticleRecordByCategoryId(articleId)
	errMsg(c, err)

	categoryList, err := service.GetAllCategoryList()
	errMsg(c, err)

	c.HTML(http.StatusOK, "views/article.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}

// 投稿页面
func PostPage(c *gin.Context) {
	categoryList, err := service.GetAllCategoryList()
	errMsg(c, err)

	c.HTML(http.StatusOK, "views/post_article.html", &categoryList)
}

//文章投稿
func PostArticle(c *gin.Context) {
	categoryId, err := strconv.ParseInt(c.PostForm("category_id"), 10, 64)
	errMsg(c, err)

	articleInfo := model.ArticleInfo{
		CategoryId:   categoryId,
		Title:        c.PostForm("title"),
		Username:     c.PostForm("author"),
		ViewCount:    0,
		CommentCount: 0,
		Summary:      c.PostForm("content")[1:100],
	}
	articleDetail := &model.ArticleDetail{
		ArticleInfo: articleInfo,
		Content:     c.PostForm("content"),
	}
	_, err = service.WriteArticle(articleDetail)
	errMsg(c, err)

	categoryList, err := service.GetAllCategoryList()
	errMsg(c, err)

	c.HTML(http.StatusOK, "views/post_article.html", &categoryList)
}

// About页面
func AboutMe(c *gin.Context) {
	c.HTML(http.StatusOK, "views/about.html", gin.H{})
}

// 留言页面
func LeavePage(c *gin.Context) {
	leavesList, err := service.GetLeavesList()
	errMsg(c, err)
	c.HTML(http.StatusOK, "views/gbook.html", &leavesList)
}

//发表留言
func WriteLeaves(c *gin.Context) {
	content, _ := c.GetPostForm("comment")
	email, _ := c.GetPostForm("email")
	username, _ := c.GetPostForm("author")
	leaves := &model.Leaves{
		Username: username,
		Content:  content,
		Email:    email,
	}
	_, err := service.WriteLeaves(leaves)
	errMsg(c, err)

	leavesList, err := service.GetLeavesList()
	errMsg(c, err)

	c.HTML(http.StatusOK, "views/gbook.html", &leavesList)
}

// 发表评论
func PostComment(c *gin.Context) {
	articleId, err := strconv.ParseInt(c.PostForm("article_id"), 10, 64)
	errMsg(c, err)
	fmt.Println(articleId)
	comment := &model.Comment{
		Username:  c.PostForm("author"),
		Content:   c.PostForm("comment"),
		ArticleID: articleId,
	}
	_, err = service.PostComment(comment)
	errMsg(c, err)

	c.Redirect(http.StatusMovedPermanently, "/")

}

// 获取评论列表
