package controller

import (
	"go_ProjectSD/model"
	"go_ProjectSD/node"
)

func CCreateTransaksi(adminID int, username string, barangs []node.BarangTransaksi) error {
	return model.MCreateTransaction(adminID, username, barangs)
}

func CGetTransactionHistory() []node.Transaksi {
	return model.MGetTransactionHistory()
}