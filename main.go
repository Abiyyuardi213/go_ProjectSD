package main

import (
	"bufio"
	"fmt"
	"go_ProjectSD/view"
	"os"
)

func MainProgram() {
	var pilih int
	var pilihMenu int
	reader := bufio.NewReader(os.Stdin) // Ini untuk membersihkan buffer setelah fmt.Scan

	for {
		view.FirstMenu()
		fmt.Print("Input your choice: ")
		fmt.Scan(&pilih)

		// Membersihkan buffer setelah fmt.Scan
		reader.ReadString('\n')

		switch pilih {
		case 1:
			view.VLoginAdmin()
			
			for {
				view.MenuAdmin()
				fmt.Print("Input your choice: ")
				fmt.Scan(&pilihMenu)

				switch pilihMenu {
				case 1:
					fmt.Println("your choice is input data")

				case 2:
					fmt.Println("your choice is update data")

				case 3:
					fmt.Println("your choice is delete data")

				case 4:
					fmt.Println("your choice is search data")

				case 5:
					fmt.Println("your choice is list data")

				case 6:
					return
				}
			}

		case 2:
			view.VRegisterAdmin()

		case 3:
			view.VDaftarAdmin()

		case 4:
			view.VLoginHistory()

		case 5:
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func main() {
	MainProgram()
}