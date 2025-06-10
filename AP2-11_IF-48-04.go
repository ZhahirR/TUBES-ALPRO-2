package main

import "fmt"

const nmax int = 100
const maxAkun int = 3

type bahanMakanan struct {
	namaBahan  string
	jumlahStok int
	tgl        int
}

type arrBahanMakanan [nmax]bahanMakanan

type akun struct {
	nama          string
	tipe          string
	milikPengguna arrBahanMakanan
	jumlahData    int
}

var akunAktif akun
var indexAkunAktif int

var daftarAkun = [maxAkun]akun{
	{nama: "User 1", tipe: "personal"},
	{nama: "User 2", tipe: "personal"},
	{nama: "User 3", tipe: "personal"},
}

// var tempBahan arrBahanMakanan

func main() {
	var pilihanMenu string
	var pilihanUrutan string

	// menampilkan menu untuk pertama kali
	menuAkun()
	menuUtama(&pilihanMenu)
	// perulangan di sini berfungsi sebagai penentu menu yang akan dimasuki beredasarkan nilai pilihanMenu
	for {
		switch pilihanMenu {
		case "1":
			menuAkun()
		case "2":
			menuTambahan(&daftarAkun[indexAkunAktif].milikPengguna, &daftarAkun[indexAkunAktif].jumlahData)
		case "3":
			menuUbah(&daftarAkun[indexAkunAktif].milikPengguna, daftarAkun[indexAkunAktif].jumlahData)
		case "4":
			menuDelete(&daftarAkun[indexAkunAktif].milikPengguna, &daftarAkun[indexAkunAktif].jumlahData)
		case "5":
			fmt.Println("1. ASCENDING ")
			fmt.Print("2. DESCENDING \nPilih : ")
			fmt.Scan(&pilihanUrutan)
			for pilihanUrutan != "1" && pilihanUrutan != "2" {
				fmt.Print("Pilih : ")
				fmt.Scan(&pilihanUrutan)
			}
			pengurutanStok(&daftarAkun[indexAkunAktif].milikPengguna, daftarAkun[indexAkunAktif].jumlahData, pilihanUrutan)
		case "6":
			fmt.Println("1. ASCENDING")
			fmt.Print("2. DESCENDING \nPilih : ")
			fmt.Scan(&pilihanUrutan)
			for pilihanUrutan != "1" && pilihanUrutan != "2" {
				fmt.Print("Pilih : ")
				fmt.Scan(&pilihanUrutan)
			}
			pengurutanExp(&daftarAkun[indexAkunAktif].milikPengguna, daftarAkun[indexAkunAktif].jumlahData, pilihanUrutan)
		case "7":
			daftarBahan(&daftarAkun[indexAkunAktif].milikPengguna, &daftarAkun[indexAkunAktif].jumlahData)
		case "8":
			fmt.Println("---------------------------------------------------")
			fmt.Println("*                                                 *")
			fmt.Println("* Terima kasih sudah menggunakan aplikasi kami :) *")
			fmt.Println("*                                                 *")
			fmt.Println("---------------------------------------------------")
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
		if pilihanMenu == "8" {
			break
		}
		if pilihanMenu != "8" {
			menuUtama(&pilihanMenu)
		}
	}
}

func menuUtama(pilihanMenu *string) {
	/*
		IS. diberikan sebuah string yang berfungsi sebagai penentu masuknya kedalam menu bagian mana berdasarkan angka uang nantinya di beri masukan oleh user, selain itu prosedur ini juga mengeluarkan tampilan menunya
		FS. string yang dimasukan oleh user tadi di kembalikan ke main fungsi
	*/

	fmt.Println("------------------------------------------------")
	fmt.Println("*                                              *")
	fmt.Println("*            APLIKASI BAHAN MAKANAN            *")
	fmt.Println("*                                              *")
	fmt.Println("------------------------------------------------")
	// fmt.Println("")
	fmt.Println("----------Menu Utama--------")
	fmt.Println("1. Akun")
	fmt.Println("2. Tambahan ")
	fmt.Println("3. Ubah (Sequential Search)")
	fmt.Println("4. Delete (Binary Search)")
	fmt.Println("5. Urutkan Berdasarkan Stok (insertionSort)")
	fmt.Println("6. Urutkan berdasarkan Kadaluarsa (SelectionsSort)")
	fmt.Println("7. Daftar Bahan ")
	fmt.Println("8. Keluar ")
	fmt.Println("----------------------------")
	fmt.Print("Pilihan: ")
	fmt.Scan(&*pilihanMenu)
	fmt.Println("")

	// akan melakukan perulangan selama nilai dari pilihan menu bukan 1,2,3,4,5 . angka tersebut merupakan string, lalu user akan terus melakukan inputan sampai nilai yang di maksudkan ada
	for *pilihanMenu != "1" && *pilihanMenu != "2" && *pilihanMenu != "3" && *pilihanMenu != "4" && *pilihanMenu != "5" && *pilihanMenu != "6" && *pilihanMenu != "7" && *pilihanMenu != "8" {
		fmt.Print("Pilihan: ")
		fmt.Scan(&*pilihanMenu)
	}

}
func menuTambahan(dataBahan *arrBahanMakanan, n *int) {
	/*
		IS. diberikan sebuah data array yang berfungsi sebagai menyimpan data yang akan di tambahkan pleh user, lalu ada n yang fungsi nya banyaknya koleksi data dari array
		FS. dataBahan akan di kembalikan setelah user menginputkan data yang akan di tambahkan, lalu n dikembalikan untuk diketahui banyaknya data yang telah di tambahkan
	*/
	// variabel piilihan berguna untuk menentukan data yang ditambahkan akan disimpan atau tidak berdasarkan inputana angka 1 atau 2

	// variabel simpanSementara akan digunakan sebagai array sementara yang nantinya akan di simpan ke dalam koleksi data utama atau tidak
	var simpanSementara bahanMakanan
	var simpan string

	fmt.Println("------------------------------------------------")
	fmt.Println("*                                              *")
	fmt.Println("*            APLIKASI BAHAN MAKANAN            *")
	fmt.Println("*                                              *")
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("----------Tambahan--------")
	fmt.Print("Nama Bahan: ")
	fmt.Scan(&simpanSementara.namaBahan)
	for {
		fmt.Print("Jumlah Stok: ")
		fmt.Scan(&simpanSementara.jumlahStok)
		if simpanSementara.jumlahStok > 0 {
			break
		}
		fmt.Println("Jumlah stok harus lebih dari 0.")
	}
	for {
		fmt.Print("Hari hingga kadaluarsa: ")
		fmt.Scan(&simpanSementara.tgl)
		if simpanSementara.tgl > 0 {
			break
		}
		fmt.Println("Hari kadaluarsa harus lebih dari 0.")
	}

	fmt.Print("Simpan data ini? (y/n): ")
	fmt.Scan(&simpan)

	if simpan == "y" || simpan == "Y" {
		if *n < nmax {
			dataBahan[*n] = simpanSementara
			*n = *n + 1
			fmt.Println("Data berhasil disimpan!")
		} else {
			fmt.Println("Gagal menyimpan: data penuh.")
		}
	} else {
		fmt.Println("Data tidak disimpan.")
	}
}

func menuUbah(dataBahan *arrBahanMakanan, n int) {
	/*
		IS. diberikan sebuah array yang memiliki beberapa data yang berfungsi sebagai pilihan yang akan di ubah
		FS. User memilih data yang akan di ubah berdasarkan nama dari suatu data yang ada di dalam array, lalu memilih mau menyimpan perubahan atau tidak,
	*/

	var nama, simpan string
	var idx int
	var simpanSementara bahanMakanan

	fmt.Println("------------------------------------------------")
	fmt.Println("*                                              *")
	fmt.Println("*            APLIKASI BAHAN MAKANAN            *")
	fmt.Println("*                                              *")
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("-------------------Menu Ubah-------------------")
	fmt.Println("")
	fmt.Println("")

	// menampilkan data dalam array fungsinya agar user lebih mudah melihat data yang ada
	daftarBahan(dataBahan, &n)

	fmt.Print("Cari Nama Data yang akan di ubah(sequencial search): ")
	// Diberikan kata kunci untuk mencari data yang akan di ubah berdasarkan nama data/item
	fmt.Scan(&nama)
	idx = sequentialSearch(*dataBahan, n, nama)
	if idx == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	fmt.Print("Nama baru: ")
	fmt.Scan(&simpanSementara.namaBahan)
	for {
		fmt.Print("Jumlah Stok baru: ")
		fmt.Scan(&simpanSementara.jumlahStok)
		if simpanSementara.jumlahStok > 0 {
			break
		}
		fmt.Println("Jumlah stok harus lebih dari 0.")
	}

	for {
		fmt.Print("Hari hingga kadaluarsa baru: ")
		fmt.Scan(&simpanSementara.tgl)
		if simpanSementara.tgl > 0 {
			break
		}
		fmt.Println("Hari kadaluarsa harus lebih dari 0.")
	}

	fmt.Print("Simpan perubahan? (y/n): ")
	fmt.Scan(&simpan)

	if simpan == "y" || simpan == "Y" {
		dataBahan[idx] = simpanSementara
		fmt.Println("Data berhasil diubah!")
	} else {
		fmt.Println("Perubahan dibatalkan.")
	}
}

func menuDelete(databahan *arrBahanMakanan, n *int) {
	/*
		IS. Diberikan data bahan dari menu utama tipe data bentukan, dan n yang merepresentasikan jumlah array arrBahanMakann
		FS. Jika data ditemukan dan pengguna mengonfirmasi penghapusan, maka satu data bahan makanan
			yang sesuai dengan nama yang dimasukkan akan dihapus dari array, dan nilai *n berkurang 1.
		    Jika pengguna membatalkan, maka data tetap tidak berubah.
	*/
	var nama string
	var konfirmasi string
	var idx int

	fmt.Println("------------------------------------------------")
	fmt.Println("*                                              *")
	fmt.Println("*            APLIKASI BAHAN MAKANAN            *")
	fmt.Println("*                                              *")
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("----------Delete--------")
	for {
		fmt.Print("Masukkan nama bahan yang ingin dihapus: ")
		fmt.Scan(&nama)
		idx = binSearch(*databahan, *n, nama)

		if idx != -1 {
			break
		} else {
			fmt.Println("Data tidak ditemukan. Silakan coba lagi.")
		}
	}

	fmt.Print("Yakin ingin menghapus data ini? (y/n): ")
	fmt.Scan(&konfirmasi)

	if konfirmasi == "y" || konfirmasi == "Y" {
		for i := idx; i < *n-1; i++ {
			databahan[i] = databahan[i+1]
		}
		*n--
		fmt.Println("Data berhasil dihapus!")
	} else {
		fmt.Println("Penghapusan dibatalkan.")
	}

}
func daftarBahan(dataBahan *arrBahanMakanan, n *int) {
	/*
		IS. Tersedia data bahan makanan sebanyak *n elemen.

		FS. Menampilkan daftar bahan makanan beserta stok dan hari kedaluwarsa.
		    Juga menampilkan bahan yang kedaluwarsa dalam <= 14 hari.
	*/

	var i int
	var found bool

	// Tampilkan nama akun aktif
	fmt.Println("------------------------------------------------")
	fmt.Printf("Akun Aktif: %s\n", daftarAkun[indexAkunAktif].nama)
	fmt.Println("------------------------------------------------")

	if *n < 1 {
		fmt.Println("Data Belum Ada")
	} else {
		fmt.Println("--------------------- Daftar Bahan Makanan ---------------------")

		// Header tabel
		fmt.Printf("*%-5s | %-20s | %-12s | %-12s\n", "No", "Nama Bahan", "Jumlah Stok", "Kadaluarsa(hari)*")
		fmt.Println("----------------------------------------------------------------")

		// Isi tabel
		for i = 0; i < *n; i++ {
			fmt.Printf("*%-5d | %-20s | %-12d | %-12d    *\n", i+1, dataBahan[i].namaBahan, dataBahan[i].jumlahStok, dataBahan[i].tgl)
		}
		fmt.Println("----------------------------------------------------------------")
		fmt.Println()
		fmt.Println("\nBahan yang hampir kadaluarsa (<= 14 hari):")
		found = false
		for i = 0; i < *n; i++ {
			if dataBahan[i].tgl <= 14 {
				fmt.Printf("- %s (Kedaluwarsa dalam %d hari)\n", dataBahan[i].namaBahan, dataBahan[i].tgl)
				found = true
			}
		}
		if !found {
			fmt.Println("Tidak ada bahan yang hampir kadaluarsa.")
		}
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
	}
}
func pengurutanStok(databahan *arrBahanMakanan, n int, pilihan string) {
	/*
		IS. Tersedia data bahan makanan sebanyak n elemen dan pilihan metode pengurutan.

		FS. Data bahan diurutkan berdasarkan jumlah stok (ascending jika pilihan "1", descending jika "2"), lalu ditampilkan.
	*/

	if pilihan == "1" {
		InsertionSortAsc(databahan, n)
		daftarBahan(&*databahan, &n)
	} else if pilihan == "2" {
		InsertionSortDesc(databahan, n)
		daftarBahan(&*databahan, &n)

	}

}
func pengurutanExp(databahan *arrBahanMakanan, n int, pilihan string) {
	/*
		IS. Tersedia data bahan makanan sebanyak n elemen dan pilihan metode pengurutan.

		FS. Data bahan diurutkan berdasarkan hari kedaluwarsa (ascending jika pilihan "1", descending jika "2"), lalu ditampilkan.
	*/

	if pilihan == "1" {
		SelectionSortAsc(databahan, n)
		daftarBahan(&*databahan, &n)
	} else if pilihan == "2" {
		SelectionSortDesc(databahan, n)
		daftarBahan(&*databahan, &n)

	}

}

func sequentialSearch(T arrBahanMakanan, n int, x string) int {
	var found bool = false
	var i int = 0
	for i < n && !found {
		found = T[i].namaBahan == x
		i++
	}
	if found {
		return i - 1
	} else {
		return -1
	}

}

func binSearch(tab arrBahanMakanan, n int, x string) int {
	var mid, left, right int
	var found int = -1
	left = 0
	right = n - 1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if x > tab[mid].namaBahan {
			right = mid - 1
		} else if x < tab[mid].namaBahan {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found
}

func InsertionSortAsc(A *arrBahanMakanan, N int) {
	/*
		IS. Tersedia array A sebanyak N elemen.

		FS. Array diurutkan secara ascending berdasarkan jumlah stok.
	*/

	var i, pass int
	var temp bahanMakanan
	pass = 1
	for pass <= N-2 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.jumlahStok < A[i-1].jumlahStok {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass = pass + 1
	}
}
func InsertionSortDesc(A *arrBahanMakanan, N int) {
	/*
		IS. Tersedia array A sebanyak N elemen.

		FS. Array diurutkan secara descending berdasarkan jumlah stok.
	*/

	var i, pass int
	var temp bahanMakanan
	for pass = 1; pass < N; pass++ {
		temp = A[pass]
		i = pass
		for i > 0 && temp.jumlahStok > A[i-1].jumlahStok {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
}

// Fungsi Selection Sort (berdasarkan hari kadaluarsa DESC)
func SelectionSortDesc(A *arrBahanMakanan, N int) {
	/*
		IS. Tersedia array A sebanyak N elemen.

		FS. Array diurutkan secara ascending berdasarkan hari kedaluwarsa.
	*/

	var i, j, maxIdx int
	var temp bahanMakanan

	for i = 0; i < N-1; i++ {
		maxIdx = i
		for j = i + 1; j < N; j++ {
			if A[j].tgl > A[maxIdx].tgl {
				maxIdx = j
			}
		}
		// Tukar elemen
		temp = A[i]
		A[i] = A[maxIdx]
		A[maxIdx] = temp
	}
}
func SelectionSortAsc(A *arrBahanMakanan, N int) {
	/*
		IS. Tersedia array A sebanyak N elemen.

		FS. Array diurutkan secara descending berdasarkan hari kedaluwarsa.
	*/

	var i, j, minIdx int
	var temp bahanMakanan

	for i = 0; i < N-1; i++ {
		minIdx = i
		for j = i + 1; j < N; j++ {
			if A[j].tgl < A[minIdx].tgl {
				minIdx = j
			}
		}
		// Tukar elemen
		temp = A[i]
		A[i] = A[minIdx]
		A[minIdx] = temp
	}
}

func menuAkun() {
	/*
		IS. Menampilkan daftar akun (3 akun).

		FS. Akun aktif dipilih oleh user berdasarkan input 1-3.
	*/

	var pilihan int
	for {
		fmt.Println("------------------------------------------------")
		fmt.Println("*                  PILIH AKUN                  *")
		fmt.Println("------------------------------------------------")
		fmt.Println(">> Akun Personal:")
		for i := 0; i < 3; i++ {
			fmt.Printf("%d. %s\n", i+1, daftarAkun[i].nama)
		}

		fmt.Print("Pilih akun (1-3): ")
		fmt.Scan(&pilihan)

		if pilihan >= 1 && pilihan <= 3 {
			break
		} else {
			fmt.Println("Pilihan tidak valid. Silakan pilih angka 1-3.")
		}
	}

	indexAkunAktif = pilihan - 1
	fmt.Println("Akun berhasil dipilih!")
	fmt.Printf("Nama: %s\n", daftarAkun[indexAkunAktif].nama)
}
