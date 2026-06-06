package main

import "fmt"

type Peserta struct {
	idPendaftaran int
	namaLengkap   string
	tanggalDaftar string
	idKursus      int
	idBidang      int
	statusAktif   bool
}

type Kursus struct {
	idKursus   int
	namaKursus string
	idBidang   int
}

type BidangMinat struct {
	idBidang   int
	namaBidang string
}

const NMAX = 1000

type tabPeserta [NMAX]Peserta
type tabKursus [NMAX]Kursus
type tabBidang [NMAX]BidangMinat

func cekIdPeserta(p tabPeserta, n int, id int) bool {
	var i int

	for i = 0; i < n; i++ {
		if p[i].idPendaftaran == id {
			return true
		}
	}

	return false
}

func cekIdKursus(k tabKursus, n int, id int) bool {
	var i int

	for i = 0; i < n; i++ {
		if k[i].idKursus == id {
			return true
		}
	}

	return false
}

func cekIdBidang(b tabBidang, n int, id int) bool {
	var i int

	for i = 0; i < n; i++ {
		if b[i].idBidang == id {
			return true
		}
	}

	return false
}

func tambahBidangMinat(b *tabBidang, n *int) {
	var idBidang int

	if *n < NMAX {
		fmt.Println("\n=== Tambah Bidang Minat ===")

		fmt.Print("Masukkan ID Bidang   : ")
		fmt.Scan(&idBidang)

		if cekIdBidang(*b, *n, idBidang) {
			fmt.Println("ID bidang sudah digunakan.")
			return
		}

		fmt.Print("Masukkan Nama Bidang : ")
		fmt.Scan(&b[*n].namaBidang)

		b[*n].idBidang = idBidang

		*n = *n + 1
		fmt.Println("Data bidang minat berhasil ditambahkan.")
	} else {
		fmt.Println("Data bidang minat sudah penuh.")
	}
}

func tampilBidangMinat(b tabBidang, n int) {
	var i int

	if n == 0 {
		fmt.Println("Belum ada data bidang minat.")
	} else {
		fmt.Println("\n=== DATA BIDANG MINAT ===")
		for i = 0; i < n; i++ {
			fmt.Println("ID Bidang   :", b[i].idBidang)
			fmt.Println("Nama Bidang :", b[i].namaBidang)
			fmt.Println("-------------------------")
		}
	}
}

func tambahDataKursus(k *tabKursus, n *int, b tabBidang, nBidang int) {
	var idKursus, idBidang int

	if nBidang == 0 {
		fmt.Println("Data bidang belum ditambahkan. Silahkan tambahkan terlebih dahulu.")
		return
	}

	if *n < NMAX {
		fmt.Println("\n=== Tambah Data Kursus ===")

		fmt.Print("Masukkan ID Kursus    : ")
		fmt.Scan(&idKursus)

		if cekIdKursus(*k, *n, idKursus) {
			fmt.Println("ID kursus sudah digunakan.")
			return
		}

		fmt.Print("Masukkan Nama Kursus  : ")
		fmt.Scan(&k[*n].namaKursus)

		fmt.Print("Masukkan ID Bidang    : ")
		fmt.Scan(&idBidang)

		if !cekIdBidang(b, nBidang, idBidang) {
			fmt.Println("ID bidang tidak ditemukan.")
			return
		}

		k[*n].idKursus = idKursus
		k[*n].idBidang = idBidang

		*n = *n + 1
		fmt.Println("Data kursus berhasil ditambahkan.")
	} else {
		fmt.Println("Data kursus sudah penuh.")
	}
}

func tampilDataKursus(k tabKursus, n int) {
	var i int

	if n == 0 {
		fmt.Println("Belum ada data kursus.")
	} else {
		fmt.Println("\n=== DATA KURSUS ===")
		for i = 0; i < n; i++ {
			fmt.Println("ID Kursus   :", k[i].idKursus)
			fmt.Println("Nama Kursus :", k[i].namaKursus)
			fmt.Println("ID Bidang   :", k[i].idBidang)
			fmt.Println("-------------------------")
		}
	}
}

func ambilIdBidangDariKursus(k tabKursus, nKursus int, idKursus int) int {
	var i int

	for i = 0; i < nKursus; i++ {
		if k[i].idKursus == idKursus {
			return k[i].idBidang
		}
	}

	return -1
}

func tambahDataPeserta(p *tabPeserta, n *int, k tabKursus, nKursus int, b tabBidang, nBidang int) {
	var idPendaftaran, idKursus, idBidang int

	if nKursus == 0 || nBidang == 0 {
		fmt.Println("Data kursus atau bidang belum ditambahkan. Silahkan tambahkan terlebih dahulu.")
		return
	}

	if *n < NMAX {
		fmt.Println("\n========== Tambah Data Peserta ==========")

		fmt.Printf("Masukkan ID Pendaftaran		 : ")
		fmt.Scan(&idPendaftaran)

		if cekIdPeserta(*p, *n, idPendaftaran) {
			fmt.Println("ID pendaftaran sudah digunakan.")
			return
		}

		fmt.Printf("Masukkan Nama Lengkap		 : ")
		fmt.Scan(&p[*n].namaLengkap)

		fmt.Printf("Masukkan Tanggal Daftar		 : ")
		fmt.Scan(&p[*n].tanggalDaftar)

		fmt.Printf("Masukkan ID Kursus		 : ")
		fmt.Scan(&idKursus)

		if !cekIdKursus(k, nKursus, idKursus) {
			fmt.Println("ID kursus tidak ditemukan.")
			return
		}

		idBidang = ambilIdBidangDariKursus(k, nKursus, idKursus)
		fmt.Printf("ID Bidang (otomatis dari kursus) : %d\n", idBidang)

		p[*n].idPendaftaran = idPendaftaran
		p[*n].idKursus = idKursus
		p[*n].idBidang = idBidang

		fmt.Printf("Status Aktif true/false		 : ")
		fmt.Scan(&p[*n].statusAktif)

		*n = *n + 1
		fmt.Println("Data peserta berhasil ditambahkan.")
	} else {
		fmt.Println("Data peserta sudah penuh.")
	}
}

func tampilDataPeserta(p tabPeserta, n int) {
	var i int

	if n == 0 {
		fmt.Println("Belum ada data peserta.")
	} else {
		fmt.Println("\n=== DATA PESERTA ===")
		for i = 0; i < n; i++ {
			fmt.Println("ID Pendaftaran :", p[i].idPendaftaran)
			fmt.Println("Nama Lengkap   :", p[i].namaLengkap)
			fmt.Println("Tanggal Daftar :", p[i].tanggalDaftar)
			fmt.Println("ID Kursus      :", p[i].idKursus)
			fmt.Println("ID Bidang      :", p[i].idBidang)
			fmt.Println("Status Aktif   :", p[i].statusAktif)
			fmt.Println("-------------------------")
		}
	}
}

func ubahDataPeserta(p *tabPeserta, n int, k tabKursus, nKursus int, b tabBidang, nBidang int) {
	var cari, i, pilih int
	var idBaru int
	var ketemu, selesai bool

	if n == 0 {
		fmt.Println("Data peserta masih kosong. Silahkan tambah data terlebih dahulu.")
		return
	}

	fmt.Print("Masukkan ID Pendaftaran yang ingin diubah: ")
	fmt.Scan(&cari)

	ketemu = false

	for i = 0; i < n && !ketemu; i++ {
		if p[i].idPendaftaran == cari {
			ketemu = true
			selesai = false

			for !selesai {
				fmt.Println("\n=== UBAH DATA PESERTA ===")
				fmt.Println("1. Ubah Nama Lengkap")
				fmt.Println("2. Ubah Tanggal Daftar")
				fmt.Println("3. Ubah ID Kursus")
				fmt.Println("4. Ubah Status Aktif")
				fmt.Println("5. Selesai")
				fmt.Print("Pilih data yang ingin diubah: ")
				fmt.Scan(&pilih)

				if pilih == 1 {
					fmt.Print("Masukkan nama baru: ")
					fmt.Scan(&p[i].namaLengkap)

				} else if pilih == 2 {
					fmt.Print("Masukkan tanggal daftar baru: ")
					fmt.Scan(&p[i].tanggalDaftar)

				} else if pilih == 3 {
					fmt.Print("Masukkan ID kursus baru: ")
					fmt.Scan(&idBaru)

					if !cekIdKursus(k, nKursus, idBaru) {
						fmt.Println("ID kursus tidak ditemukan")
					} else {
						p[i].idKursus = idBaru
						p[i].idBidang = ambilIdBidangDariKursus(k, nKursus, idBaru)
						fmt.Println("ID Bidang otomatis diperbarui menjadi:", p[i].idBidang)
					}

				} else if pilih == 4 {
					fmt.Print("Masukkan status aktif baru true/false: ")
					fmt.Scan(&p[i].statusAktif)

				} else if pilih == 5 {
					selesai = true

				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}

			fmt.Println("Data peserta berhasil diubah.")
		}
	}

	if !ketemu {
		fmt.Println("Data peserta tidak ditemukan.")
	}
}

func cariIDPeserta(p tabPeserta, n int, cari int) int {
	var i int

	for i = 0; i < n; i++ {
		if p[i].idPendaftaran == cari {
			return i
		}
	}

	return -1
}

func hapusDataPeserta(p *tabPeserta, n *int) {
	var cari, idx, i int

	if *n == 0 {
		fmt.Println("Data peserta masih kosong. Silahkan tambah data terlebih dahulu.")
		return
	}

	fmt.Print("Masukkan ID Pendaftaran yang ingin dihapus: ")
	fmt.Scan(&cari)

	idx = cariIDPeserta(*p, *n, cari)

	if idx == -1 {
		fmt.Println("Data peserta tidak ditemukan.")
	} else {
		for i = idx; i < *n-1; i++ {
			p[i] = p[i+1]
		}

		*n = *n - 1
		fmt.Println("Data peserta berhasil dihapus.")
	}
}

func cariPesertaBidang(p tabPeserta, n int) {
	var i, cariBidang int
	var ketemu bool

	if n == 0 {
		fmt.Println("Data peserta masih kosong. Silahkan tambah data terlebih dahulu.")
		return
	}

	fmt.Print("Masukkan ID Bidang yang dicari: ")
	fmt.Scan(&cariBidang)

	ketemu = false

	fmt.Println("\n=== HASIL PENCARIAN BIDANG MINAT ===")

	for i = 0; i < n; i++ {
		if p[i].idBidang == cariBidang {
			fmt.Println("ID Pendaftaran :", p[i].idPendaftaran)
			fmt.Println("Nama Lengkap   :", p[i].namaLengkap)
			fmt.Println("Tanggal Daftar :", p[i].tanggalDaftar)
			fmt.Println("ID Kursus      :", p[i].idKursus)
			fmt.Println("ID Bidang      :", p[i].idBidang)
			fmt.Println("Status Aktif   :", p[i].statusAktif)
			fmt.Println("-------------------------")

			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Data peserta dengan bidang tersebut tidak ditemukan.")
	}
}

func cariPesertaNama(p tabPeserta, n int) {
	var cari string
	var kiri, kanan, tengah int
	var i int
	var temp tabPeserta
	var ketemu bool

	if n == 0 {
		fmt.Println("Data peserta masih kosong.")
		return
	}

	fmt.Print("Masukkan nama lengkap yang dicari: ")
	fmt.Scan(&cari)

	for i = 0; i < n; i++ {
		temp[i] = p[i]
	}

	urutNamaPeserta(&temp, n)

	kiri = 0
	kanan = n - 1
	ketemu = false

	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2
		if temp[tengah].namaLengkap == cari {
			ketemu = true
		} else if temp[tengah].namaLengkap < cari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if ketemu {
		fmt.Println("\n=== DATA PESERTA DITEMUKAN ===")

		i = tengah
		for i >= 0 && temp[i].namaLengkap == cari {
			fmt.Println("ID Pendaftaran :", temp[i].idPendaftaran)
			fmt.Println("Nama Lengkap   :", temp[i].namaLengkap)
			fmt.Println("Tanggal Daftar :", temp[i].tanggalDaftar)
			fmt.Println("ID Kursus      :", temp[i].idKursus)
			fmt.Println("ID Bidang      :", temp[i].idBidang)
			fmt.Println("Status Aktif   :", temp[i].statusAktif)
			fmt.Println("-------------------------")
			i--
		}

		i = tengah + 1
		for i < n && temp[i].namaLengkap == cari {
			fmt.Println("ID Pendaftaran :", temp[i].idPendaftaran)
			fmt.Println("Nama Lengkap   :", temp[i].namaLengkap)
			fmt.Println("Tanggal Daftar :", temp[i].tanggalDaftar)
			fmt.Println("ID Kursus      :", temp[i].idKursus)
			fmt.Println("ID Bidang      :", temp[i].idBidang)
			fmt.Println("Status Aktif   :", temp[i].statusAktif)
			fmt.Println("-------------------------")
			i++
		}
	} else {
		fmt.Println("Data peserta tidak ditemukan.")
	}
}

func urutIDPeserta(p *tabPeserta, n int) {
	var pass, idx, i int
	var temp Peserta

	if n == 0 {
		fmt.Println("Data peserta masih kosong. Tidak ada data yang bisa diurutkan.")
		return
	}

	for pass = 1; pass <= n-1; pass++ {
		idx = pass - 1

		for i = pass; i < n; i++ {
			if p[i].idPendaftaran < p[idx].idPendaftaran {
				idx = i
			}
		}

		temp = p[pass-1]
		p[pass-1] = p[idx]
		p[idx] = temp
	}

	fmt.Println("Data peserta berhasil diurutkan berdasarkan ID Pendaftaran.")
}

func urutNamaPeserta(p *tabPeserta, n int) {
	var i, j int
	var temp Peserta

	if n == 0 {
		fmt.Println("Data peserta masih kosong. Tidak ada data yang bisa diurutkan.")
		return
	}

	for i = 1; i < n; i++ {
		temp = p[i]
		j = i - 1

		for j >= 0 && p[j].namaLengkap > temp.namaLengkap {
			p[j+1] = p[j]
			j = j - 1
		}

		p[j+1] = temp
	}
}

func statistikPeserta(p tabPeserta, nPeserta int, b tabBidang, nBidang int) {
	var i, j, jumlah, totalAktif int

	if nBidang == 0 || nPeserta == 0 {
		fmt.Println("Data peserta atau bidang minat masih kosong. Belum ada statistik yang bisa ditampilkan.")
		return
	}

	fmt.Println("\n=== STATISTIK PESERTA ===")

	for i = 0; i < nBidang; i++ {
		jumlah = 0

		for j = 0; j < nPeserta; j++ {
			if p[j].idBidang == b[i].idBidang {
				jumlah = jumlah + 1
			}
		}

		fmt.Println("Bidang Minat :", b[i].namaBidang)
		fmt.Println("Jumlah       :", jumlah)
		fmt.Println("-------------------------")
	}

	totalAktif = 0
	for i = 0; i < nPeserta; i++ {
		if p[i].statusAktif == true {
			totalAktif = totalAktif + 1
		}
	}

	fmt.Println("Total Peserta Aktif:", totalAktif)
}

func main() {
	var dataPeserta tabPeserta
	var dataKursus tabKursus
	var dataBidang tabBidang

	var nPeserta, nKursus, nBidang int
	var pilih int

	for pilih != 14 {
		fmt.Println("\n===== MENU KURSUSIN =====")
		fmt.Println("1. Tambah Data Peserta")
		fmt.Println("2. Ubah Data Peserta")
		fmt.Println("3. Hapus Data Peserta")
		fmt.Println("4. Tampil Data Peserta")
		fmt.Println("5. Tambah Data Kursus")
		fmt.Println("6. Tampil Data Kursus")
		fmt.Println("7. Tambah Bidang Minat")
		fmt.Println("8. Tampil Bidang Minat")
		fmt.Println("9. Cari Peserta Berdasarkan Bidang Minat")
		fmt.Println("10. Cari Peserta Berdasarkan Nama")
		fmt.Println("11. Urutkan Peserta Berdasarkan ID")
		fmt.Println("12. Urutkan Peserta Berdasarkan Nama")
		fmt.Println("13. Statistik Peserta")
		fmt.Println("14. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahDataPeserta(&dataPeserta, &nPeserta, dataKursus, nKursus, dataBidang, nBidang)

		} else if pilih == 2 {
			ubahDataPeserta(&dataPeserta, nPeserta, dataKursus, nKursus, dataBidang, nBidang)

		} else if pilih == 3 {
			hapusDataPeserta(&dataPeserta, &nPeserta)

		} else if pilih == 4 {
			tampilDataPeserta(dataPeserta, nPeserta)

		} else if pilih == 5 {
			tambahDataKursus(&dataKursus, &nKursus, dataBidang, nBidang)

		} else if pilih == 6 {
			tampilDataKursus(dataKursus, nKursus)

		} else if pilih == 7 {
			tambahBidangMinat(&dataBidang, &nBidang)

		} else if pilih == 8 {
			tampilBidangMinat(dataBidang, nBidang)

		} else if pilih == 9 {
			cariPesertaBidang(dataPeserta, nPeserta)

		} else if pilih == 10 {
			cariPesertaNama(dataPeserta, nPeserta)

		} else if pilih == 11 {
			urutIDPeserta(&dataPeserta, nPeserta)

		} else if pilih == 12 {
			urutNamaPeserta(&dataPeserta, nPeserta)
			fmt.Println("Data peserta berhasil diurutkan berdasarkan nama lengkap.")

		} else if pilih == 13 {
			statistikPeserta(dataPeserta, nPeserta, dataBidang, nBidang)

		} else if pilih == 14 {
			fmt.Println("Program selesai.")

		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
