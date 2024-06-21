package model

import (
	"errors"
	"go_ProjectSD/database"
	"go_ProjectSD/node"
	"math/rand"
	"time"
)

func MCheckTransactionIDExist(transactionID int) bool {
	current := database.DatabaseTransaksi.Next
	for current != nil {
		if current.ID == transactionID {
			return true
		}
		current = current.Next
	}
	return false
}

func MCreateTransaction(adminID int, username string, barangs []node.BarangTransaksi) error {
	rand.Seed(time.Now().UnixNano())
	var transactionID int
	exist := true

	for exist {
		transactionID = rand.Intn(9000) + 1000
		exist = MCheckTransactionIDExist(transactionID)
	}

	transactionTime := time.Now()
	newTransaction := &node.Transaksi{
		ID: transactionID,
		AdminID: adminID,
		Username: username,
		Barang: barangs,
		TransaksiTime: transactionTime,
		Next: nil,
	}

	if database.DatabaseTransaksi.Next == nil {
		database.DatabaseTransaksi.Next = newTransaction
	} else {
		current := database.DatabaseTransaksi
		for current.Next != nil {
			current = *current.Next
		}
		current.Next = newTransaction
	}

	for _, barangTransaksi := range barangs {
		err := MUpdatedBarangQuantity(barangTransaksi.SerialNumber, barangTransaksi.Quantity)
		if err != nil {
			return nil
		}
	}
	return nil
}

func MUpdatedBarangQuantity(serialNumber int, barangQuantity int) error {
	current := database.DatabaseBarang.Next
	for current != nil {
		if current.DBBarang.SerialNumber == serialNumber {
			if current.DBBarang.Stock >= barangQuantity {
				current.DBBarang.Stock -= barangQuantity
				return nil
			} else {
				return errors.New("stok tidak mencukupi untuk transaksi ini")
			}
		}
		current = current.Next
	}
	return errors.New("barang tidak ditemukan")
}

func MGetTransactionHistory() []node.Transaksi {
	var transactions []node.Transaksi
	current := database.DatabaseTransaksi.Next
	for current != nil {
		transactions = append(transactions, *current)
		current = current.Next
	}
	return transactions
}