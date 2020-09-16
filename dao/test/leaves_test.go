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

func TestInsertLeaves(t *testing.T) {
	leaves := &model.Leaves{
		Username:   "yuanjg",
		Email:      "yuanjg@163.com",
		Content:    "写的超级棒，nice，学习了",
		CreateTime: time.Now(),
	}
	leaveId, err := dao.InsertLeaves(leaves)
	if err != nil {
		return
	}
	t.Logf("LeaveId: %d\n", leaveId)
}

func TestGetLeaveList(t *testing.T) {
	leavesList, err := dao.GetLeavesList()
	if err != nil {
		return
	}
	for _, v := range leavesList {
		t.Logf("leaves: %#v", v)
	}
}
