package service

import (
	"code.lssj.me/bloger/dao"
	"code.lssj.me/bloger/model"
	"log"
)

// 获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = dao.GetAllCategoryList()
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
