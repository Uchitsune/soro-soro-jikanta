package main

import "fmt"

func main() {
	var hari int

	fmt.Print("Masukkan angka hari (1-11): ")
	fmt.Scan(&hari)

	switch hari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	case 7:
		fmt.Println("Minggu")
	case 8:
		fmt.Println("Bayar Pajak")
	case 9:
		fmt.Println("Bayar Utang")
	case 10:
		fmt.Println("hari maret :)")
	case 11:
		fmt.Println("Valentine")
	case 12:
		fmt.Println("Kalkulus Putaaa")
	default:
		fmt.Println("Angka tidak valid!")
	}
}
