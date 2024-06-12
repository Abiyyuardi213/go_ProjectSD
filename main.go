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
					view.VTransaksiKeluar()

				case 7:
					view.VHistoryTransaksi()

				case 8:
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
