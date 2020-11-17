package services

import (
	"context"
	"database/sql"
	"fmt"

	cm "pnp/Framework/git/order/common"

	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) ProductHandler(ctx context.Context, req cm.Products) (res cm.Products) {
	var db *sql.DB
	var err error
	defer panicRecovery()

	host := cm.Config.Connection.Host
	port := cm.Config.Connection.Port
	user := cm.Config.Connection.User
	pass := cm.Config.Connection.Password
	data := cm.Config.Connection.Database

	var mySQL = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, data)

	db, err = sql.Open("mysql", mySQL)

	if err != nil {
		panic(err.Error())
	}

	res.ProductName = req.ProductName

	var product cm.Products
	var products []cm.Products

	sql := `SELECT
				ProductID,
				ProductName,
				IFNULL(SupplierID,'') SupplierID,
				IFNULL(CategoryID,'') CategoryID,
				IFNULL(QuantityPerUnit,'') QuantityPerUnit,
				IFNULL(UnitPrice,'') UnitPrice,
				IFNULL(UnitsInStock,'') UnitInStock,
				IFNULL(UnitsOnOrder,'') UnitOnOrder ,
				IFNULL(ReorderLevel,'') ReorderLevel,
				IFNULL(Discontinued,'') Discontinued ,
				IFNULL(Description,'') Description

			
				FROM products WHERE ProductName = ?`

	result, err := db.Query(sql, req.ProductName)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&product.ProductID, &product.ProductName, &product.SupplierID, &product.CategoryID, &product.QuantityPerUnit, &product.UnitPrice, &product.UnitsInStock, &product.UnitsOnOrder, &product.ReorderLevel, &product.Discontinued, &product.Description)

		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)

	}

	res = product

	return
}
