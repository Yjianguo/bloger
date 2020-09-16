package dao

import (
	"code.lssj.me/bloger/dao/db"
	"code.lssj.me/bloger/model"
	"log"
	"time"
)

// 写留言
func InsertLeaves(leaves *model.Leaves) (leaveId int64, err error) {
	sqlStr := `insert into leaves(username, email, content, create_time) values (?,?,?,?)`
	result, err := db.DB.Exec(sqlStr, leaves.Username, leaves.Email, leaves.Content, time.Now())
	if err != nil {
		log.Fatal(err)
		return
	}
	leaveId, err = result.LastInsertId()
	return
}

//查询留言
func GetLeavesList() (leaveList []*model.Leaves, err error) {
	sqlStr := `select username, email, content, create_time from leaves order by create_time desc limit 10`
	err = db.DB.Select(&leaveList, sqlStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
