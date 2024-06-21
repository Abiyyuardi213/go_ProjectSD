package view

import (
	"fmt"
	"go_ProjectSD/controller"
	"go_ProjectSD/node"
)

func VAddBarang() {
	var name, price string
	var stock int

	fmt.Println("----- INPUT DATA BARU -----")
	fmt.Print("-- Nama Barang : ")
	fmt.Scan(&name)
	fmt.Print("-- Harga Barang   : ")
	fmt.Scan(&price)
	fmt.Print("-- Jumlah Stok : ")
	fmt.Scan(&stock)

	newData := node.DataBarang{
		Name: name,
		Price: price,
		Stock: stock,
	}

	addedData, err := controller.CAddData(newData)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("Data berhasil ditambahkan.")
	fmt.Printf("-- Serial Number : %d\n", addedData.SerialNumber)
	fmt.Printf("-- Nama Barang   : %s\n", addedData.Name)
	fmt.Printf("-- Harga Barang  : %s\n", addedData.Price)
	fmt.Printf("-- Jumlah Barang : %d\n", addedData.Stock)
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
		fmt.Printf("Harga Barang  : %s\n", data.Price)
		fmt.Printf("Jumlah Stok   : %d\n", data.Stock)
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
	fmt.Printf("-- Harga Barang  : %s\n", data.Price)
	fmt.Printf("-- Jumlah Barang : %d\n", data.Stock)
	fmt.Println("-------------------------------------")
}

func VUpdateBarang() {
	var serialNumber, stock int
	var name, price, confirm string

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
	fmt.Printf("-- Harga Barang  : %s\n", data.Price)
	fmt.Printf("-- Jumlah Stok   : %d\n", data.Stock)
	fmt.Println("----------------------------------------")
	fmt.Print("Apakah anda yakin akan mengupdate data diatas? (y/n): ")
	fmt.Scan(&confirm)

	if confirm == "y" || confirm == "Y" {
		fmt.Print("-- Nama Barang  : ")
		fmt.Scan(&name)
		fmt.Print("-- Harga Barang : ")
		fmt.Scan(&price)
		fmt.Print("-- Jumlah Stok  : ")
		fmt.Scan(&stock)

		updatedData := node.DataBarang{
			SerialNumber: serialNumber,
			Name: name,
			Price: price,
			Stock: stock,
		}

		_, err := controller.CUpdateData(serialNumber, updatedData)
		if err != nil {
			fmt.Println("Error : ", err)
			return
		}

		fmt.Println("Data berhasil diupdate.")
		fmt.Printf("-- Serial Number : %d\n", updatedData.SerialNumber)
		fmt.Printf("-- Nama Barang   : %s\n", updatedData.Name)
		fmt.Printf("-- Harga Barang  : %s\n", updatedData.Price)
		fmt.Printf("-- Jumlah Stok   : %d\n", updatedData.Stock)
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
	fmt.Printf("-- Harga Barang  : %s\n", data.Price)
	fmt.Printf("-- Jumlah Stok   : %d\n", data.Stock)
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
		fmt.Printf("-- Harga Barang  : %s\n", deletedData.Price)
		fmt.Printf("-- Jumlah Stok   : %d\n", deletedData.Stock)
	} else {
		fmt.Println("Penghapusan data dibatalkan.")
	}
}