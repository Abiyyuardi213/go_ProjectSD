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
	var loggedInAdminID int
	var loggedInAdminUsername string

	reader := bufio.NewReader(os.Stdin) // Ini untuk membersihkan buffer setelah fmt.Scan

	for {
		view.FirstMenu()
		fmt.Print("Input your choice: ")
		fmt.Scan(&pilih)

		// Membersihkan buffer setelah fmt.Scan
		reader.ReadString('\n')

		switch pilih {
		case 1:
			loggedInAdminID, loggedInAdminUsername = view.VLoginAdmin()

			adminLoop:
			for {
				view.MenuAdmin()
				fmt.Print("Input your choice: ")
				fmt.Scan(&pilihMenu)

				reader.ReadString('\n')

				switch pilihMenu {
				case 1:
					view.VAddBarang()

				case 2:
					view.VUpdateBarang()

				case 3:
					view.VDeleteBarang()

				case 4:
					view.VSearchBarang()

				case 5:
					view.VReadAllData()

				case 6:
					if loggedInAdminID != 0 && loggedInAdminUsername != "" {
						view.VCreateTransaksi(loggedInAdminID, loggedInAdminUsername)
					} else {
						fmt.Println("Admin belum login.")
					}

				case 7:
					view.VHistoryTransaksi()

				case 8:
					fmt.Println("Returning to main menu...")
					break adminLoop
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
