package db

import (
	"database/sql"
	"seckill/utils"
	"strconv"
	"time"
)

type Order struct {
	db *sql.DB
}

func selectTableByInt64(id int64) int {
	now := time.Now().Format("2006-01-02")
	result := int(id % 2)
	for k, cur := range sliceMap[now] {
		for _, v := range cur {
			if v == result {
				return k
			}
		}
	}
	return 0
}

func (o *Order) OrderAdd(uid int64, pid int, timeStamp int64) (int64, error) {
	id := utils.GenID()
	exec, err := o.db.Exec("insert into user_order_"+strconv.Itoa(selectTableByInt64(id))+"(id,uid,pid,created_at) values(?,?,?,?)", id, uid, pid, timeStamp)
	if err != nil {
		return -1, err
	}
	return exec.RowsAffected()
}

func (o *Order) OrderStatusAdd(id int64, uid int64, pid int) error {
	_, err := o.db.Exec("insert into orderStatus values (?,?,?,?)", id, uid, pid, "Created")
	return err
}

func (o *Order) Try(id int64, uid int64, pid int) (int64, error) {
	var status string
	tx, err := o.db.Begin()
	if err != nil {
		return -1, err
	}
	row := tx.QueryRow("select status from orderStatus where id=? and uid=? and pid=?", id, uid, pid)
	if row.Err() != nil {
		return -1, err
	}
	row.Scan(&status)
	if status != "Created" {
		return 0, nil
	}
	exec, err := tx.Exec("update orderStatus set status=? where id=? and uid=? and pid=?", "Doing", id, uid, pid)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	if affect, err := exec.RowsAffected(); err != nil || affect == 0 {
		tx.Rollback()
		return affect, err
	}
	tx.Commit()
	return 1, nil
}

func (o *Order) Commit(id int64, uid int64, pid int) (int64, error) {
	var status string
	tx, err := o.db.Begin()
	if err != nil {
		return -1, err
	}
	row := tx.QueryRow("select status from orderStatus where id=? and uid=? and pid=?", id, uid, pid)
	if row.Err() != nil {
		return -1, err
	}
	row.Scan(&status)
	if status != "Doing" {
		return 0, nil
	}
	exec, err := tx.Exec("update orderStatus set status=? where id=? and uid=? and pid=?", "Finished", id, uid, pid)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	if affect, err := exec.RowsAffected(); err != nil || affect == 0 {
		tx.Rollback()
		return affect, err
	}
	tx.Commit()
	return 1, nil
}

func (o *Order) Cancel(id int64, uid int64, pid int) (int64, error) {
	var status string
	tx, err := o.db.Begin()
	if err != nil {
		return -1, err
	}
	row := tx.QueryRow("select status from orderStatus where id=? and uid=? and pid=?", id, uid, pid)
	if row.Err() != nil {
		return -1, err
	}
	row.Scan(&status)
	if status != "Doing" {
		return 0, nil
	}
	exec, err := tx.Exec("update orderStatus set status=? where id=? and uid=? and pid=?", "Created", id, uid, pid)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	if affect, err := exec.RowsAffected(); err != nil || affect == 0 {
		tx.Rollback()
		return affect, err
	}
	tx.Commit()
	return 1, nil
}
