package db

import (
	"database/sql"
	"seckill/utils"
	"strconv"
	"time"
)

var sliceMap = map[string]map[int][]int{
	"2023-10-6": {
		0: []int{1, 3, 5},
		1: []int{2, 4},
	},
}

type User struct {
	db *sql.DB
}

func selectTableByString(email string) int {
	now := time.Now().Format("2006-01-02")
	result := utils.StringHash(email) % 2
	for k, cur := range sliceMap[now] {
		for _, v := range cur {
			if v == result {
				return k
			}
		}
	}
	return 0
}

func (u *User) UserAdd(email string, password string) (int64, error) {
	id := utils.GenID()
	ret, err := u.db.Exec("insert into userInfo_"+strconv.Itoa(selectTableByString(email))+"values(?,?,?)", id, email, password)
	if err != nil {
		return -1, err
	}
	return ret.RowsAffected()
}

func (u *User) UserLogin(email string, password string) (int64, bool) {
	row := u.db.QueryRow("select id from userInfo_"+strconv.Itoa(selectTableByString(email))+"where email=? and password=?", email, password)
	if row.Err() != nil {
		return 0, false
	}
	var id int64
	row.Scan(&id)
	return id, true
}
