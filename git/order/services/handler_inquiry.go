package services

import (
	"context"
	"database/sql"
	"fmt"

	cm "pnp/Framework/git/order/common"

	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) InquiryHandler(ctx context.Context, req cm.InquiryRequest) (res cm.InquiryResponse) {

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

	var inquiryResponse cm.InquiryResponse
	var list cm.PaymentChannel

	sql := `SELECT
				trx_id,
				IFNULL(payment_status_code,'')
			FROM transaksi WHERE merchant_id = ?`

	result, err := db.Query(sql, req.MerchantID)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err := result.Scan(&list.PgCode, &list.PgName)

		if err != nil {
			panic(err.Error())
		}

		inquiryResponse.PaymentChannel = append(inquiryResponse.PaymentChannel, list)

	}

	inquiryResponse.Merchant = req.Merchant
	inquiryResponse.MerchantID = req.MerchantID

	res = inquiryResponse

	return
}
