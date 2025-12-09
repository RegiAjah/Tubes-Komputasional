package main

import "fmt"

type Book struct {
	Title      string
	Author     string
	IsBorrowed bool
}

var library []Book

func main() {
	for {
		fmt.Println("\n=== MENU PERPUSTAKAAN ===")
		fmt.Println("1. Lihat daftar buku")
		fmt.Println("2. Tambah buku")
		fmt.Println("3. Pinjam buku")
		fmt.Println("4. Kembalikan buku")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			showBooks()

		case 2:
			var title, author string

			fmt.Print("Judul buku (tanpa spasi): ")
			fmt.Scan(&title)

			fmt.Print("Nama penulis (tanpa spasi): ")
			fmt.Scan(&author)

			addBook(title, author)

		case 3:
			var title string
			fmt.Print("Judul buku yang dipinjam (tanpa spasi): ")
			fmt.Scan(&title)

			borrowBook(title)

		case 4:
			var title string
			fmt.Print("Judul buku yang dikembalikan (tanpa spasi): ")
			fmt.Scan(&title)

			returnBook(title)

		case 5:
			fmt.Println("Terima kasih, program selesai.")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func showBooks() {
	if len(library) == 0 {
		fmt.Println("Belum ada buku di perpustakaan.")
		return
	}

	fmt.Println("\n=== DAFTAR BUKU ===")
	for i, b := range library {
		status := "Tersedia"
		if b.IsBorrowed {
			status = "Dipinjam"
		}
		fmt.Printf("%d. %s - %s (%s)\n", i+1, b.Title, b.Author, status)
	}
}

func addBook(title, author string) {
	library = append(library, Book{
		Title:      title,
		Author:     author,
		IsBorrowed: false,
	})

	fmt.Println("Buku berhasil ditambahkan!")
}

func borrowBook(title string) {
	for i := range library {
		if library[i].Title == title {
			if library[i].IsBorrowed {
				fmt.Println("Buku sedang dipinjam!")
				return
			}
			library[i].IsBorrowed = true
			fmt.Println("Berhasil meminjam buku.")
			return
		}
	}
	fmt.Println("Buku tidak ditemukan!")
}

func returnBook(title string) {
	for i := range library {
		if library[i].Title == title {
			if !library[i].IsBorrowed {
				fmt.Println("Buku ini tidak sedang dipinjam.")
				return
			}
			library[i].IsBorrowed = false
			fmt.Println("Buku berhasil dikembalikan!")
			return
		}
	}
	fmt.Println("Buku tidak ditemukan!")
}
