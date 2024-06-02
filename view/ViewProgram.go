package view

import "fmt"

func FirstMenu() {
	fmt.Println("PT. Indomarco Jaya")
	fmt.Println("1. Login Admin")
	fmt.Println("2. Register Admin")
	fmt.Println("3. List Admin")
	fmt.Println("4. Admin Login History")
	fmt.Println("5. Exit")
}

func MenuAdmin() {
	fmt.Println("----- WAREHOUSE MANAGEMENT -----")
	fmt.Println("--------------------------------")
	fmt.Println("1. Input Data")
	fmt.Println("2. Update Data")
	fmt.Println("3. Delete Data")
	fmt.Println("4. Search Data")
	fmt.Println("5. Data List")
	fmt.Println("6. Transaction of Outgoing Goods")
	fmt.Println("7. Transaction of Ingoing Goods")
	fmt.Println("8. Back to Home Page")
	fmt.Println("--------------------------------")
}