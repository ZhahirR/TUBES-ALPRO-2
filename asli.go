package main

import "fmt"

const nmax int = 100

type bahanMakanan struct {
	namaBahan  string
	jumlahStok int
	tgl        int
}
type arrBahanMakanan [nmax]bahanMakanan

var tempBahan arrBahanMakanan

func main() {
	var pilihanMenu int
	menuUtama()
	fmt.Scan(&pilihanMenu)
	for pilihanMenu != 5 {
		switch pilihanMenu {
		case 1:
			menuTambahan()
		case 2:
			menuUbah()
		case 3:
			menuDelete()
		case 4:
			daftarBahan()
		case 5:
			fmt.Println("Terima kasih sudah menggunakan aplikasi kami :)")
			break
		default:
			menuUtama()
		}
		menuUtama()
		fmt.Scan(&pilihanMenu)

	}
}

func menuUtama() {
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("             APLIKASI DEMO TUBES")
	fmt.Println("")
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("----------Menu Utama--------")
	fmt.Println("1. Tambahan ")
	fmt.Println("2. Ubah ")
	fmt.Println("3. Delete ")
	fmt.Println("4. Daftar Bahan ")
	fmt.Println("5. Keluar ")
	fmt.Println("----------------------------")
	fmt.Print("Pilihan: ")
}
func menuTambahan() {
	var pilihan int
	var simpanSementara arrBahanMakanan

	fmt.Println("------------------------------------------------")
	fmt.Println("             APLIKASI DEMO TUBES")
	fmt.Println("")
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("----------Tambahan--------")
	fmt.Print("Nama Bahan: ")
	fmt.Scan(&tempBahan[0].namaBahan)

	fmt.Print("Jumlah Stok(masukan angka tidak boleh selain angka): ")
	fmt.Scan(&tempBahan[0].jumlahStok)

	fmt.Print("Masukan Berapa lama sampai Kadaluarsa(masukan angka tidak boleh selain angka): ")
	fmt.Scan(&tempBahan[0].tgl)

	fmt.Println("----------------------------")
	fmt.Println("1. Simpan, 2. Batal")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		simpanSementara[0].namaBahan = tempBahan[0].namaBahan
		simpanSementara[0].jumlahStok = tempBahan[0].jumlahStok
		simpanSementara[0].tgl = tempBahan[0].tgl

	}
	//fmt.Printf("Nama Bahan : %s \nJumlah Stok : %d\nKadaluarsa : %d\n", simpanSementara[0].namaBahan, simpanSementara[0].jumlahStok, simpanSementara[0].tgl)

}

func menuUbah() {
	fmt.Println("------------------------------------------------")
	fmt.Println("             APLIKASI DEMO TUBES")
	fmt.Println("")
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("----------Ubah--------")
	fmt.Print("Masukan Nomor Bahan: ")
	fmt.Scan()
	fmt.Println("----------------------------")
}

func menuDelete() {
	fmt.Println("------------------------------------------------")
	fmt.Println("             APLIKASI DEMO TUBES")
	fmt.Println("")
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("----------Delete--------")
	fmt.Println("Masukan Nama Bahan: ")
	fmt.Println("----------------------------")
}
func daftarBahan() {

	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("             APLIKASI DEMO TUBES")
	fmt.Println("")
	fmt.Println("------------------------------------------------")
	fmt.Println("")
	fmt.Println("----------Daftar--------")
	fmt.Println("Daftar Bahan: ")
	
	fmt.Println("----------------------------")

}

func sequentialSearch(T arrBahanMakanan, n int, x string) bool {
	var found bool = false
	var i int = 0
	for i < n && !found {
		found = T[i] == x
		i++
	}
	return found
}

func binSearch(tab arrBahanMakanan, n int, x int) bool {
	var mid, idx, left, right int
	idx = -1
	left = 0
	right = n - 1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if x < A[mid] {
			right = mid - 1
		} else if x > A[mid] {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

func InsertionSort(A *arrBahanMakanan, N int) {
	var i, pass int
	var temp int
	pass = 1
	for pass <= N-1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp < A[i-1] {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass = pass + 1
	}
}

func SelectionSort(A *arrBahanMakanan, N int) {
	var i, idx, pass int
	var temp int

	pass = 1
	for i < N {
		idx = pass - 1
		i = pass
		for i < N {
			if A[i] > A[idx] {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}
