package model

import (
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