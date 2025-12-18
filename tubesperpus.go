package main

import "fmt"

type Buku struct {
	//TIPE DATA
	Judul      string 
	Penulis     string
	Peminjaman bool
}

var perpus []Buku
// [] untuk mendefinisikan didalam variabel perpus ada buku

func main() {
	// menampilkan tampilan menu

	for {
	// untuk looping (mengulang tampilan menu)

		fmt.Println("\n=== MENU PERPUSTAKAAN ===")
		// \n untuk membuat baris baru (enter)
		fmt.Println("1. Lihat daftar buku")
		fmt.Println("2. Tambah buku")
		fmt.Println("3. Pinjam buku")
		fmt.Println("4. Kembalikan buku")
		fmt.Println("5. Keluarkan/Hapus buku")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)
		// untuk melihat pengguna menginput menu brp & digunakan untuk mencari alamat

		switch pilihan {
		// switch dan case itu percabangan
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
		
		// Di dalam switch pilihan pada func main()
		case 5:
			var Judul, Penulis string
			fmt.Print("Judul buku yang akan dihapus: ")
			fmt.Scan(&Judul)
			fmt.Print("Nama penulis buku tersebut: ")
			fmt.Scan(&Penulis)

			hapusbuku(Judul, Penulis)

		case 6:
			fmt.Println("Terima kasih, program selesai.")
			return
			// untuk pengulangan

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func listBuku() {
	if len(perpus) == 0 {
	// len itu untuk mendeteksi jumlah buku di dalam perpus

		fmt.Println("Belum ada buku di perpustakaan.")
		return
	}

	fmt.Println("\n=== DAFTAR BUKU ===")
	for i, b := range perpus {
	// i (index) untuk nomor urut / posisi elemen
	// b (value) untuk menampung isi / data dari elemen
	// range untuk membongkar isi slice

		status := "Tersedia"
		if b.Peminjaman {
			status = "Dipinjam"
		}
		fmt.Printf("%d. %s - %s (%s)\n", i+1, b.Judul, b.Penulis, status)
		// %d (desimal) buat angka - i+1 
		// %s (string) buat judul, penulis, status
	}
}

func tambahbuku(Judul, Penulis string) {
	// cek apakah judul sudah ada (Case Sensitive)
	for _, b := range perpus {
		// Jika judul DAN penulis sama persis, anggap buku sudah ada
		if b.Judul == Judul && b.Penulis == Penulis {
			fmt.Printf("\nGagal: Buku '%s' karya %s sudah terdaftar!\n", Judul, Penulis)
			return
		}
	}

	// append untuk mendifinisikan var perpus dengan elemen baru yaitu judul,penulis dan status peminjaman
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
			// kita menggunakan fungsi boolean true untuk mendeteksi apakah buku berhasil dipinjam
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

func hapusbuku(Judul string, Penulis string) {
	indexHapus := -1
	// indexhapus digunakan untuk mengurangi jumlah buku yang ada di daftar

	// cari indeks buku berdasarkan judul
	for i, b := range perpus {
		if b.Judul == Judul && b.Penulis == Penulis {
			indexHapus = i
			break
			// break untuk memberhintikan fungsi hapusbuku dan balik ke mmenu
		}	
	}

	// jika ketemu, hapus dari slice
	if indexHapus != -1 {
		// Teknik menghapus elemen slice: 
		// Mengambil elemen sebelum indexHapus dan menggabungkannya dengan elemen setelah indexHapus
		perpus = append(perpus[:indexHapus], perpus[indexHapus+1:]...)
		fmt.Printf("Buku '%s' karya %s berhasil dihapus!\n", Judul, Penulis)
	} else {
		fmt.Println("Gagal: Buku dengan judul dan penulis tersebut tidak ditemukan.")
	}
}