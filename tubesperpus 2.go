package main

import "fmt"

type Buku struct {
	Judul      string
	Penulis     string
	Peminjaman bool
}

var perpus []Buku

func main() {
	for {
		fmt.Println("\n=== MENU PERPUSTAKAAN ===")
		fmt.Println("1. Lihat daftar buku")
		fmt.Println("2. Tambah buku")
		fmt.Println("3. Pinjam buku")
		fmt.Println("4. Kembalikan buku")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			listBuku()

		case 2:
			var Judul, Penulis string

			fmt.Print("Judul buku (tanpa spasi): ")
			fmt.Scan(&Judul)

			fmt.Print("Nama penulis (tanpa spasi): ")
			fmt.Scan(&Penulis)

			tambahbuku(Judul, Penulis)

		case 3:
			var Judul string
			fmt.Print("Judul buku yang dipinjam (tanpa spasi): ")
			fmt.Scan(&Judul)

			pinjambuku(Judul)

		case 4:
			var Judul string
			fmt.Print("Judul buku yang dikembalikan (tanpa spasi): ")
			fmt.Scan(&Judul)

			kembalikanbuku(Judul)

		case 5:
			fmt.Println("Terima kasih, program selesai.")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func listBuku() {
	if len(perpus) == 0 {
		fmt.Println("Belum ada buku di perpustakaan.")
		return
	}

	fmt.Println("\n=== DAFTAR BUKU ===")
	for i, b := range perpus {
		status := "Tersedia"
		if b.Peminjaman {
			status = "Dipinjam"
		}
		fmt.Printf("%d. %s - %s (%s)\n", i+1, b.Judul, b.Penulis, status)
	}
}

func tambahbuku(Judul, Penulis string) {
	perpus = append(perpus, Buku{
		Judul:      Judul,
		Penulis:     Penulis,
		Peminjaman: false,
	})

	fmt.Println("Buku berhasil ditambahkan!")
}

func pinjambuku(Judul string) {
	for i := range perpus {
		if perpus[i].Judul == Judul {
			if perpus[i].Peminjaman {
				fmt.Println("Buku sedang dipinjam!")
				return
			}
			perpus[i].Peminjaman = true
			fmt.Println("Berhasil meminjam buku.")
			return
		}
	}
	fmt.Println("Buku tidak ditemukan!")
}

func kembalikanbuku(Judul string) {
	for i := range perpus {
		if perpus[i].Judul == Judul {
			if !perpus[i].Peminjaman {
				fmt.Println("Buku ini tidak sedang dipinjam.")
				return
			}
			perpus[i].Peminjaman = false
			fmt.Println("Buku berhasil dikembalikan!")
			return
		}
	}
	fmt.Println("Buku tidak ditemukan!")
}
