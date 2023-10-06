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

func (p *Product) Try(id, num int) (int64, error) {
	exec, err := p.db.Exec("update product set freezeNum=freezeNum+? where id=?", num, id)
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
