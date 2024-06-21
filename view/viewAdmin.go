package view

import (
	"bufio"
	"fmt"
	"go_ProjectSD/controller"
	"go_ProjectSD/node"
	"os"
	"strings"
)

func VRegisterAdmin() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Register Admin")

	fmt.Print("Input Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Input Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Input Phone: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	fmt.Print("Input Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Input Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	data := node.DataAdmin{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Username: username,
		Password: password,
	}

	controller.CRegisterAdmin(data)
}

func VLoginAdmin() (int, string) {
	var username, password string
	fmt.Print("Username : ")
	fmt.Scan(&username)
	fmt.Print("Password : ")
	fmt.Scan(&password)

	adminID, err := controller.CLoginAdmin(username, password)
	if err != nil {
		fmt.Println("Login gagal:", err)
		return 0, ""
	}
	fmt.Println("Login Berhasil")
	return adminID, username
}

func VDaftarAdmin() {
	admins := controller.CDaftarAdmin()
	if len(admins) == 0 {
		fmt.Println("No admins founds!")
		return
	}
	fmt.Println("List of admins : ")
	fmt.Println("------------------------")
	for _, admin := range admins {
		fmt.Printf("ID Admin       : %d\n", admin.Id)
		fmt.Printf("Nama Admin     : %s\n", admin.Name)
		fmt.Printf("Email Admin    : %s\n", admin.Email)
		fmt.Printf("Telepon Admin  : %s\n", admin.Phone)
		fmt.Printf("Username Admin : %s\n", admin.Username)
		fmt.Println("------------------------")
	}
}

func VLoginHistory() {
	history := controller.CLoginHistory()
	if len(history) == 0 {
		fmt.Println("No login history found!")
		return
	}
	fmt.Println("----- ADMIN LOGIN HISTORY -----")
	fmt.Println("===============================")
	for _, record := range history {
		fmt.Printf("Admin ID   : %d\n", record.AdminID)
		fmt.Printf("Username   : %s\n", record.Username)
		fmt.Printf("Login Time : %s\n", record.LoginTime.Format("2006-01-02 15:04:05"))
		fmt.Println("===============================")
	}
}