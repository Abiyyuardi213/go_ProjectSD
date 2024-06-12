package controller

import (
	"fmt"
	"go_ProjectSD/model"
	"go_ProjectSD/node"
)

func CTransaksiKeluar(barangName string, quantity int) (node.DataBarang, error) {
	updatedData, err := model.MTransaksiKeluar(barangName, quantity)
	if err != nil {
		return node.DataBarang{}, fmt.Errorf("failed to process transaction: %v", err)
	}
	return updatedData, nil
}

func CHistoryTransaksi() []node.TransaksiKeluar {
	return model.MGetHistoryTransaksi()
}