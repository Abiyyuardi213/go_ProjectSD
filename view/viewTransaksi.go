package view

import (
	"fmt"
	"go_ProjectSD/controller"
	"go_ProjectSD/node"
	"time"
)

func VCreateTransaksi(adminID int, username string) {
	var barangCount int
	var barangs []node.BarangTransaksi

	fmt.Println("----- TRANSAKSI BARANG -----")
	fmt.Print("Jumlah jenis barang : ")
	fmt.Scan(&barangCount)

	for i := 0; i < barangCount; i++ {
		var serialNumber, quantity int
		fmt.Printf("Serial Number barang ke-%d: ", i+1)
		fmt.Scan(&serialNumber)
		fmt.Printf("Jumlah item barang ke-%d  : ", i+1)
		fmt.Scan(&quantity)

		barangs = append(barangs, node.BarangTransaksi{
			SerialNumber: serialNumber,
			Quantity: quantity,
		})
	}

	waktuTransaksi := time.Now()
	fmt.Println("----- DETAIL TRANSAKSI -----")
	fmt.Printf("-- Admin ID : %d\n", adminID)
	fmt.Printf("-- Username : %s\n", username)
	fmt.Printf("-- Waktu    : %s\n", waktuTransaksi.Format("2006-01-02 15:04:05"))
	fmt.Println("-- Keterangan Barang : ")
	for i, barang := range barangs {
		fmt.Printf(" %d. Serial Number : %d\n", i+1, barang.SerialNumber)
		fmt.Printf("    Quanity       : %d\n", barang.Quantity)
	}

	var confirm string
	fmt.Print("Apakah data diatas sudah benar? (y/n) : ")
	fmt.Scan(&confirm)

	if confirm != "y" && confirm != "Y" {
		fmt.Println("Transaksi dibatalkan.")
		return
	}

	err := controller.CCreateTransaksi(adminID, username, barangs)
	if err != nil {
		fmt.Println("Gagal membuat transaksi : ", err)
		return
	}

	fmt.Println("----- TRANSAKSI BERHASIL -----")
	fmt.Println("Detail Transaksi : ")
	fmt.Printf("-- Admin ID : %d\n", adminID)
	fmt.Printf("-- Username : %s\n", username)
	fmt.Printf("-- Waktu    : %s\n", waktuTransaksi.Format("2006-01-02 15:04:05"))
	fmt.Println("-- Jenis Barang : ")
	for i, barang := range barangs {
		fmt.Printf(" %d. Serial Number : %d\n", i+1, barang.SerialNumber)
		fmt.Printf("    Jumlah Barang : %d\n", barang.Quantity)
	}
}

func VHistoryTransaksi() {
	transactions := controller.CGetTransactionHistory()

	fmt.Println("----- HISTORY TRANSAKSI -----")
	for _, transaksi := range transactions {
		fmt.Printf("-- ID Transaksi   : %d\n", transaksi.ID)
		fmt.Printf("-- Admin ID       : %d\n", transaksi.AdminID)
		fmt.Printf("-- Username       : %s\n", transaksi.Username)
		fmt.Println("-- Barang         : ")
		for _, barang := range transaksi.Barang {
			fmt.Printf("  - Serial Number : %d\n", barang.SerialNumber)
			fmt.Printf("    Quantity      : %d\n", barang.Quantity)
		}
		fmt.Printf("-- Waktu     : %s\n", transaksi.TransaksiTime.Format("2006-01-02 15:04:05"))
		fmt.Println("---------------------------------------------")
	}
	fmt.Println("==========================================")
}