package controller

import (
	"fmt"
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

func CLoginAdmin(username, password string) {
	adminData, err := model.MLoginAdmin(username, password)
	if err != nil {
		fmt.Println("Login Failed : ", err.Error())
		return
	}
	fmt.Println("Login successfull!")
	fmt.Println("Welcome,", adminData.Name)
}

func CDaftarAdmin() []node.DataAdmin {
	return model.MDaftarAdmin()
}

func CLoginHistory() []node.LoginHistory {
	return model.MGetLoginHistory()
}