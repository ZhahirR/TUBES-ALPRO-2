package main

import "fmt"

const nmax int = 100

type bahanMakanan struct {
	namaBahan  string
	jumlahStok int
	tgl        int // hari sampai kedaluwarsa
}

type arrBahanMakanan [nmax]bahanMakanan

var data arrBahanMakanan
var jumlahData int = 0

func main() {
	var pilihanMenu int
	for {
		menuUtama()
		fmt.Scan(&pilihanMenu)
		
		switch pilihanMenu {
		case 1:
			menuTambah()
		case 2:
			menuUbah()
		case 3:
			menuHapus()
		case 4:
			daftarBahan()
		case 5:
			laporan()
		case 6:
			fmt.Println("Terima kasih sudah menggunakan aplikasi kami :)")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuUtama() {
	fmt.Println("\n=== APLIKASI STOK BAHAN MAKANAN ===")
	fmt.Println("1. Tambah Bahan")
	fmt.Println("2. Ubah Bahan")
	fmt.Println("3. Hapus Bahan")
	fmt.Println("4. Tampilkan Daftar Bahan")
	fmt.Println("5. Laporan Total Bahan")
	fmt.Println("6. Keluar")
	fmt.Print("Pilih menu: ")
}

func menuTambah() {
	if jumlahData >= nmax {
		fmt.Println("Data penuh, tidak bisa menambah.")
		return
	}
	
	var bahan bahanMakanan
	fmt.Print("Nama Bahan: ")
	fmt.Scan(&bahan.namaBahan)
	fmt.Print("Jumlah Stok: ")
	fmt.Scan(&bahan.jumlahStok)
	fmt.Print("Hari sampai kedaluwarsa: ")
	fmt.Scan(&bahan.tgl)
	
	// Validasi input
	if bahan.jumlahStok < 0 || bahan.tgl < 0 {
		fmt.Println("Error: Stok dan hari kedaluwarsa tidak boleh negatif")
		return
	}
	
	data[jumlahData] = bahan
	jumlahData++
	fmt.Println("Data berhasil ditambahkan.")
}

func menuUbah() {
	if jumlahData == 0 {
		fmt.Println("Belum ada data bahan.")
		return
	}
	
	var nama string
	fmt.Print("Masukkan nama bahan yang ingin diubah: ")
	fmt.Scan(&nama)
	
	idx := sequentialSearch(nama)
	if idx == -1 {
		fmt.Println("Bahan tidak ditemukan.")
		return
	}
	
	fmt.Printf("Data saat ini: %s | Stok: %d | Kedaluwarsa: %d hari\n", 
		data[idx].namaBahan, data[idx].jumlahStok, data[idx].tgl)
	
	fmt.Print("Nama Baru (kosongkan jika tidak ingin diubah): ")
	var namaBaru string
	fmt.Scan(&namaBaru)
	if namaBaru != "" {
		data[idx].namaBahan = namaBaru
	}
	
	fmt.Print("Jumlah Stok Baru (0 jika tidak ingin diubah): ")
	var stokBaru int
	fmt.Scan(&stokBaru)
	if stokBaru > 0 {
		data[idx].jumlahStok = stokBaru
	}
	
	fmt.Print("Hari sampai kedaluwarsa Baru (0 jika tidak ingin diubah): ")
	var tglBaru int
	fmt.Scan(&tglBaru)
	if tglBaru > 0 {
		data[idx].tgl = tglBaru
	}
	
	fmt.Println("Data berhasil diubah.")
}

func menuHapus() {
	if jumlahData == 0 {
		fmt.Println("Belum ada data bahan.")
		return
	}
	
	var nama string
	fmt.Print("Masukkan nama bahan yang ingin dihapus: ")
	fmt.Scan(&nama)
	
	idx := sequentialSearch(nama)
	if idx == -1 {
		fmt.Println("Bahan tidak ditemukan.")
		return
	}
	
	// Konfirmasi penghapusan
	fmt.Printf("Anda yakin ingin menghapus %s? (y/n): ", data[idx].namaBahan)
	var confirm string
	fmt.Scan(&confirm)
	if confirm != "y" && confirm != "Y" {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}
	
	for i := idx; i < jumlahData-1; i++ {
		data[i] = data[i+1]
	}
	jumlahData--
	fmt.Println("Data berhasil dihapus.")
}

func daftarBahan() {
	if jumlahData == 0 {
		fmt.Println("Belum ada data bahan.")
		return
	}
	
	var pilihan int
	fmt.Println("\nTampilkan berdasarkan:")
	fmt.Println("1. Cari berdasarkan Nama")
	fmt.Println("2. Urutkan berdasarkan Tanggal Kedaluwarsa (Selection Sort)")
	fmt.Println("3. Urutkan berdasarkan Jumlah Stok (Insertion Sort)")
	fmt.Println("4. Tampilkan semua bahan")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)
	
	switch pilihan {
	case 1:
		var nama string
		fmt.Print("Masukkan nama bahan: ")
		fmt.Scan(&nama)
		idx := sequentialSearch(nama)
		if idx != -1 {
			fmt.Printf("Ditemukan: %s | Stok: %d | Kedaluwarsa: %d hari\n", 
				data[idx].namaBahan, data[idx].jumlahStok, data[idx].tgl)
		} else {
			fmt.Println("Bahan tidak ditemukan.")
		}
	case 2:
		selectionSort()
		fmt.Println("\nDaftar diurutkan berdasarkan tanggal kedaluwarsa:")
		tampilSemua()
	case 3:
		insertionSort()
		fmt.Println("\nDaftar diurutkan berdasarkan jumlah stok:")
		tampilSemua()
	case 4:
		tampilSemua()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func tampilSemua() {
	fmt.Println("\nDaftar Bahan Makanan:")
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("%d. %s | Stok: %d | Kedaluwarsa: %d hari\n", 
			i+1, data[i].namaBahan, data[i].jumlahStok, data[i].tgl)
	}
}

func laporan() {
	if jumlahData == 0 {
		fmt.Println("Belum ada data bahan.")
		return
	}
	
	total := 0
	for i := 0; i < jumlahData; i++ {
		total += data[i].jumlahStok
	}
	
	fmt.Printf("\nTotal seluruh stok bahan: %d\n", total)
	fmt.Printf("Jumlah jenis bahan: %d\n", jumlahData)
	
	// Tampilkan bahan yang hampir kadaluarsa (<= 7 hari)
	fmt.Println("\nBahan yang hampir kadaluarsa (<= 7 hari):")
	found := false
	for i := 0; i < jumlahData; i++ {
		if data[i].tgl <= 14 {
			fmt.Printf("- %s (Kedaluwarsa dalam %d hari)\n", data[i].namaBahan, data[i].tgl)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada bahan yang hampir kadaluarsa.")
	}
}

func sequentialSearch(x string) int {
	for i := 0; i < jumlahData; i++ {
		if data[i].namaBahan == x {
			return i
		}
	}
	return -1
}

func insertionSort() {
	var i, j int
	var key bahanMakanan
	
	for i = 1; i < jumlahData; i++ {
		key = data[i]
		j = i - 1
		for j >= 0 && data[j].jumlahStok > key.jumlahStok {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func selectionSort() {
	var i, j, min int
	
	for i = 0; i < jumlahData-1; i++ {
		min = i
		for j = i + 1; j < jumlahData; j++ {
			if data[j].tgl < data[min].tgl {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
}
