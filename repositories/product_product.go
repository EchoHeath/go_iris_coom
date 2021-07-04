package repositories

import (
	"database/sql"
	"go-electricity/datamodels"
)

type IProduct interface {
	Conn() error
	Insert(product *datamodels.Product) (int64, error)
	Delete(id int64) bool
	Update(product *datamodels.Product) error
	SelectByKey(id int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

type ProductManager struct {
	table     string
	mysqlConn *sql.DB
}

func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{table: table, mysqlConn: db}
}

func (p *ProductManager) Conn() error {

}