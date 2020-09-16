package test

import (
	"code.lssj.me/bloger/dao"
	"code.lssj.me/bloger/dao/db"
	"log"
	"testing"
)

func init() {
	err := db.Init()
	if err != nil {
		log.Fatal("数据库初始化失败", err)
	}
}

// 测试单个分类
func TestGetCategoryById(t *testing.T) {
	category, err := dao.GetCategoryById(1)
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("Category: %#v", category)
}

// 测试多个分类
func TestGetCategoryList(t *testing.T) {
	ids := []int64{1, 2, 3}
	categoryList, err := dao.GetCategoryList(ids)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range categoryList {
		t.Logf("category: %#v", v)
	}

}

// 测试全部分类
func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := dao.GetAllCategoryList()
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range categoryList {
		t.Logf("category: %#v", v)
	}
}
