package main

import "fmt"

type Film struct {
	judul, genre string
	rating       float64
	tonton       int
}

const NMAX int = 9999

type DaftarFilm [NMAX]Film

func main() {
	var A DaftarFilm
	var n int
	var pilih int
	var judul string

	for {
		fmt.Println("|================================================|")
		fmt.Println("|==========Aplikasi Review Film Favorit==========|")
		fmt.Println("|================================================|")
		fmt.Println("|1. Tambah Film                                  |")
		fmt.Println("|2. Lihat Daftar Film                            |")
		fmt.Println("|3. Edit Film                                    |")
		fmt.Println("|4. Hapus Film                                   |")
		fmt.Println("|5. Keluar                                       |")
		fmt.Println("|================================================|")

		fmt.Print("Pilih Menu : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahFilm(&A, &n)

		case 2:
			tampilFilm(A, n)

		case 3:
			editFilm(&A, n)

		case 4:
			fmt.Print("Masukkan judul film yang akan dihapus : ")
			fmt.Scan(&judul)
			hapusFilm(&A, &n, judul)

		case 5:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return

		default:
			fmt.Println("Pilihan tidak valid")
		}

		fmt.Println()
	}
}

// Menampilkan daftar film
func tampilFilm(A DaftarFilm, n int) {
	var i int

	if n == 0 {
		fmt.Println("Daftar film kosong")
		return
	}

	for i = 0; i < n; i++ {
		fmt.Printf("Judul Film : %s\n", A[i].judul)
		fmt.Printf("Rating     : %.2f\n", A[i].rating)
		fmt.Printf("Viewer     : %d\n", A[i].tonton)
		fmt.Printf("Genre      : %s\n", A[i].genre)
		fmt.Println("-----------------------------")
	}
}

// Tambah film
func tambahFilm(A *DaftarFilm, n *int) {
	fmt.Print("Judul : ")
	fmt.Scan(&A[*n].judul)

	fmt.Print("Genre : ")
	fmt.Scan(&A[*n].genre)

	fmt.Print("Viewer : ")
	fmt.Scan(&A[*n].tonton)

	fmt.Print("Rating : ")
	fmt.Scan(&A[*n].rating)

	*n++
	fmt.Println("Film berhasil ditambahkan")
}

// Hapus film
func hapusFilm(A *DaftarFilm, n *int, judul string) {
	var idx int = -1
	var i int

	for i = 0; i < *n; i++ {
		if A[i].judul == judul {
			idx = i
			break
		}
	}

	if idx != -1 {
		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}

		*n--
		fmt.Println("Film berhasil dihapus")
	} else {
		fmt.Println("Film tidak ditemukan")
	}
}

// Edit film
func editFilm(A *DaftarFilm, n int) {
	var judul string
	var idx int

	fmt.Print("Masukkan judul film : ")
	fmt.Scan(&judul)

	idx = cariFilm(*A, n, judul)

	if idx != -1 {

		fmt.Print("Genre baru : ")
		fmt.Scan(&A[idx].genre)

		fmt.Print("Viewer baru : ")
		fmt.Scan(&A[idx].tonton)

		fmt.Print("Rating baru : ")
		fmt.Scan(&A[idx].rating)

		fmt.Println("Data berhasil diubah")

	} else {
		fmt.Println("Film tidak ditemukan")
	}
}

// Cari film berdasarkan judul
func cariFilm(A DaftarFilm, n int, judul string) int {
	var i int

	for i = 0; i < n; i++ {
		if A[i].judul == judul {
			return i
		}
	}

	return -1
}