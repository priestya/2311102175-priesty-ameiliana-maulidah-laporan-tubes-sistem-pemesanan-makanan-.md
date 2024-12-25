package main

import (
	"fmt"
)

// Struct Menu untuk menyimpan informasi makanan/minuman
type Menu struct {
	Nama  string
	Harga float64
}

// Struct Pesanan untuk menyimpan pesanan pelanggan
type Pesanan struct {
	Menu   Menu
	Jumlah int
}

// Fungsi untuk menampilkan menu restoran
func tampilkanMenu(menu []Menu) {
	fmt.Println("\n==== MENU RESTORAN ====")
	for i, item := range menu {
		fmt.Printf("%d. %s - Rp %.2f\n", i+1, item.Nama, item.Harga)
	}
	fmt.Println("0. Selesai dan Hitung Total")
}

// Fungsi untuk menghitung total harga pesanan
func hitungTotal(pesanan []Pesanan) float64 {
	total := 0.0
	for _, p := range pesanan {
		total += p.Menu.Harga * float64(p.Jumlah)
	}
	return total
}

func main() {
	// Daftar menu restoran
	menu := []Menu{
		{"Nasi Goreng", 20000},
		{"Mie Goreng", 18000},
		{"Ayam Bakar", 25000},
		{"Es Teh", 5000},
		{"Es Jeruk", 7000},
	}

	var pesanan []Pesanan
	var selesai bool

	for !selesai {
		// Tampilkan daftar menu
		tampilkanMenu(menu)

		// Input pilihan menu
		var pilihan, jumlah int
		fmt.Print("Pilih menu (nomor): ")
		fmt.Scan(&pilihan)

		if pilihan == 0 {
			selesai = true // Keluar dari loop
		} else if pilihan > 0 && pilihan <= len(menu) {
			// Input jumlah pesanan
			fmt.Print("Jumlah: ")
			fmt.Scan(&jumlah)

			// Menambahkan pesanan ke daftar
			pesanan = append(pesanan, Pesanan{menu[pilihan-1], jumlah})
			fmt.Printf("Pesanan ditambahkan: %s x%d\n", menu[pilihan-1].Nama, jumlah)
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}

	// Tampilkan rincian pesanan dan total harga
	fmt.Println("\n===== RINCIAN PESANAN =====")
	totalHarga := hitungTotal(pesanan)
	for _, p := range pesanan {
		fmt.Printf("%s x%d = Rp %.2f\n", p.Menu.Nama, p.Jumlah, p.Menu.Harga*float64(p.Jumlah))
	}
	fmt.Printf("Total Harga: Rp %.2f\n", totalHarga)
	fmt.Println("Terima kasih telah memesan di restoran kami!")
}
