package main

import "fmt"

type Film struct {
	judul      string
	genre      string
	tahunRilis int
	deskripsi  string
	rating     float64
}

const NMAX int = 9999

type DaftarFilm [NMAX]Film

func main() {
	var A DaftarFilm
	var n int = 10
	var pilih int
	var judul, genre string

	A[0] = Film{"Interstellar", "SciFi", 2014, "LuarAngkasa", 9.0}
	A[1] = Film{"Joker", "Drama", 2019, "ArthurFleck", 8.5}
	A[2] = Film{"Avengers", "Action", 2012, "Superhero", 8.8}
	A[3] = Film{"Titanic", "Romance", 1997, "KapalTitanic", 8.2}
	A[4] = Film{"Inception", "SciFi", 2010, "MimpiBerlapis", 9.1}
	A[5] = Film{"Frozen", "Animation", 2013, "PetualanganElsa", 7.8}
	A[6] = Film{"Batman", "Action", 2022, "KsatriaMalam", 8.4}
	A[7] = Film{"Coco", "Animation", 2017, "DuniaArwah", 8.6}
	A[8] = Film{"Parasite", "Thriller", 2019, "KesenjanganSosial", 8.9}
	A[9] = Film{"Up", "Animation", 2009, "BalonUdara", 8.7}

	for {
		fmt.Println("|================================================|")
		fmt.Println("|==========Aplikasi Review Film Favorit==========|")
		fmt.Println("|================================================|")
		fmt.Println("|1. Tambah Film                                  |")
		fmt.Println("|2. Lihat Daftar Film                            |")
		fmt.Println("|3. Edit Film                                    |")
		fmt.Println("|4. Hapus Film                                   |")
		fmt.Println("|5. Cari Film                                    |")
		fmt.Println("|6. Urutkan Film                                 |")
		fmt.Println("|7. Statistik Film                               |")
		fmt.Println("|8. Keluar                                       |")
		fmt.Println("|================================================|")

		fmt.Print("Pilih Menu : ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahFilm(&A, &n)
		case 2:
			tampilFilm(A, n)

		case 3:
			tampilFilm(A, n)
			fmt.Print("Masukkan judul film yang akan diedit : ")
			fmt.Scan(&judul)
			editFilm(&A, n, judul)
		case 4:
			tampilFilm(A, n)
			fmt.Print("Masukkan judul film yang akan dihapus : ")
			fmt.Scan(&judul)
			hapusFilm(&A, &n, judul)

		case 5:
			fmt.Print("1. Cari Film Berdasarkan Judul\n")
			fmt.Print("2. Cari Film Berdasarkan Genre\n")
			fmt.Print("Pilih opsi : ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				fmt.Print("Masukkan judul film yang akan dicari : ")
				fmt.Scan(&judul)
				cariFilmJudul(A, n, judul)

			case 2:
				fmt.Print("Masukkan genre film yang akan dicari : ")
				fmt.Scan(&genre)
				
			default:
				fmt.Println("Pilihan tidak valid")
			}

		case 6:
			fmt.Print("1. Urutkan Film Berdasarkan Rating\n")
			fmt.Print("2. Urutkan Film Berdasarkan Tahun Rilis\n")
			fmt.Print("Pilih opsi : ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				urutFilmRating(&A, n)
				tampilFilm(A, n)
				fmt.Println("Film berhasil diurutkan berdasarkan rating tertinggi")
			case 2:
				urutFilmTahunRilisInsertionSortDescending(&A, n)
				tampilFilm(A, n)
				fmt.Println("Film berhasil diurutkan berdasarkan tahun rilis yang terbaru")
			default:
				fmt.Println("Pilihan tidak valid")
			}
		case 7:
			fmt.Print("1. Statistik Berdasarkan Genre\n")
			fmt.Print("2. Statistik Berdasarkan Rating\n")
			fmt.Print("Pilih opsi : ")
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				statistikFilmGenre(A, n)
			case 2:
				rata := rataRataRating(A, n)
				fmt.Printf("Rata-rata rating film: %.2f\n", rata)
				fmt.Println("Rata-rata rating film berhasil dihitung")
			default:
				fmt.Println("Pilihan tidak valid")
			}
		case 8:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
		}
	}
}
//----------------------------------------------------CRUD--------------------------------------------
//menambahkan film baru ke daftar film
func tambahFilm(A *DaftarFilm, n *int) {
	var judul, genre, deskripsi string
	var tahunRilis int
	var rating float64

	fmt.Print("Masukkan judul film : ")
	fmt.Scan(&judul)
	fmt.Print("Masukkan genre film : ")
	fmt.Scan(&genre)
	fmt.Print("Masukkan tahun rilis film : ")
	fmt.Scan(&tahunRilis)
	fmt.Print("Masukkan deskripsi film : ")
	fmt.Scan(&deskripsi)
	fmt.Print("Masukkan rating film : ")
	fmt.Scan(&rating)

	if *n < NMAX {
		A[*n].judul = judul
		A[*n].genre = genre
		A[*n].tahunRilis = tahunRilis
		A[*n].deskripsi = deskripsi
		A[*n].rating = rating

		*n = *n + 1
		fmt.Println("Film berhasil ditambahkan")
	} else {
		fmt.Println("Daftar film penuh")
	}
}

// menampilkan daftar film
func tampilFilm(A DaftarFilm, n int) {
	var i int

	if n == 0 {
		fmt.Println("Daftar film kosong")
		return
	}

	fmt.Println("===================================================================================================")
	fmt.Printf("| %-3s | %-20s | %-10s | %-5s | %-30s | %-6s |\n",
		"No", "Judul", "Genre", "Tahun", "Deskripsi", "Rating")
	fmt.Println("===================================================================================================")

	for i = 0; i < n; i++ {
		fmt.Printf("| %-3d | %-20s | %-10s | %-5d | %-30s | %-6.2f |\n",
			i+1,
			A[i].judul,
			A[i].genre,
			A[i].tahunRilis,
			A[i].deskripsi,
			A[i].rating)
	}

	fmt.Println("===================================================================================================")
}

// edit film selectioon sort tanpa break
func editFilm(A *DaftarFilm, n int, judul string) {
	var i int
	var ketemu bool

	ketemu = false

	for i = 0; i < n; i++ {
		if (*A)[i].judul == judul {
			ketemu = true

			fmt.Print("Masukkan judul baru : ")
			fmt.Scan(&(*A)[i].judul)

			fmt.Print("Masukkan genre baru : ")
			fmt.Scan(&(*A)[i].genre)

			fmt.Print("Masukkan tahun rilis baru : ")
			fmt.Scan(&(*A)[i].tahunRilis)

			fmt.Print("Masukkan deskripsi baru : ")
			fmt.Scan(&(*A)[i].deskripsi)

			fmt.Print("Masukkan rating baru : ")
			fmt.Scan(&(*A)[i].rating)

			fmt.Println("Film berhasil diedit")
		}
	}

	if !ketemu {
		fmt.Println("Film tidak ditemukan")
	}
}

// menghapus film dari daftar film
func hapusFilm(A *DaftarFilm, n *int, judul string) {
	var i int
	var ketemu bool = false

	for i = 0; i < *n && !ketemu; i++ {
		if A[i].judul == judul {
			A[i] = A[*n-1]
			*n--
			ketemu = true
		}
	}

	if ketemu {
		fmt.Println("Film berhasil dihapus")
	} else {
		fmt.Println("Film tidak ditemukan")
	}
}
//----------------------------------------------------------------------------CRUD--------------------------------------------


//----------------------------------------------------CARI--------------------------------------------
//cari film berdasarkan judul menggunakan sequential search
func cariFilmJudul(A DaftarFilm, n int, judul string) {
	var i int
	var ketemu bool

	ketemu = false
	i = 0

	for i < n && !ketemu {
		if A[i].judul == judul {
			ketemu = true
			fmt.Print("----------------------------------------------\n")
			fmt.Println("Film ditemukan")
			fmt.Println("Judul      :", A[i].judul)
			fmt.Println("Genre      :", A[i].genre)
			fmt.Println("Tahun Rilis:", A[i].tahunRilis)
			fmt.Println("Deskripsi  :", A[i].deskripsi)
			fmt.Printf("Rating     : %.2f\n", A[i].rating)
			fmt.Print("----------------------------------------------\n")
		}
		i++
	}

	if !ketemu {
		fmt.Println("Film tidak ditemukan")
	}
}

//cari film berdasarkan genre menggunakan binary search

//----------------------------------------------------CARI--------------------------------------------

//----------------------------------------------------URUTAN--------------------------------------------
//mengurutkan film berdasarkan rating tertinggi menggunakan selection sort
func urutFilmRating(A *DaftarFilm, n int) {
	var i, j, idxMin int
	var temp Film

	for i = 0; i < n-1; i++ {
		idxMin = i

		for j = i + 1; j < n; j++ {
			if (*A)[j].rating > (*A)[idxMin].rating {
				idxMin = j
			}
		}

		temp = (*A)[i]
		(*A)[i] = (*A)[idxMin]
		(*A)[idxMin] = temp
	}
}
//mengurutkan film berdasarkan tahun rilis terbaru menggunakan insertion sort
func urutFilmTahunRilisInsertionSortDescending(A *DaftarFilm, n int) {
	var i, j int
	var temp Film
	for i = 1; i < n; i++ {
		temp = A[i]
		j = i - 1
		for j >= 0 && A[j].tahunRilis < temp.tahunRilis {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = temp
	}
}
//----------------------------------------------------URUTAN--------------------------------------------

//------------------------------------STATISTIK--------------------------------------------
//rata-rata rating film
func rataRataRating(A DaftarFilm, n int) float64 {
	var total float64
	var i int

	total = 0
	for i = 0; i < n; i++ {
		total += A[i].rating
	}

	return total / float64(n)
}
// statistik film berdasarkan genre
func statistikFilmGenre(A DaftarFilm, n int) {
	var genre string
	var count int = 0
	fmt.Print("Masukkan genre film yang akan dicari : ")
	fmt.Scan(&genre)
	for i := 0; i < n; i++ {
		if A[i].genre == genre {
			count++
		}
	}
	fmt.Printf("Jumlah film dengan genre %s: %d\n", genre, count)
}
//------------------------------------STATISTIK--------------------------------------------