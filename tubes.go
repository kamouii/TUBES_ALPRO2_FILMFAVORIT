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
	var judul string

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
			

		case 2:
			tampilFilm(A, n)

		case 3:
			

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
				fmt.Scan(&judul)
				
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
				
			case 2:
				
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
						
					case 2:
						
					default:
						fmt.Println("Pilihan tidak valid")
					}
		case 8:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			return
				}
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

		*n = *n - 1
		fmt.Println("Film berhasil dihapus")
	} else {
		fmt.Println("Film tidak ditemukan")
	}
}

func cariFilmJudul(A DaftarFilm, n int, judul string)  {
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
