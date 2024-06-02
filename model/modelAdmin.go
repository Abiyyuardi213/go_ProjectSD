package model

import (
	"errors"
	"go_ProjectSD/database"
	"go_ProjectSD/node"
	"math/rand"
	"time"
)

func MRegisterAdmin(data node.DataAdmin) (node.DataAdmin, error) {
	rand.Seed(time.Now().UnixNano())

	var id int
	exists := true

	//loop untuk memastikan ID yang dihasilkan bernilai unik
	for exists {
		id = rand.Intn(900) + 100 //id acak antara 100 dan 999
		exists = MCheckIDExists(id)
	}

	data.Id = id

	newNode := &node.AdminLL{
		DBAdmin: data,
		Next: nil,
	}

	if database.DatabaseAdmin.Next == nil {
		database.DatabaseAdmin.Next = newNode
	} else {
		current := database.DatabaseAdmin.Next
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}

	return data, nil
}

func MCheckIDExists(id int) bool {
	current := database.DatabaseAdmin.Next
	for current != nil {
		if current.DBAdmin.Id == id {
			return true
		}
		current = current.Next
	}
	return false
}

//fungsi untuk login admin
func MLoginAdmin(username, password string) (node.DataAdmin, error) {
	current := database.DatabaseAdmin.Next
	for current != nil {
		if current.DBAdmin.Username == username && current.DBAdmin.Password == password {
			SaveLoginHistory(current.DBAdmin.Id, current.DBAdmin.Username)
			return current.DBAdmin, nil
		}
		current = current.Next
	}

	return node.DataAdmin{}, errors.New("invalid username or password")
}

func MDaftarAdmin() []node.DataAdmin {
	var admins []node.DataAdmin
	current := database.DatabaseAdmin.Next
	for current != nil {
		admins = append(admins, current.DBAdmin)
		current = current.Next
	}
	return admins
}

func SaveLoginHistory(adminID int, username string) {
	loginTime := time.Now()
	newHistory := &node.LoginHistory{
		AdminID: adminID,
		Username: username,
		LoginTime: loginTime,
		Next: nil,
	}

	if database.AdminLoginHistory.Next == nil {
		database.AdminLoginHistory.Next = newHistory
	} else {
		current := database.AdminLoginHistory
		for current.Next != nil {
			current = *current.Next
		}
		current.Next = newHistory
	}
}

func MGetLoginHistory() []node.LoginHistory {
	var history []node.LoginHistory
	current := database.AdminLoginHistory.Next
	for current != nil {
		history = append(history, *current)
		current = current.Next
	}
	return history
}