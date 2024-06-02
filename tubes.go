package main

import (
	"fmt"
)

// deklarasi akun user yang bisa memakai apk
const MAXAKUN int = 100

// deklarasi transaksi maksimal per akun
const MAXTRANSAKSI int = 200

// struct yang ada di total akun
type Total_Akun struct {
	id, username, alamat_email, password string
	saldo                                float64
	transaksi                            [MAXTRANSAKSI]T_transaksi
	transaksiCount                       int
	status                               bool
}

type T_transaksi struct {
	tanggal, bulan, tahun int
	uang                  float64
	kegiatan, catatan     string
}

// tabAkun adalah alias array akun dengan maks elemen MAXAKUN
type tabAkun [MAXAKUN]Total_Akun

// variabel global untuk menjalankan perintah
var tAkun tabAkun
var n int
var currentUserIdx int

func main() {
	menu_pertama()
}

//===================================AKSES ADMIN======================================/

// untuk admin agar bisa mengecek akun yang ingin di setujui
func cek_registrasi(A *tabAkun, n *int) {
	var pilihan int
	fmt.Println("=========================================")
	fmt.Println("1. melihat akun-akun yang ada")
	fmt.Println("2. memilih akun akun yang akan disetujui")
	fmt.Println("3. memilih akun akun yang akan ditolak")
	fmt.Println("4. menghapus akun")
	fmt.Println("5. exit")
	fmt.Println("=========================================")
	fmt.Print("pilih antara (1/2/3/4/5): ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		lihat_regis(A, n)
	case 2:
		persetujuan(A, n)
	case 3:
		penolakan(A, n)
	case 4:
		menghapus_akun(A, n)
	case 5:
		keluar_admin()
	default:
		fmt.Println("pilihan yang anda masukkan tidak ada di opsi")
		cek_registrasi(A, n)
	}
}

// admin ingin melihat akun akun yang belum di setujui
func lihat_regis(A *tabAkun, n *int) {
	var pilihan int
	fmt.Println("==================================================================================")
	fmt.Printf("%6s %6s %6s %6s %6s\n", "====ID====", "====username====", "=======email=======", "====password====", "====status====")
	for i := 0; i < *n; i++ {
		T_akun := (*A)[i]
		fmt.Printf("%7s %13s %22s %14s %14t\n", T_akun.id, T_akun.username, T_akun.alamat_email, T_akun.password, T_akun.status)
	}
	fmt.Println("==================================================================================")
	fmt.Println("1. exit")
	fmt.Print("pilih antara (1): ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		cek_registrasi(A, n)
	} else {
		fmt.Println("pilihan yang anda masukkan tidak ada di opsi")
		lihat_regis(A, n)
	}
}

// admin ingin menghapus akun yang diberikan persetujuan
func menghapus_akun(A *tabAkun, n *int) {
	var id string
	fmt.Print("tulis ID akun yang ingin kamu hapus: ")
	fmt.Scan(&id)

	idx := pengecekan(A, *n, id)
	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			(*A)[i] = (*A)[i+1]
		}
		*n--
		fmt.Println("Akun berhasil dihapus")
	} else {
		fmt.Println("Akun tidak ditemukan")
	}
	cek_registrasi(A, n)
}

// Fungsi untuk melakukan binary search pada array tabAkun yang sudah terurut berdasarkan id
func pengecekan(A *tabAkun, n int, id string) int {
	low := 0
	high := n - 1

	for low <= high {
		mid := (low + high) / 2
		if (*A)[mid].id == id {
			return mid
		} else if (*A)[mid].id < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// admin ingin keluar dari menu admin miliki
func keluar_admin() {
	menu_pertama()
}

//===================================AKSES USER DAN ADMIN===================================/

// untuk melakukan pemilihan antara login dan registrasi
func menu_pertama() {
	var pilihan int
	fmt.Println("================")
	fmt.Println("1. login")
	fmt.Println("2. sign up")
	fmt.Println("================")
	fmt.Print("pilih antara (1/2): ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		login(&tAkun, &n)
	case 2:
		registrasi(&tAkun, &n)
	default:
		fmt.Println("pilihan yang anda masukkan tidak ada di opsi")
		menu_pertama()
	}
}

// procedure untuk melakukan registrasi akun
func registrasi(A *tabAkun, n *int) {
	var id, username, alamat_email, password string
	fmt.Print("masukkan ID anda: ")
	fmt.Scan(&id)
	fmt.Println("CATATAN: penggunaan spasi digantikan oleh uderscore '_'")
	fmt.Print("masukkan username anda: ")
	fmt.Scan(&username)
	fmt.Print("masukkan alamat email anda: ")
	fmt.Scan(&alamat_email)
	fmt.Print("masukkan password anda: ")
	fmt.Scan(&password)

	if id == "000" {
		fmt.Println("ID yang anda masukkan tidak valid")
		registrasi(A, n)
		return
	}

	if cek_email(*A, *n, alamat_email) == 1 && cek_id(*A, *n, id) == 1 {
		A[*n].id = id
		A[*n].username = username
		A[*n].alamat_email = alamat_email
		A[*n].password = password
		*n++

	} else if cek_id(*A, *n, id) == -1 && cek_email(*A, *n, alamat_email) == -1 {
		fmt.Println("id dan email yang di masukkan sudah terdaftar")
		fmt.Println("mohon masukkan lagi ID dan email yang belum ada")
	} else if cek_email(*A, *n, alamat_email) == -1 {
		fmt.Println("email yang di masukkan sudah terdaftar")
		fmt.Println("mohon masukkan lagi email yang belum")
	} else if cek_id(*A, *n, id) == -1 {
		fmt.Println("ID yang di masukkan sudah terdaftar")
		fmt.Println("mohon masukkan lagi ID yang belum ada")
	}
	menu_pertama()
}

// procedure untuk melakukan login akun yang sudah di setujui oleh admin
func login(A *tabAkun, n *int) {
	var id, username, alamat_email, password string
	fmt.Print("masukkan ID anda: ")
	fmt.Scan(&id)
	fmt.Print("masukkan username anda: ")
	fmt.Scan(&username)
	fmt.Print("masukkan alamat email anda: ")
	fmt.Scan(&alamat_email)
	fmt.Print("masukkan password anda: ")
	fmt.Scan(&password)

	if id == "000" && username == "root" && alamat_email == "root" && password == "root" {
		cek_registrasi(A, n)
	} else {
		idx := cek_login(*A, *n, id)
		if idx != -1 && A[idx].status && A[idx].id == id && A[idx].username == username && A[idx].alamat_email == alamat_email && A[idx].password == password {
			currentUserIdx = idx
			menu_kedua()
		} else {
			fmt.Println("akses anda di tolak")
			menu_pertama()
		}
	}
}

// pengecekan apakah id yang di masukkan user ada sudah di setujui atau belum
func cek_login(A tabAkun, n int, id string) int {
	for i := 0; i < n; i++ {
		if id == A[i].id {
			return i
		}
	}
	return -1
}

func cek_email(A tabAkun, n int, alamat_email string) int {
	for i := 0; i < n; i++ {
		if A[i].alamat_email == alamat_email {
			return -1
		}
	}
	return 1
}

func cek_id(A tabAkun, n int, id string) int {
	for i := 0; i < n; i++ {
		if A[i].id == id {
			return -1
		}
	}
	return 1
}

//===================================AKSES USER===================================/

// setujui akun
func persetujuan(A *tabAkun, n *int) {
	var id string
	fmt.Print("tulis ID akun yang ingin kamu setujui: ")
	fmt.Scan(&id)

	for i := 0; i < *n; i++ {
		if id == (*A)[i].id {
			if (*A)[i].status {
				fmt.Println("Akun tersebut sudah disetujui")
			} else {
				(*A)[i].status = true
				(*A)[i].saldo = 100000
				fmt.Println("Akun berhasil disetujui")
			}
		}
	}

	cek_registrasi(A, n)
}

// tolak akun
func penolakan(A *tabAkun, n *int) {
	var id string
	fmt.Print("tulis ID akun yang ingin kamu tolak: ")
	fmt.Scan(&id)

	for i := 0; i < *n; i++ {
		if id == (*A)[i].id {
			(*A)[i].status = false
		}
	}
	cek_registrasi(A, n)
}

// menu utama untuk user
func menu_kedua() {
	var A tabAkun
	var pilihan int
	fmt.Println("===========================================")
	fmt.Println("1. ketik 1 jika ingin cek saldo")
	fmt.Println("2. ketik 2 jika ingin mengirim uang")
	fmt.Println("3. ketik 3 jika ingin melakukan pembayaran")
	fmt.Println("4. ketik 4 jika ingin melihat riwayat transaksi")
	fmt.Println("5. logout")
	fmt.Println("===========================================")
	fmt.Print("pilih antara (1/2/3/4/5): ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		cek_saldo()
	case 2:
		kirim_uang(&A)
	case 3:
		pembayaran(&A)
	case 4:
		riwayat_transaksi()
	case 5:
		menu_pertama()
	default:
		menu_kedua()
	}
}

// berfungsi agar user bisa melihat saldo yang dimiliki
func cek_saldo() {
	var pilihan int
	fmt.Printf("Saldo Anda: %.2f\n", tAkun[currentUserIdx].saldo)
	fmt.Println("1. exit")
	fmt.Print("pilih antara (1): ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		menu_kedua()
	} else {
		fmt.Println("pilihan yang anda masukkan tidak ada di opsi")
		cek_saldo()
	}
}

// berfungsi agar user bisa mengirim uang kepada rekening lain
func kirim_uang(A *tabAkun) {
	var pilihan int
	var kirim float64
	var id string
	var tanggal, bulan, tahun int
	var catatan string

	fmt.Print("\nmasukkan ID yang ingin kamu kirimkan uang kamu: ")
	fmt.Scan(&id)

	idx := cek_login(tAkun, n, id)
	if idx == -1 || tAkun[idx].status == false || id == tAkun[currentUserIdx].id {
		fmt.Println("ID yang kamu masukkan tidak valid")
		menu_kedua()
		return
	}

	fmt.Print("masukkan jumlah uang yang ingin kamu kirimkan: ")
	fmt.Scan(&kirim)

	if tAkun[currentUserIdx].saldo < kirim {
		fmt.Println("Saldo tidak mencukupi")
		menu_kedua()
	} else if kirim < 0 {
		fmt.Println("tidak bisa mengirim kurang dari 0")
		menu_kedua()
	}

	fmt.Println("CATATAN: contoh penulisan tanggal 25 11 2005")
	fmt.Print("masukkan tanggal sekarang DD MM YYYY: ")
	fmt.Scan(&tanggal, &bulan, &tahun)
	if tanggal < 0 || tanggal > 32 || bulan < 0 || bulan > 13 || tahun < 2000 || tahun > 2025 {
		fmt.Println("tanggal yang anda masukkan tidak terdefinisi")
		fmt.Println("mohon mengulang kembali dengan benar")
		kirim_uang(A)
		return
	}

	fmt.Println("CATATAN: jika ingin menggukan spasi, gunakanlah underscores seperti _")
	fmt.Print("masukkan catatan: ")
	fmt.Scan(&catatan)

	// orang yang meneransfer uang
	tAkun[currentUserIdx].transaksi[tAkun[currentUserIdx].transaksiCount] = T_transaksi{
		tanggal:  tanggal,
		bulan:    bulan,
		tahun:    tahun,
		uang:     kirim,
		kegiatan: "melakukan_transfer",
		catatan:  catatan,
	}
	tAkun[currentUserIdx].transaksiCount++

	// orang yang menerima uang
	tAkun[idx].transaksi[tAkun[idx].transaksiCount] = T_transaksi{
		tanggal:  tanggal,
		bulan:    bulan,
		tahun:    tahun,
		uang:     kirim,
		kegiatan: "menerima_transfer",
		catatan:  catatan,
	}
	tAkun[idx].transaksiCount++

	tAkun[currentUserIdx].saldo -= kirim
	menerima_uang(kirim, idx)
	fmt.Println("Transaksi berhasil")
	fmt.Println("1. exit")
	fmt.Print("pilih antara (1): ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		menu_kedua()
	} else {
		fmt.Println("pilihan yang anda masukkan tidak ada di opsi")
		fmt.Println("anda akan otomatis dikeluarkan")
		menu_kedua()
	}
}

// berfungsi menambah uang kepada akun yang diberikan uang
func menerima_uang(menerima float64, idx int) {
	tAkun[idx].saldo += menerima
}

// berfungsi agar user bisa membayar pembelian
func pembayaran(A *tabAkun) {
	var pilihan int
	var bayar float64
	var tanggal, bulan, tahun int
	var catatan string

	fmt.Println("===========================================")
	fmt.Println("1. ketik 1 jika ingin membayar makanan")
	fmt.Println("2. ketik 2 jika ingin membayar pulsa")
	fmt.Println("3. ketik 3 jika ingin membayar listrik")
	fmt.Println("4. ketik 4 jika ingin membayar BPJS")
	fmt.Println("5. ketik 5 jika ingin membayar lainnya")
	fmt.Println("6. logout")
	fmt.Println("===========================================")
	fmt.Print("pilih antara (1/2/3/4/5/6): ")
	fmt.Scan(&pilihan)

	if pilihan == 6 {
		menu_kedua()
		return
	} else if pilihan == 1 || pilihan == 2 || pilihan == 3 || pilihan == 4 || pilihan == 5 {
		fmt.Print("\nmasukkan jumlah uang yang ingin kamu bayar: ")
		fmt.Scan(&bayar)

		if tAkun[currentUserIdx].saldo < bayar {
			fmt.Println("Saldo tidak mencukupi")
			menu_kedua()
		} else if bayar < 0 {
			fmt.Println("tidak bisa mengirim kurang dari 0")
			menu_kedua()
		}

		fmt.Println("CATATAN: contoh penulisan tanggal 25 11 2005")
		fmt.Print("masukkan tanggal sekarang DD MM YYYY: ")
		fmt.Scan(&tanggal, &bulan, &tahun)
		if tanggal < 1 || tanggal > 31 || bulan < 1 || bulan > 12 || tahun < 2000 || tahun > 2024 {
			fmt.Println("tanggal yang anda masukkan tidak terdefinisi")
			fmt.Println("mohon mengulang kembali dengan benar")
			pembayaran(A)
			return
		}

		fmt.Println("CATATAN: jika ingin menggukan spasi, gunakanlah underscores seperti _")
		fmt.Print("masukkan catatan: ")
		fmt.Scan(&catatan)

		if pilihan == 1 {
			tAkun[currentUserIdx].transaksi[tAkun[currentUserIdx].transaksiCount] = T_transaksi{
				tanggal:  tanggal,
				bulan:    bulan,
				tahun:    tahun,
				uang:     bayar,
				kegiatan: "membayar_makanan",
				catatan:  catatan,
			}
			tAkun[currentUserIdx].transaksiCount++
		}
		if pilihan == 2 {
			tAkun[currentUserIdx].transaksi[tAkun[currentUserIdx].transaksiCount] = T_transaksi{
				tanggal:  tanggal,
				bulan:    bulan,
				tahun:    tahun,
				uang:     bayar,
				kegiatan: "membayar_pulsa",
				catatan:  catatan,
			}
			tAkun[currentUserIdx].transaksiCount++
		}
		if pilihan == 3 {
			tAkun[currentUserIdx].transaksi[tAkun[currentUserIdx].transaksiCount] = T_transaksi{
				tanggal:  tanggal,
				bulan:    bulan,
				tahun:    tahun,
				uang:     bayar,
				kegiatan: "membayar_listrik",
				catatan:  catatan,
			}
			tAkun[currentUserIdx].transaksiCount++
		}
		if pilihan == 4 {
			tAkun[currentUserIdx].transaksi[tAkun[currentUserIdx].transaksiCount] = T_transaksi{
				tanggal:  tanggal,
				bulan:    bulan,
				tahun:    tahun,
				uang:     bayar,
				kegiatan: "membayar_BPJS",
				catatan:  catatan,
			}
			tAkun[currentUserIdx].transaksiCount++
		}
		if pilihan == 5 {
			tAkun[currentUserIdx].transaksi[tAkun[currentUserIdx].transaksiCount] = T_transaksi{
				tanggal:  tanggal,
				bulan:    bulan,
				tahun:    tahun,
				uang:     bayar,
				kegiatan: "lainnya",
				catatan:  catatan,
			}
			tAkun[currentUserIdx].transaksiCount++
		}

		tAkun[currentUserIdx].saldo -= bayar
		fmt.Println("Pembayaran berhasil")
		menu_kedua()
	}
}

// berfungsi agar user bisa melihat transaksi apa saja yang dilakukan oleh akun tersebut
func riwayat_transaksi() {
	var memilih int
	fmt.Print("1. urut terbaru\n")
	fmt.Print("2. urut terlama\n")
	fmt.Print("pilih antara (1/2): ")
	fmt.Scan(&memilih)

	if memilih == 1 {
		urut_tanggal_transaksi_terbaru()
	} else if memilih == 2 {
		urut_tanggal_transaksi_terlama()
	}
	for i := 0; i < tAkun[currentUserIdx].transaksiCount; i++ {
		trans := tAkun[currentUserIdx].transaksi[i]
		fmt.Printf("Tanggal: %02d/%02d/%d\n", trans.tanggal, trans.bulan, trans.tahun)
		fmt.Printf("Catatan: %s\n", trans.catatan)
		fmt.Printf("uang: %.2f\n", trans.uang)
		fmt.Printf("Kegiatan: %s\n", trans.kegiatan)
		fmt.Println("------------------------------------")
	}
	fmt.Println("1. exit")
	var pilihan int
	fmt.Print("pilih antara (1): ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		menu_kedua()
	} else {
		fmt.Println("pilihan yang anda masukkan tidak ada di opsi")
		riwayat_transaksi()
	}
}

// untuk mengurutkan tanggal dari paling besar hingga paling kecil dengan selection sort
func urut_tanggal_transaksi_terbaru() {
	for i := 0; i < tAkun[currentUserIdx].transaksiCount-1; i++ {
		maxIdx := i
		for j := i + 1; j < tAkun[currentUserIdx].transaksiCount; j++ {
			if compareDate(tAkun[currentUserIdx].transaksi[j], tAkun[currentUserIdx].transaksi[maxIdx]) {
				maxIdx = j
			}
		}
		temp := tAkun[currentUserIdx].transaksi[maxIdx]
		tAkun[currentUserIdx].transaksi[maxIdx] = tAkun[currentUserIdx].transaksi[i]
		tAkun[currentUserIdx].transaksi[i] = temp
	}
}

// untuk mengurutkan tanggal dari paling besar hingga paling kecil dengan insersion sort
func urut_tanggal_transaksi_terlama() {
	for i := 1; i < tAkun[currentUserIdx].transaksiCount; i++ {
		key := tAkun[currentUserIdx].transaksi[i]
		j := i - 1
		for j >= 0 && !compareDate(key, tAkun[currentUserIdx].transaksi[j]) {
			tAkun[currentUserIdx].transaksi[j+1] = tAkun[currentUserIdx].transaksi[j]
			j = j - 1
		}
		tAkun[currentUserIdx].transaksi[j+1] = key
	}

}

func compareDate(t1, t2 T_transaksi) bool {
	if t1.tahun > t2.tahun {
		return true
	}
	if t1.tahun < t2.tahun {
		return false
	}
	if t1.bulan > t2.bulan {
		return true
	}
	if t1.bulan < t2.bulan {
		return false
	}
	return t1.tanggal > t2.tanggal
}

// berfungsi agar user bisa keluar dari menu pertama
func logout() {
	menu_pertama()
}