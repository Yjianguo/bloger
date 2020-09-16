package service

import (
	"code.lssj.me/bloger/dao"
	"code.lssj.me/bloger/model"
)

// 写留言

func WriteLeaves(leaves *model.Leaves) (leaveId int64, err error) {
	leaveId, err = dao.InsertLeaves(leaves)
	if err != nil {
		panic(err)
		return
	}
	return
}

//展示留言
func GetLeavesList() (leavesList []*model.Leaves, err error) {
	leavesList, err = dao.GetLeavesList()
	if err != nil {
		panic(err)
		return
	}
	return
}
