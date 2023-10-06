package db

import (
	"database/sql"
	"seckill/conf"
)

type Dao struct {
	User    *User
	Product *Product
	Order   *Order
}

func NewDao() Dao {
	var err error
	userDb, err := sql.Open("mysql", conf.MysqlUserDSN)
	if err != nil {
		panic(err)
	}
	orderDb, err := sql.Open("mysql", conf.MysqlOrderStatusDSN)
	if err != nil {
		panic(err)
	}
	productDb, err := sql.Open("mysql", conf.MysqlProductDSN)
	if err != nil {
		panic(err)
	}
	return Dao{
		User: &User{
			userDb,
		},
		Order: &Order{
			orderDb,
		},
		Product: &Product{
			productDb,
		},
	}
}
