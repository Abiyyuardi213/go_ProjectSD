package model

import (
	"fmt"
	"go_ProjectSD/database"
	"go_ProjectSD/node"
	"math/rand"
	"time"
)

func MAddData(data node.DataBarang) (node.DataBarang, error) {
	rand.Seed(time.Now().UnixNano())

	var serialNumber int
	exist := true

	for exist {
		serialNumber = rand.Intn(9999) + 1000
		exist = MCheckSerialNumber(serialNumber)
	}

	data.SerialNumber = serialNumber

	newNode := &node.BarangLL{
		DBBarang: data,
		Next: nil,
	}

	if database.DatabaseBarang.Next == nil {
		database.DatabaseBarang.Next = newNode
	} else {
		current := database.DatabaseBarang.Next
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}

	return data, nil
}

func MCheckSerialNumber(serialNumber int) bool {
	current := database.DatabaseBarang.Next
	for current != nil {
		if current.DBBarang.SerialNumber == serialNumber {
			return true
		}
		current = current.Next
	}
	return false
}

func MReadAllData() []node.DataBarang {
	var allData []node.DataBarang
	current := database.DatabaseBarang.Next

	for current != nil {
		allData = append(allData, current.DBBarang)
		current = current.Next
	}

	return allData
}

func MUpdateData(serialNumber int, updatedData node.DataBarang) (node.DataBarang, error) {
	current := database.DatabaseBarang.Next

	for current != nil {
		if current.DBBarang.SerialNumber == serialNumber {
			current.DBBarang.Name = updatedData.Name
			current.DBBarang.Stock = updatedData.Stock
			current.DBBarang.Price = updatedData.Price
			return current.DBBarang, nil
		}
		current = current.Next
	}
	return node.DataBarang{}, fmt.Errorf("data dengan serial number %d tidak ditemukan", serialNumber)
}

func MDeleteData(serialNumber int) (node.DataBarang, error) {
	if database.DatabaseBarang.Next == nil {
		return node.DataBarang{}, fmt.Errorf("database kosong")
	}

	var deletedData node.DataBarang
	if database.DatabaseBarang.Next.DBBarang.SerialNumber == serialNumber {
		deletedData = database.DatabaseBarang.Next.DBBarang
		database.DatabaseBarang.Next = database.DatabaseBarang.Next.Next
		return deletedData, nil
	}

	current := database.DatabaseBarang.Next
	for current.Next != nil {
		if current.Next.DBBarang.SerialNumber == serialNumber {
			deletedData = current.Next.DBBarang
			current.Next = current.Next.Next
			return deletedData, nil
		}
		current = current.Next
	}

	return node.DataBarang{}, fmt.Errorf("data dengan serial number %d tidak ditemukan", serialNumber)
}