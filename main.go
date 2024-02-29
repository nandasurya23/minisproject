package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Buku struct {
	Kode         string
	Judul        string
	Pengarang    string
	Penerbit     string
	JumlahHalaman int
	TahunTerbit  int
}

var DaftarBuku []Buku

func TambahBuku() {
	input := bufio.NewReader(os.Stdin)

	var bukuBaru Buku

	fmt.Println("Tambah Buku Baru")
	fmt.Print("Kode Buku: ")
	bukuBaru.Kode, _ = input.ReadString('\n')
	bukuBaru.Kode = strings.TrimSpace(bukuBaru.Kode)

	fmt.Print("Judul Buku: ")
	bukuBaru.Judul, _ = input.ReadString('\n')
	bukuBaru.Judul = strings.TrimSpace(bukuBaru.Judul)

	fmt.Print("Pengarang: ")
	bukuBaru.Pengarang, _ = input.ReadString('\n')
	bukuBaru.Pengarang = strings.TrimSpace(bukuBaru.Pengarang)

	fmt.Print("Penerbit: ")
	bukuBaru.Penerbit, _ = input.ReadString('\n')
	bukuBaru.Penerbit = strings.TrimSpace(bukuBaru.Penerbit)

	fmt.Print("Jumlah Halaman: ")
	fmt.Scanln(&bukuBaru.JumlahHalaman)

	fmt.Print("Tahun Terbit: ")
	fmt.Scanln(&bukuBaru.TahunTerbit)

	DaftarBuku = append(DaftarBuku, bukuBaru)

	fmt.Println("Buku berhasil ditambahkan!")
}

func TampilkanDaftarBuku() {
	fmt.Println("Daftar Buku:")
	for _, buku := range DaftarBuku {
		fmt.Printf("Kode: %s, Judul: %s, Pengarang: %s, Penerbit: %s, Jumlah Halaman: %d, Tahun Terbit: %d\n",
			buku.Kode, buku.Judul, buku.Pengarang, buku.Penerbit, buku.JumlahHalaman, buku.TahunTerbit)
	}
}

func HapusBuku() {
	input := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Kode Buku yang Ingin Dihapus: ")
	kodeBuku, _ := input.ReadString('\n')
	kodeBuku = strings.TrimSpace(kodeBuku)

	for i, buku := range DaftarBuku {
		if buku.Kode == kodeBuku {
			DaftarBuku = append(DaftarBuku[:i], DaftarBuku[i+1:]...)
			fmt.Println("Buku berhasil dihapus!")
			return
		}
	}

	fmt.Println("Buku dengan kode tersebut tidak ditemukan.")
}

func UbahBuku() {
	input := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Kode Buku yang Ingin Diubah: ")
	kodeBuku, _ := input.ReadString('\n')
	kodeBuku = strings.TrimSpace(kodeBuku)

	for i, buku := range DaftarBuku {
		if buku.Kode == kodeBuku {
			fmt.Println("Masukkan informasi baru:")

			var bukuBaru Buku

			fmt.Print("Judul Buku: ")
			bukuBaru.Judul, _ = input.ReadString('\n')
			bukuBaru.Judul = strings.TrimSpace(bukuBaru.Judul)

			fmt.Print("Pengarang: ")
			bukuBaru.Pengarang, _ = input.ReadString('\n')
			bukuBaru.Pengarang = strings.TrimSpace(bukuBaru.Pengarang)

			fmt.Print("Penerbit: ")
			bukuBaru.Penerbit, _ = input.ReadString('\n')
			bukuBaru.Penerbit = strings.TrimSpace(bukuBaru.Penerbit)

			fmt.Print("Jumlah Halaman: ")
			fmt.Scanln(&bukuBaru.JumlahHalaman)

			fmt.Print("Tahun Terbit: ")
			fmt.Scanln(&bukuBaru.TahunTerbit)

			DaftarBuku[i] = bukuBaru
			fmt.Println("Buku berhasil diubah!")
			return
		}
	}

	fmt.Println("Buku dengan kode tersebut tidak ditemukan.")
}

func main() {
	for {
		fmt.Println("\n===== APLIKASI MANAJEMEN DAFTAR BUKU PERPUSTAKAAN =====")
		fmt.Println("1. Tambah Buku")
		fmt.Println("2. Tampilkan Daftar Buku")
		fmt.Println("3. Hapus Buku")
		fmt.Println("4. Ubah Buku")
		fmt.Println("5. Keluar")
		fmt.Print("Pilihan: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			TambahBuku()
		case 2:
			TampilkanDaftarBuku()
		case 3:
			HapusBuku()
		case 4:
			UbahBuku()
		case 5:
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}
