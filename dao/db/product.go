package db

import (
	"database/sql"
	"seckill/dao/db/model"
)

type Product struct {
	db *sql.DB
}

func (p *Product) ProductAdd(product *model.ProductInfo) error {
	_, err := p.db.Exec("insert into product(name,price,pic,des,num) values (?,?,?,?,?)", product.Name, product.Price, product.Pic, product.Des, product.Num)
	return err
}

func (p *Product) ProductList(offset int, limit int) (products []*model.ProductInfo, err error) {
	query, err := p.db.Query("select * from product limit ?,?", (offset-1)*limit, limit)
	if err != nil {
		return nil, err
	}
	var product *model.ProductInfo
	for query.Next() {
		query.Scan(&product.ID, &product.Name, &product.Num, &product.Pic, &product.Des, &product.Price)
		products = append(products, product)
	}
	return
}

func (p *Product) ProductInfo(id int) (product *model.ProductInfo, err error) {
	row := p.db.QueryRow("select * from product where id=?", id)
	if err = row.Err(); err != nil {
		return
	}
	row.Scan(&product)
	return
}

func (p *Product) Try(id, num int) (int64, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}
	row := tx.QueryRow("select num from product where id=?", id)
	if row.Err() != nil {
		return 0, err
	}
	var realNum int
	row.Scan(&realNum)
	if num < realNum {
		tx.Rollback()
		return 0, err
	}
	exec, err := tx.Exec("update product set freezeNum=freezeNum+? where id=?", num, id)
	if err != nil {
		return -1, err
	}
	return exec.RowsAffected()
}

func (p *Product) Commit(id, num int) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("update product set num=num-freezeNum where id=?", id)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("update product set freezeNum =freezeNum-?  where id=?", num, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}

func (p *Product) Cancel(id, num int) error {
	_, err := p.db.Exec("update product set freezeNum=freezeNum-? where id=?", num, id)
	return err
}
