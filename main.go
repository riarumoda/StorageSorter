package main

import (
	"fmt"
	"time"
)

// ----> Kamus Global <-----

type Produk struct {
	namaProduk string
	merek      string
	jenis      string
	harga      float64
	stok       int
}

type Transaksi struct {
	pembeli          string
	barangTerjual    Produk
	jumlahDibeli     int
	tanggalTransaksi time.Time
	subtotal         float64
}

const MAXPRODUCT int = 1024

type Data [MAXPRODUCT]Produk

type Buku [MAXPRODUCT]Transaksi

//---------------------------------------------------------------

// ----> Main Program <----
func main() {
	// Variabel untuk menentukan pilihan menu dan mencegah infinite-loop
	var determinator int

	// Variabel yang akan dijadikan penyimpanan data produk
	var dataProduk Data

	// Variabel yang akan digunakan untuk membatasi iterasi ketika membaca data dan menampilkan data
	var nData int

	// Start Menu
	menuStart()
	// Lifecycle dari Main Feature
	for determinator != 3 {
		menuProcess()
		fmt.Println("Masukkan nomor menu:")
		fmt.Print(">>>> ")
		fmt.Scan(&determinator)
		if determinator == 1 {
			konfigurasiDataProduk(&dataProduk, &nData)
		} else if determinator == 2 {
			pencatatanTransaksi()
		}
	}
	menuEnd()
}

// ----> Prosedur Untuk Menu Utama <----
func konfigurasiDataProduk(data *Data, n *int) {
	var det int
	menuHeaderKonfigurasiDataProduk()
	menuOptionsKonfigurasiDataProduk()
	fmt.Println("Masukkan menu: ")
	fmt.Print(">>>> ")
	fmt.Scan(&det)
	for det != 3 {
		if det == 1 && *n < MAXPRODUCT {
			inputDataProduk(data, n)
		} else if det == 1 && *n >= MAXPRODUCT {
			fmt.Println("Memori telah habis!")
		} else if det == 2 {
			tampilSemuaDataProduk(*data, *n)
		}
		menuOptionsKonfigurasiDataProduk()
		fmt.Println("Masukkan menu: ")
		fmt.Print(">>>> ")
		fmt.Scan(&det)
	}

}

// TODO: Buat Prosedur untuk Fitur Pencatatan Transaksi beserta logic programnya
func pencatatanTransaksi() {

}

// Input data produk
func inputDataProduk(data *Data, n *int) {
	fmt.Println("---------------------------")
	fmt.Println("P E N G I S I A N   D A T A")
	fmt.Println("---------------------------")
	fmt.Println("Nama Produk:")
	fmt.Print(">>>> ")
	fmt.Scan(&data[*n].namaProduk)
	for data[*n].namaProduk == "" {
		fmt.Println("Nama produk kosong! Mohon masukkan nama produk: ")
		fmt.Print(">>>> ")
		fmt.Scan(&data[*n].namaProduk)
	}
	fmt.Println("---------------------------")
	fmt.Println("Merek Produk:")
	fmt.Print(">>>>")
	fmt.Scan(&data[*n].merek)
	for data[*n].merek == "" {
		fmt.Println("Merek produk kosong! Mohon masukkan merek produk: ")
		fmt.Print(">>>> ")
		fmt.Scan(&data[*n].merek)
	}
	fmt.Println("---------------------------")
	fmt.Println("Jenis Produk:")
	fmt.Print(">>>> ")
	fmt.Scan(&data[*n].jenis)
	for data[*n].jenis == "" {
		fmt.Println("Jenis produk kosong! Mohon masukkan jenis produk: ")
		fmt.Print(">>>> ")
		fmt.Scan(&data[*n].jenis)
	}
	fmt.Println("---------------------------")
	fmt.Println("Harga Produk:")
	fmt.Print(">>>> Rp.")
	fmt.Scan(&data[*n].harga)
	for data[*n].harga == 0 {
		fmt.Println("Harga produk kosong! Mohon masukkan harga produk: ")
		fmt.Print(">>>> Rp.")
		fmt.Scan(&data[*n].harga)
	}
	fmt.Println("---------------------------")
	fmt.Println("Stok Produk:")
	fmt.Print(">>>> ")
	fmt.Scan(&data[*n].stok)
	for data[*n].stok == -1 {
		fmt.Println("Stok produk kosong! Mohon masukkan stok produk: ")
		fmt.Print(">>>> ")
		fmt.Scan(&data[*n].stok)
	}
	fmt.Println("---------------------------")
	fmt.Println("Data Ke-", *n+1, "Berhasil Diinput")
	*n++
}

//TODO: Buat Prosedur dan Function untuk Sub-Menu dari Konfigurasi Data Produk

func tampilSemuaDataProduk(data Data, n int) {
	var det int
	menuHeaderKonfigurasiDataProduk()
	for det != 6 {
		menuOptionsTampilSemuaProduk()
		fmt.Println("Masukkan menu: ")
		fmt.Print(">>>> ")
		if det == 1 {
			showTampilSemuaProduk(data, n)
		} else if det == 2 {
			editTampilSemuaProduk(&data)
		}
	}
}

// Tampilkan data
// TODO: need a better placement for the interface header
// TODO_2: is calling tampilSemuaDataProduk() is necessary?
func showTampilSemuaProduk(data Data, n int) {
	fmt.Println("|------------------------------------------------------------------------------|")
	fmt.Println("|                      T A B E L - D A T A - P R O D U K                       |")
	fmt.Println("|------------------------------------------------------------------------------|")
	fmt.Println("| No. | Nama Produk | Merek Produk | Jenis Produk | Harga Produk | Stok Produk |")
	for i := 0; i < n; i++ {
		fmt.Printf("| %d | %s | %s | %s | %.2f | %d |\n", i+1, data[i].namaProduk, data[i].merek, data[i].jenis, data[i].harga, data[i].stok)
	}
	fmt.Println("|------------------------------------------------------------------------------|")
	fmt.Println()
	tampilSemuaDataProduk(data, n)
}

// Edit Data
/*
 * TODO: asumsi user mengira index 0 dari data adalah Data pertama
 * atau 1, akan akan aneh jika user mengedit data 1 dan yang teredit
 * menjadi data kedua.
 * TODO_2: asumsi user adalah tinkerer dan mencoba untuk input 0 untuk
 * select kolom data, kita buat user untuk input ulang dengan benar.
 * extra untuk mengecek apakah user meng-select kolom data lebih dari
 * MAXPRODUCT.
 * TODO_3: do we need to use by-pointers for the variables?
 */
func editTampilSemuaProduk(data *Data) {
	var n int
	fmt.Print("Masukkan kolom data yang akan diedit: ")
	fmt.Scan(&n)
	if n == 0 && n > MAXPRODUCT {
		fmt.Println("Masukkan kolom data yang benar!")
		fmt.Print("Masukkan kolom data yang akan diedit: ")
		fmt.Scan(&n)
	} else {
		fmt.Println("Nama Produk:")
		fmt.Print(">>>> ")
		fmt.Scan(&data[n-1].namaProduk)
		for data[n-1].namaProduk == "" {
			fmt.Println("Nama produk kosong! Mohon masukkan nama produk: ")
			fmt.Print(">>>> ")
			fmt.Scan(&data[n-1].namaProduk)
		}
		fmt.Println("---------------------------")
		fmt.Println("Merek Produk:")
		fmt.Print(">>>>")
		fmt.Scan(&data[n-1].merek)
		for data[n-1].merek == "" {
			fmt.Println("Merek produk kosong! Mohon masukkan merek produk: ")
			fmt.Print(">>>> ")
			fmt.Scan(&data[n-1].merek)
		}
		fmt.Println("---------------------------")
		fmt.Println("Jenis Produk:")
		fmt.Print(">>>> ")
		fmt.Scan(&data[n-1].jenis)
		for data[n-1].jenis == "" {
			fmt.Println("Jenis produk kosong! Mohon masukkan jenis produk: ")
			fmt.Print(">>>> ")
			fmt.Scan(&data[n-1].jenis)
		}
		fmt.Println("---------------------------")
		fmt.Println("Harga Produk:")
		fmt.Print(">>>> Rp.")
		fmt.Scan(&data[n-1].harga)
		for data[n-1].harga == 0 {
			fmt.Println("Harga produk kosong! Mohon masukkan harga produk: ")
			fmt.Print(">>>> Rp.")
			fmt.Scan(&data[n-1].harga)
		}
		fmt.Println("---------------------------")
		fmt.Println("Stok Produk:")
		fmt.Print(">>>> ")
		fmt.Scan(&data[n-1].stok)
		for data[n-1].stok == -1 {
			fmt.Println("Stok produk kosong! Mohon masukkan stok produk: ")
			fmt.Print(">>>> ")
			fmt.Scan(&data[n-1].stok)
		}
		fmt.Println("---------------------------")
		fmt.Println("Data Ke-", n, "Berhasil Diinput")
	}
}

// -----> Menampilkan menu secara estetik pada CLI <----------
func menuStart() {
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("-----------------------------------------------")
	fmt.Println("                 STORAGE SORTER                ")
	fmt.Println("-----------------------------------------------")
	fmt.Println(" Aplikasi Manajemen Inventaris Toko Elektronik ")
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")
	fmt.Println("                 Developed By:                 ")
	fmt.Println("                  Kelompok 1                   ")
	fmt.Println("-----------------------------------------------")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
}

func menuProcess() {
	fmt.Println("|-----------------------------------------------|")
	fmt.Println("|               M E N U   F I T U R             |")
	fmt.Println("|-----------------------------------------------|")
	fmt.Println("1. Konfigurasi Data Produk")
	fmt.Println("2. Pencatatan Transaksi")
	fmt.Println("3. Exit Program")
	fmt.Println("|-----------------------------------------------|")
}

func menuEnd() {
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("-----------------------------------------------")
	fmt.Println("                Program Selesai                ")
	fmt.Println("-----------------------------------------------")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
}

func menuOptionsTampilSemuaProduk() {
	menuHeaderKonfigurasiDataProduk()
	fmt.Println("1. Lihat data")
	fmt.Println("2. Edit Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Cari Data")
	fmt.Println("5. Urutkan Data")
}

func menuHeaderKonfigurasiDataProduk() {
	fmt.Println("|---------------------------------|")
	fmt.Println("| M E N U - P R O D U K - D A T A |")
	fmt.Println("|---------------------------------|")
}

func menuOptionsKonfigurasiDataProduk() {
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Tampilkan Semua Data")
	fmt.Println("3. Kembali ke Menu Utama")
}

// ---------------------------------------------------------------
