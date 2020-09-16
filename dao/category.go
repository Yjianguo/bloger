package dao

import (
	"code.lssj.me/bloger/dao/db"
	"code.lssj.me/bloger/model"
	"github.com/jmoiron/sqlx"
	"log"
)

//添加分类
func InsertCategory(category *model.Category) (categoryId int64, err error) {
	sql := "insert into category (category_name, category_no) values (?,?)"
	result, err := db.DB.Exec(sql, category.CategoryName, category.CategoryNo)
	if err != nil {
		log.Fatal("SQL exec failed.", err)
		return
	}
	categoryId, err = result.LastInsertId()
	if err != nil {
		log.Fatal("Insert failed.", err)
		return
	}
	return
}

// 查询单个分类
func GetCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sql := "select id, category_name, category_no from category where id=?"
	err = db.DB.Get(category, sql, id)
	if err != nil {
		log.Fatal("Query failed.", err)
		return
	}
	return
}

//查询多个分类
func GetCategoryList(ids []int64) (categoryList []*model.Category, err error) {
	sql := "select id, category_name, category_no from category where id in (?)"
	sqlStr, args, err := sqlx.In(sql, ids)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.DB.Select(&categoryList, sqlStr, args...)
	if err != nil {
		log.Fatal("select failed", err)
		return
	}
	return
}

//查询全部分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	sql := "select id, category_name, category_no from category order by category_no asc"
	err = db.DB.Select(&categoryList, sql)
	if err != nil {
		log.Fatal("select failed", err)
		return
	}
	return
}
