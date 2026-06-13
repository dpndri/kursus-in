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

func validasiIDPositif(label string) int {
	var id int
	var valid bool

	valid = false
	for !valid {
		fmt.Print(label)
		fmt.Scan(&id)

		if id <= 0 {
			fmt.Println("[!] ID harus berupa bilangan positif (lebih dari 0).")
		} else {
			valid = true
		}
	}
	return id
}

func validasiNama(label string) string {
	var nama string
	var valid bool
	var i int
	var adaHuruf bool

	valid = false
	for !valid {
		fmt.Print(label)
		fmt.Scan(&nama)

		adaHuruf = false
		for i = 0; i < len(nama); i++ {
			if (nama[i] >= 'a' && nama[i] <= 'z') || (nama[i] >= 'A' && nama[i] <= 'Z') {
				adaHuruf = true
			}
		}

		if len(nama) == 0 {
			fmt.Println("[!] Nama tidak boleh kosong.")
		} else if !adaHuruf {
			fmt.Println("[!] Nama harus mengandung huruf.")
		} else {
			valid = true
		}
	}
	return nama
}

func validasiTanggal(label string) string {
	var tgl string
	var valid bool
	var i int
	var semuaValid bool

	valid = false
	for !valid {
		fmt.Print(label)
		fmt.Scan(&tgl)

		if len(tgl) != 10 {
			fmt.Println("[!] Format tanggal tidak valid. Gunakan format DD-MM-YYYY (contoh: 01-01-2026).")
		} else if tgl[2] != '-' || tgl[5] != '-' {
			fmt.Println("[!] Format tanggal tidak valid. Gunakan tanda '-' sebagai pemisah (contoh: 01-01-2026).")
		} else {
			semuaValid = true

			for i = 0; i < 2; i++ {
				if tgl[i] < '0' || tgl[i] > '9' {
					semuaValid = false
				}
			}
			for i = 3; i < 5; i++ {
				if tgl[i] < '0' || tgl[i] > '9' {
					semuaValid = false
				}
			}
			for i = 6; i < 10; i++ {
				if tgl[i] < '0' || tgl[i] > '9' {
					semuaValid = false
				}
			}

			if !semuaValid {
				fmt.Println("[!] Tanggal hanya boleh berisi angka dan tanda '-' (contoh: 01-01-2026).")
			} else {
				valid = true
			}
		}
	}
	return tgl
}

func validasiStatusAktif(label string) bool {
	var input string
	var valid bool

	valid = false
	for !valid {
		fmt.Print(label)
		fmt.Scan(&input)

		if input == "true" {
			return true
		} else if input == "false" {
			return false
		} else {
			fmt.Println("[!] Input tidak valid. Masukkan 'true' atau 'false'.")
		}
	}
	return false
}

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

	if *n >= NMAX {
		fmt.Println("[!] Data bidang minat sudah penuh.")
		return
	}

	fmt.Println("\n--- Tambah Bidang Minat ---")

	idBidang = validasiIDPositif("ID Bidang   : ")
	if cekIdBidang(*b, *n, idBidang) {
		fmt.Println("[!] ID bidang sudah digunakan.")
		return
	}

	b[*n].namaBidang = validasiNama("Nama Bidang : ")
	b[*n].idBidang = idBidang

	*n = *n + 1
	fmt.Println("[v] Data bidang minat berhasil ditambahkan.")
}

func tampilBidangMinat(b tabBidang, n int) {
	var i int

	if n == 0 {
		fmt.Println("[!] Belum ada data bidang minat.")
		return
	}

	fmt.Println("\n======= DATA BIDANG MINAT =======")
	for i = 0; i < n; i++ {
		fmt.Printf("  %-14s : %d\n", "ID Bidang", b[i].idBidang)
		fmt.Printf("  %-14s : %s\n", "Nama Bidang", b[i].namaBidang)
		fmt.Println("---------------------------------")
	}
}

func tambahDataKursus(k *tabKursus, n *int, b tabBidang, nBidang int) {
	var idKursus, idBidang int

	if nBidang == 0 {
		fmt.Println("[!] Data bidang belum ada. Silakan tambah bidang minat terlebih dahulu.")
		return
	}

	if *n >= NMAX {
		fmt.Println("[!] Data kursus sudah penuh.")
		return
	}

	fmt.Println("\n--- Tambah Data Kursus ---")

	idKursus = validasiIDPositif("ID Kursus   : ")
	if cekIdKursus(*k, *n, idKursus) {
		fmt.Println("[!] ID kursus sudah digunakan.")
		return
	}

	k[*n].namaKursus = validasiNama("Nama Kursus : ")
	idBidang = validasiIDPositif("ID Bidang   : ")

	if !cekIdBidang(b, nBidang, idBidang) {
		fmt.Println("[!] ID bidang tidak ditemukan.")
		return
	}

	k[*n].idKursus = idKursus
	k[*n].idBidang = idBidang

	*n = *n + 1
	fmt.Println("[v] Data kursus berhasil ditambahkan.")
}

func tampilDataKursus(k tabKursus, n int) {
	var i int

	if n == 0 {
		fmt.Println("[!] Belum ada data kursus.")
		return
	}

	fmt.Println("\n======= DATA KURSUS =======")
	for i = 0; i < n; i++ {
		fmt.Printf("  %-14s : %d\n", "ID Kursus", k[i].idKursus)
		fmt.Printf("  %-14s : %s\n", "Nama Kursus", k[i].namaKursus)
		fmt.Printf("  %-14s : %d\n", "ID Bidang", k[i].idBidang)
		fmt.Println("---------------------------")
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

func cetakDetailPeserta(p Peserta) {
	fmt.Printf("  %-14s : %d\n", "ID Pendaftaran", p.idPendaftaran)
	fmt.Printf("  %-14s : %s\n", "Nama Lengkap", p.namaLengkap)
	fmt.Printf("  %-14s : %s\n", "Tanggal Daftar", p.tanggalDaftar)
	fmt.Printf("  %-14s : %d\n", "ID Kursus", p.idKursus)
	fmt.Printf("  %-14s : %d\n", "ID Bidang", p.idBidang)
	fmt.Printf("  %-14s : %v\n", "Status Aktif", p.statusAktif)
	fmt.Println("---------------------------------")
}

func tambahDataPeserta(p *tabPeserta, n *int, k tabKursus, nKursus int, b tabBidang, nBidang int) {
	var idPendaftaran, idKursus, idBidang int

	if nKursus == 0 || nBidang == 0 {
		fmt.Println("[!] Data kursus atau bidang belum ada. Silakan tambahkan terlebih dahulu.")
		return
	}

	if *n >= NMAX {
		fmt.Println("[!] Data peserta sudah penuh.")
		return
	}

	fmt.Println("\n--- Tambah Data Peserta ---")

	idPendaftaran = validasiIDPositif("ID Pendaftaran              : ")
	if cekIdPeserta(*p, *n, idPendaftaran) {
		fmt.Println("[!] ID pendaftaran sudah digunakan.")
		return
	}

	p[*n].namaLengkap = validasiNama("Nama Lengkap                : ")
	p[*n].tanggalDaftar = validasiTanggal("Tanggal Daftar (DD-MM-YYYY) : ")
	idKursus = validasiIDPositif("ID Kursus                   : ")

	if !cekIdKursus(k, nKursus, idKursus) {
		fmt.Println("[!] ID kursus tidak ditemukan.")
		return
	}

	idBidang = ambilIdBidangDariKursus(k, nKursus, idKursus)
	fmt.Printf("  ID Bidang (otomatis)        : %d\n", idBidang)

	p[*n].idPendaftaran = idPendaftaran
	p[*n].idKursus = idKursus
	p[*n].idBidang = idBidang
	p[*n].statusAktif = validasiStatusAktif("Status Aktif (true/false)   : ")

	*n = *n + 1
	fmt.Println("[v] Data peserta berhasil ditambahkan.")
}

func tampilDataPeserta(p tabPeserta, n int) {
	var i int

	if n == 0 {
		fmt.Println("[!] Belum ada data peserta.")
		return
	}

	fmt.Println("\n======= DATA PESERTA =======")
	for i = 0; i < n; i++ {
		cetakDetailPeserta(p[i])
	}
}

func ubahDataPeserta(p *tabPeserta, n int, k tabKursus, nKursus int, b tabBidang, nBidang int) {
	var cari, i, pilih int
	var idBaru int
	var ketemu, selesai bool

	if n == 0 {
		fmt.Println("[!] Data peserta masih kosong. Silakan tambah data terlebih dahulu.")
		return
	}

	cari = validasiIDPositif("ID Pendaftaran yang ingin diubah : ")
	ketemu = false

	for i = 0; i < n && !ketemu; i++ {
		if p[i].idPendaftaran == cari {
			ketemu = true
			selesai = false

			for !selesai {
				fmt.Println("\n--- Ubah Data Peserta ---")
				fmt.Println("  1. Ubah Nama Lengkap")
				fmt.Println("  2. Ubah Tanggal Daftar")
				fmt.Println("  3. Ubah ID Kursus")
				fmt.Println("  4. Ubah Status Aktif")
				fmt.Println("  5. Selesai")
				fmt.Print("Pilihan : ")
				fmt.Scan(&pilih)

				if pilih == 1 {
					p[i].namaLengkap = validasiNama("Nama baru : ")
				} else if pilih == 2 {
					p[i].tanggalDaftar = validasiTanggal("Tanggal daftar baru (DD-MM-YYYY) : ")
				} else if pilih == 3 {
					idBaru = validasiIDPositif("ID kursus baru : ")
					if !cekIdKursus(k, nKursus, idBaru) {
						fmt.Println("[!] ID kursus tidak ditemukan.")
					} else {
						p[i].idKursus = idBaru
						p[i].idBidang = ambilIdBidangDariKursus(k, nKursus, idBaru)
						fmt.Printf("  ID Bidang otomatis diperbarui : %d\n", p[i].idBidang)
					}
				} else if pilih == 4 {
					p[i].statusAktif = validasiStatusAktif("Status aktif baru (true/false) : ")
				} else if pilih == 5 {
					selesai = true
				} else {
					fmt.Println("[!] Pilihan tidak valid.")
				}
			}
			fmt.Println("[v] Data peserta berhasil diubah.")
		}
	}

	if !ketemu {
		fmt.Println("[!] Data peserta tidak ditemukan.")
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
		fmt.Println("[!] Data peserta masih kosong. Silakan tambah data terlebih dahulu.")
		return
	}

	cari = validasiIDPositif("ID Pendaftaran yang ingin dihapus : ")
	idx = cariIDPeserta(*p, *n, cari)

	if idx == -1 {
		fmt.Println("[!] Data peserta tidak ditemukan.")
	} else {
		for i = idx; i < *n-1; i++ {
			p[i] = p[i+1]
		}
		*n = *n - 1
		fmt.Println("[v] Data peserta berhasil dihapus.")
	}
}

func cariPesertaBidang(p tabPeserta, n int) {
	var i, cariBidang int
	var ketemu bool

	if n == 0 {
		fmt.Println("[!] Data peserta masih kosong. Silakan tambah data terlebih dahulu.")
		return
	}

	cariBidang = validasiIDPositif("ID Bidang yang dicari : ")
	ketemu = false
	fmt.Println("\n======= Hasil Pencarian Berdasarkan Bidang =======")

	for i = 0; i < n; i++ {
		if p[i].idBidang == cariBidang {
			cetakDetailPeserta(p[i])
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("[!] Tidak ada peserta dengan ID bidang tersebut.")
	}
}

func urutNamaPesertaAscending(temp *tabPeserta, n int) {
	var i, j int
	var tempPeserta Peserta

	if n == 0 {
		fmt.Println("[!] Data peserta masih kosong. Tidak ada data yang bisa diurutkan.")
		return
	}

	for i = 1; i < n; i++ {
		tempPeserta = temp[i]
		j = i - 1
		for j >= 0 && temp[j].namaLengkap > tempPeserta.namaLengkap {
			temp[j+1] = temp[j]
			j = j - 1
		}
		temp[j+1] = tempPeserta
	}
	fmt.Println("[v] Data peserta berhasil diurutkan berdasarkan nama (A-Z).")
}

func urutNamaPesertaDescending(p *tabPeserta, n int) {
	var i, j int
	var temp Peserta

	if n == 0 {
		fmt.Println("[!] Data peserta masih kosong. Tidak ada data yang bisa diurutkan.")
		return
	}

	for i = 1; i < n; i++ {
		temp = p[i]
		j = i - 1
		for j >= 0 && p[j].namaLengkap < temp.namaLengkap {
			p[j+1] = p[j]
			j = j - 1
		}
		p[j+1] = temp
	}
	fmt.Println("[v] Data peserta berhasil diurutkan berdasarkan nama (Z-A).")
}

func urutIDPesertaAscending(p *tabPeserta, n int) {
	var pass, idx, i int
	var temp Peserta

	if n == 0 {
		fmt.Println("[!] Data peserta masih kosong. Tidak ada data yang bisa diurutkan.")
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
	fmt.Println("[v] Data peserta berhasil diurutkan berdasarkan ID (ascending).")
}

func urutIDPesertaDescending(p *tabPeserta, n int) {
	var pass, idx, i int
	var temp Peserta

	if n == 0 {
		fmt.Println("[!] Data peserta masih kosong. Tidak ada data yang bisa diurutkan.")
		return
	}

	for pass = 1; pass <= n-1; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if p[i].idPendaftaran > p[idx].idPendaftaran {
				idx = i
			}
		}
		temp = p[pass-1]
		p[pass-1] = p[idx]
		p[idx] = temp
	}
	fmt.Println("[v] Data peserta berhasil diurutkan berdasarkan ID (Descending).")
}

func cariPesertaNama(p tabPeserta, n int) {
	var cari string
	var kiri, kanan, tengah, i int
	var temp tabPeserta
	var ketemu bool

	if n == 0 {
		fmt.Println("[!] Data peserta masih kosong.")
		return
	}

	fmt.Print("Nama lengkap yang dicari : ")
	fmt.Scan(&cari)

	for i = 0; i < n; i++ {
		temp[i] = p[i]
	}

	urutNamaPesertaAscending(&temp, n)

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
		fmt.Println("\n======= Hasil Pencarian Berdasarkan Nama =======")
		i = tengah
		for i >= 0 && temp[i].namaLengkap == cari {
			cetakDetailPeserta(temp[i])
			i--
		}
		i = tengah + 1
		for i < n && temp[i].namaLengkap == cari {
			cetakDetailPeserta(temp[i])
			i++
		}
	} else {
		fmt.Println("[!] Data peserta tidak ditemukan.")
	}
}

func statistikPeserta(p tabPeserta, nPeserta int, k tabKursus, nKursus int, b tabBidang, nBidang int) {
	var i, j, jumlah, totalAktif, totalTidakAktif int
	var persentase float64

	if nBidang == 0 || nPeserta == 0 {
		fmt.Println("[!] Data peserta atau bidang minat masih kosong. Belum ada statistik.")
		return
	}

	fmt.Println("\n======= STATISTIK PESERTA =======")
	fmt.Printf("  %-28s : %d\n", "Total Peserta Terdaftar", nPeserta)

	totalAktif = 0
	totalTidakAktif = 0
	for i = 0; i < nPeserta; i++ {
		if p[i].statusAktif == true {
			totalAktif++
		} else {
			totalTidakAktif++
		}
	}

	persentase = float64(totalAktif) / float64(nPeserta) * 100
	fmt.Printf("  %-28s : %d (%.2f%%)\n", "Total Peserta Aktif", totalAktif, persentase)

	persentase = float64(totalTidakAktif) / float64(nPeserta) * 100
	fmt.Printf("  %-28s : %d (%.2f%%)\n", "Total Peserta Tidak Aktif", totalTidakAktif, persentase)

	fmt.Println("\n  -- Distribusi per Bidang Minat --")
	for i = 0; i < nBidang; i++ {
		jumlah = 0
		for j = 0; j < nPeserta; j++ {
			if p[j].idBidang == b[i].idBidang {
				jumlah++
			}
		}
		persentase = float64(jumlah) / float64(nPeserta) * 100
		fmt.Printf("  %-20s : %d peserta (%.2f%%)\n", b[i].namaBidang, jumlah, persentase)
	}

	if nKursus > 0 {
		fmt.Println("\n  -- Distribusi per Kursus --")
		for i = 0; i < nKursus; i++ {
			jumlah = 0
			for j = 0; j < nPeserta; j++ {
				if p[j].idKursus == k[i].idKursus {
					jumlah++
				}
			}
			persentase = float64(jumlah) / float64(nPeserta) * 100
			fmt.Printf("  %-20s : %d peserta (%.2f%%)\n", k[i].namaKursus, jumlah, persentase)
		}
	}
	fmt.Println("---------------------------------")
}

func menuDataPeserta(p *tabPeserta, nPeserta *int, k tabKursus, nKursus int, b tabBidang, nBidang int) {
	var pilih, subPilih int

	for pilih != 5 {
		fmt.Println("\n========= MENU DATA PESERTA =========")
		fmt.Println("  1. Kelola Data (Tambah/Ubah/Hapus/Tampil)")
		fmt.Println("  2. Urutkan Data")
		fmt.Println("  3. Cari Data")
		fmt.Println("  4. Statistik Peserta")
		fmt.Println("  5. Kembali ke Menu Utama")
		fmt.Println("----------------------------------------")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			for subPilih != 5 {
				fmt.Println("\n------ Kelola Data Peserta ------")
				fmt.Println("  1. Tambah Data Peserta")
				fmt.Println("  2. Ubah Data Peserta")
				fmt.Println("  3. Hapus Data Peserta")
				fmt.Println("  4. Tampil Data Peserta")
				fmt.Println("  5. Kembali")
				fmt.Println("----------------------------------")
				fmt.Print("Pilihan : ")
				fmt.Scan(&subPilih)

				if subPilih == 1 {
					tambahDataPeserta(p, nPeserta, k, nKursus, b, nBidang)
				} else if subPilih == 2 {
					ubahDataPeserta(p, *nPeserta, k, nKursus, b, nBidang)
				} else if subPilih == 3 {
					hapusDataPeserta(p, nPeserta)
				} else if subPilih == 4 {
					tampilDataPeserta(*p, *nPeserta)
				} else if subPilih == 5 {
					fmt.Println("Kembali ke Menu Data Peserta...")
				} else {
					fmt.Println("[!] Pilihan tidak valid.")
				}
			}
			subPilih = 0

		} else if pilih == 2 {
			for subPilih != 5 {
				fmt.Println("\n------ Urutkan Data Peserta ------")
				fmt.Println("  1. Urut ID Peserta (Ascending)")
				fmt.Println("  2. Urut ID Peserta (Descending)")
				fmt.Println("  3. Urut Nama Peserta (Ascending A-Z)")
				fmt.Println("  4. Urut Nama Peserta (Descending Z-A)")
				fmt.Println("  5. Kembali")
				fmt.Println("-------------------------------------")
				fmt.Print("Pilihan : ")
				fmt.Scan(&subPilih)

				if subPilih == 1 {
					urutIDPesertaAscending(p, *nPeserta)
					tampilDataPeserta(*p, *nPeserta)
				} else if subPilih == 2 {
					urutIDPesertaDescending(p, *nPeserta)
					tampilDataPeserta(*p, *nPeserta)
				} else if subPilih == 3 {
					urutNamaPesertaAscending(p, *nPeserta)
					tampilDataPeserta(*p, *nPeserta)
				} else if subPilih == 4 {
					urutNamaPesertaDescending(p, *nPeserta)
					tampilDataPeserta(*p, *nPeserta)
				} else if subPilih == 5 {
					fmt.Println("Kembali ke Menu Data Peserta...")
				} else {
					fmt.Println("[!] Pilihan tidak valid.")
				}
			}
			subPilih = 0

		} else if pilih == 3 {
			for subPilih != 3 {
				fmt.Println("\n------ Cari Data Peserta ------")
				fmt.Println("  1. Cari Berdasarkan Bidang Minat")
				fmt.Println("  2. Cari Berdasarkan Nama Lengkap")
				fmt.Println("  3. Kembali")
				fmt.Println("----------------------------------")
				fmt.Print("Pilihan : ")
				fmt.Scan(&subPilih)

				if subPilih == 1 {
					cariPesertaBidang(*p, *nPeserta)
				} else if subPilih == 2 {
					cariPesertaNama(*p, *nPeserta)
				} else if subPilih == 3 {
					fmt.Println("Kembali ke Menu Data Peserta...")
				} else {
					fmt.Println("[!] Pilihan tidak valid.")
				}
			}
			subPilih = 0

		} else if pilih == 4 {
			statistikPeserta(*p, *nPeserta, k, nKursus, b, nBidang)
		} else if pilih == 5 {
			fmt.Println("Kembali ke Menu Utama...")
		} else {
			fmt.Println("[!] Pilihan tidak valid.")
		}
	}
}

func menuDataKursus(k *tabKursus, nKursus *int, b tabBidang, nBidang int) {
	var pilih int

	for pilih != 3 {
		fmt.Println("\n======= MENU DATA KURSUS =======")
		fmt.Println("  1. Tambah Data Kursus")
		fmt.Println("  2. Tampil Data Kursus")
		fmt.Println("  3. Kembali ke Menu Utama")
		fmt.Println("---------------------------------")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahDataKursus(k, nKursus, b, nBidang)
		} else if pilih == 2 {
			tampilDataKursus(*k, *nKursus)
		} else if pilih == 3 {
			fmt.Println("Kembali ke Menu Utama...")
		} else {
			fmt.Println("[!] Pilihan tidak valid.")
		}
	}
}

func menuDataBidang(b *tabBidang, nBidang *int) {
	var pilih int

	for pilih != 3 {
		fmt.Println("\n======= MENU DATA BIDANG MINAT =======")
		fmt.Println("  1. Tambah Bidang Minat")
		fmt.Println("  2. Tampil Bidang Minat")
		fmt.Println("  3. Kembali ke Menu Utama")
		fmt.Println("--------------------------------------")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahBidangMinat(b, nBidang)
		} else if pilih == 2 {
			tampilBidangMinat(*b, *nBidang)
		} else if pilih == 3 {
			fmt.Println("Kembali ke Menu Utama...")
		} else {
			fmt.Println("[!] Pilihan tidak valid.")
		}
	}
}

func main() {
	var dataPeserta tabPeserta
	var dataKursus tabKursus
	var dataBidang tabBidang

	var nPeserta, nKursus, nBidang int
	var pilih int

	for pilih != 4 {
		fmt.Println("\n======= MENU UTAMA KURSUSIN =======")
		fmt.Println("  1. Data Peserta")
		fmt.Println("  2. Data Kursus")
		fmt.Println("  3. Data Bidang Minat")
		fmt.Println("  4. Keluar")
		fmt.Println("-----------------------------------")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			menuDataPeserta(&dataPeserta, &nPeserta, dataKursus, nKursus, dataBidang, nBidang)
		} else if pilih == 2 {
			menuDataKursus(&dataKursus, &nKursus, dataBidang, nBidang)
		} else if pilih == 3 {
			menuDataBidang(&dataBidang, &nBidang)
		} else if pilih == 4 {
			fmt.Println("Terima kasih telah menggunakan KursusIn.")
		} else {
			fmt.Println("[!] Pilihan tidak valid.")
		}
	}
}
