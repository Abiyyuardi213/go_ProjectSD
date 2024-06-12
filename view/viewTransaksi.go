package view

import (
	"fmt"
	"go_ProjectSD/controller"
	"go_ProjectSD/database"
)

func VTransaksiKeluar() {
	var barangName string
	var quantity int

	fmt.Println("----- TRANSAKSI BARANG KELUAR -----")
	fmt.Print("-- Nama Barang   : ")
	fmt.Scan(&barangName)
	fmt.Print("-- Jumlah Barang : ")
	fmt.Scan(&quantity)

	updatedData, err := controller.CTransaksiKeluar(barangName, quantity)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("Transaksi barang keluar berhasil.")
	fmt.Printf("-- Nama Barang   : %s\n", updatedData.Name)
	fmt.Printf("-- Jumlah Barang : %s\n", updatedData.Stock)
	fmt.Printf("-- Admin         : %s\n", database.LoggedInAdmin.Username)
	fmt.Println("----------------------------------------")
}

func VHistoryTransaksi() {
	history := controller.CHistoryTransaksi()
	if len(history) == 0 {
		fmt.Println("No transaction history found!")
		return
	}
	fmt.Println("----- TRANSACTION HISTORY -----")
	fmt.Println("===============================")
	for _, record := range history {
		fmt.Printf("-- Admin ID : %d\n", record.AdminID)
		fmt.Printf("-- Username : %s\n", record.Username)
		fmt.Printf("-- Nama Barang : %s\n", record.NamaBarang)
		fmt.Printf("-- Jumlah Barang : %d\n", record.Quantity)
		fmt.Printf("-- Waktu         : %s\n", record.Time.Format("2006-01-02 15:04:05"))
		fmt.Println("===============================")
	}
}