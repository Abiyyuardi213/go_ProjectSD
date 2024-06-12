package view

import (
	"fmt"
	"go_ProjectSD/controller"
	"go_ProjectSD/node"
)

func VAddBarang() {
	var name, stock, price string

	fmt.Println("----- INPUT DATA BARU -----")
	fmt.Print("-- Nama Barang : ")
	fmt.Scan(&name)
	fmt.Print("-- Jumlah Stok : ")
	fmt.Scan(&stock)
	fmt.Print("-- Harga Barang   : ")
	fmt.Scan(&price)

	newData := node.DataBarang{
		Name: name,
		Stock: stock,
		Price: price,
	}

	addedData, err := controller.CAddData(newData)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("Data berhasil ditambahkan.")
	fmt.Printf("-- Serial Number : %d\n", addedData.SerialNumber)
	fmt.Printf("-- Nama Barang   : %s\n", addedData.Name)
	fmt.Printf("-- Jumlah Barang : %s\n", addedData.Stock)
	fmt.Printf("-- Harga Barang  : %s\n", addedData.Price)
	fmt.Println("-------------------------------------------")
}

func VReadAllData() {
	allData, err := controller.CReadAllData()
	if err != nil {
		fmt.Println("Erros: ", err)
		return
	}

	fmt.Println("Daftar semua barang : ")
	for _, data := range allData {
		fmt.Printf("Serial Number : %d\n", data.SerialNumber)
		fmt.Printf("Nama Barang   : %s\n", data.Name)
		fmt.Printf("Jumlah Stok   : %s\n", data.Stock)
		fmt.Printf("Harga Barang  : %s\n", data.Price)
		fmt.Println("-------------------------------------")
	}
}

func VSearchBarang() {
	var serialNumber int

	fmt.Println("----- CARI DATA BARANG -----")
	fmt.Print("-- Serial Number : ")
	fmt.Scan(&serialNumber)

	data, err := controller.CSearchData(serialNumber)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("Data barang ditemukan : ")
	fmt.Printf("-- Serial Number : %d\n", data.SerialNumber)
	fmt.Printf("-- Nama Barang   : %s\n", data.Name)
	fmt.Printf("-- Jumlah Barang : %s\n", data.Stock)
	fmt.Printf("-- Harga Barang  : %s\n", data.Price)
	fmt.Println("-------------------------------------")
}

func VUpdateBarang() {
	var serialNumber int
	var name, stock, price, confirm string

	fmt.Println("----- UPDATE DATA BARANG -----")
	fmt.Print("-- Serial Number : ")
	fmt.Scan(&serialNumber)

	data, err := controller.CSearchData(serialNumber)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("Data barang yang akan diupdate: ")
	fmt.Printf("-- Serial Number : %d\n", data.SerialNumber)
	fmt.Printf("-- Nama Barang   : %s\n", data.Name)
	fmt.Printf("-- Jumlah Stok   : %s\n", data.Stock)
	fmt.Printf("-- Harga Barang  : %s\n", data.Price)
	fmt.Println("----------------------------------------")
	fmt.Print("Apakah anda yakin akan mengupdate data diatas? (y/n): ")
	fmt.Scan(&confirm)

	if confirm == "y" || confirm == "Y" {
		fmt.Print("-- Nama Barang  : ")
		fmt.Scan(&name)
		fmt.Print("-- Jumlah Stok  : ")
		fmt.Scan(&stock)
		fmt.Print("-- Harga Barang : ")
		fmt.Scan(&price)

		updatedData := node.DataBarang{
			SerialNumber: serialNumber,
			Name: name,
			Stock: stock,
			Price: price,
		}

		_, err := controller.CUpdateData(serialNumber, updatedData)
		if err != nil {
			fmt.Println("Error : ", err)
			return
		}

		fmt.Println("Data berhasil diupdate.")
		fmt.Printf("-- Serial Number : %d\n", updatedData.SerialNumber)
		fmt.Printf("-- Nama Barang   : %s\n", updatedData.Name)
		fmt.Printf("-- Jumlah Stok   : %s\n", updatedData.Stock)
		fmt.Printf("-- Harga Barang  : %s\n", updatedData.Price)
		fmt.Println("-------------------------------------------")
	} else {
		fmt.Println("Update data dibatalkan")
	}
}

func VDeleteBarang() {
	var serialNumber int
	var confirmation string
	
	fmt.Println("----- HAPUS DATA BARANG -----")
	fmt.Print("-- Serial Number : ")
	fmt.Scan(&serialNumber)

	data, err := controller.CDeleteData(serialNumber)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("Data barang yang akan di hapus : ")
	fmt.Printf("-- Serial Number : %d\n", data.SerialNumber)
	fmt.Printf("-- Nama Barang   : %s\n", data.Name)
	fmt.Printf("-- Jumlah Stok   : %s\n", data.Stock)
	fmt.Printf("-- Harga Barang  : %s\n", data.Price)
	fmt.Println("-------------------------------------")
	fmt.Print("Apakah anda yakin akan menghapus data diatas? (y/n): ")
	fmt.Scan(&confirmation)

	if confirmation == "y" || confirmation == "Y" {
		deletedData, err := controller.CDeleteData(serialNumber)
		if err != nil {
			fmt.Println("Error : ", err)
			return
		}
		fmt.Println("Data berhasil di hapus.")
		fmt.Printf("-- Serial Number : %d\n", deletedData.SerialNumber)
		fmt.Printf("-- Nama Barang   : %s\n", deletedData.Name)
		fmt.Printf("-- Jumlah Stok   : %s\n", deletedData.Stock)
		fmt.Printf("-- Harga Barang  : %s\n", deletedData.Price)
	} else {
		fmt.Println("Penghapusan data dibatalkan.")
	}
}