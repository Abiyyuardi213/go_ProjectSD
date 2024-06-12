package model

import (
	"fmt"
	"go_ProjectSD/database"
	"go_ProjectSD/node"
	"strconv"
	"time"
)

func MTransaksiKeluar(barangName string, quantity int) (node.DataBarang, error) {
	current := database.DatabaseBarang.Next
	var updatedBarang node.DataBarang

	for current != nil {
		if current.DBBarang.Name == barangName {
			stock, _ := strconv.Atoi(current.DBBarang.Stock)
			if stock < quantity {
				return node.DataBarang{}, fmt.Errorf("stok barang tidak mencukupi")
			}
			stock -= quantity
			current.DBBarang.Stock = strconv.Itoa(stock)
			updatedBarang = current.DBBarang

			newTransaksi := &node.TransaksiKeluar{
				AdminID: database.LoggedInAdmin.Id,
				Username: database.LoggedInAdmin.Username,
				NamaBarang: barangName,
				Quantity: quantity,
				Time: time.Now(),
				Next: nil,
			}

			if database.DatabaseTransaksiOut.Next == nil {
				database.DatabaseTransaksiOut.Next = newTransaksi
			} else {
				transaksiCurrent := database.DatabaseTransaksiOut.Next
				for transaksiCurrent.Next != nil {
					transaksiCurrent = transaksiCurrent.Next
				}
				transaksiCurrent.Next = newTransaksi
			}
			return updatedBarang, nil
		}
		current = current.Next
	}

	return node.DataBarang{}, fmt.Errorf("barang dengan nama %s tidak ditemukan", barangName)
}

func MGetHistoryTransaksi() []node.TransaksiKeluar {
	var history []node.TransaksiKeluar
	current := database.DatabaseTransaksiOut.Next
	for current != nil {
		history = append(history, *current)
		current = current.Next
	}
	return history
}