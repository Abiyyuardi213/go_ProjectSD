package controller

import (
	"errors"
	"fmt"
	"go_ProjectSD/database"
	"go_ProjectSD/model"
	"go_ProjectSD/node"
)

func CRegisterAdmin(data node.DataAdmin) {
	registeredData, err := model.MRegisterAdmin(data)
	if err != nil {
		fmt.Println("Failed to register Admin!", err)
		return
	}
	fmt.Println("Admin registered successfully!")
	fmt.Println("Admin ID : ", registeredData.Id)
}

func CLoginAdmin(username, password string) (int, error) {
	current := database.DatabaseAdmin.Next
	for current != nil {
		if current.DBAdmin.Username == username && current.DBAdmin.Password == password {
			return current.DBAdmin.Id, nil
		}
		current = current.Next
	}
	return 0, errors.New("username atau password salah")
}

func CDaftarAdmin() []node.DataAdmin {
	return model.MDaftarAdmin()
}

func CLoginHistory() []node.LoginHistory {
	return model.MGetLoginHistory()
}