package controller

import (
	"fmt"
	"go_ProjectSD/model"
	"go_ProjectSD/node"
)

func CAddData(data node.DataBarang) (node.DataBarang, error) {
	newData, err := model.MAddData(data)
	if err != nil {
		return newData, fmt.Errorf("failed to add data: %v", err)
	}

	return newData, nil
}

func CReadAllData() ([]node.DataBarang, error) {
	allData := model.MReadAllData()

	if len(allData) == 0 {
		return nil, fmt.Errorf("tidak ada barang dalam database")
	}

	return allData, nil
}

func CSearchData(serialNumber int) (node.DataBarang, error) {
	data, err := model.MSearchData(serialNumber)
	if err != nil {
		return node.DataBarang{}, fmt.Errorf("failed to search data: %v", err)
	}
	return data, nil
}

func CUpdateData(serialNumber int, updatedData node.DataBarang) (node.DataBarang, error) {
	updatedData.SerialNumber = serialNumber
	updatedData, err := model.MUpdateData(serialNumber, updatedData)
	if err != nil {
		return node.DataBarang{}, fmt.Errorf("failed to update data: %v", err)
	}
	return updatedData, nil
}

func CDeleteData(serialNumber int) (node.DataBarang, error) {
	deletedData, err := model.MDeleteData(serialNumber)
	if err != nil {
		return node.DataBarang{}, fmt.Errorf("failed to delete data : %v", err)
	}

	return deletedData, nil
}