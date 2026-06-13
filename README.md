# KursusIn — Sistem Pendaftaran Kursus Online Terpadu

> **Tugas Besar | Mata Kuliah Algoritma dan Pemrograman 2**
> Nomor Soal: 22 — Sistem Pendaftaran Kursus Online Terpadu (KursusIn)

---

## 📝 Deskripsi Aplikasi

**KursusIn** adalah aplikasi berbasis *Command Line Interface* (CLI) yang dikembangkan menggunakan bahasa pemrograman **Go (Golang)** untuk mengelola registrasi peserta pada berbagai program pelatihan digital. Data utama yang dikelola meliputi data peserta, katalog kursus, dan bidang minat. Pengguna aplikasi adalah admin pendaftaran atau koordinator kursus.

Aplikasi ini mengimplementasikan konsep-konsep pemrograman prosedural secara murni, meliputi modularitas (*subprogram/fungsi*), tipe data bentukan (*struct*), array statis, validasi input, serta algoritma pengurutan dan pencarian tanpa bergantung pada *library* eksternal pihak ketiga — hanya menggunakan paket bawaan `fmt`.

---

## 📌 Daftar Isi

- [Spesifikasi Sistem](#-spesifikasi-sistem)
- [Fitur Utama](#-fitur-utama)
- [Struktur Data](#-struktur-data)
- [Algoritma yang Diimplementasikan](#-algoritma-yang-diimplementasikan)
- [Validasi Input](#-validasi-input)
- [Struktur Menu CLI](#-struktur-menu-cli)
- [Cara Menjalankan](#-cara-menjalankan)
- [Identitas Kelompok](#-identitas-kelompok)

---

## ✅ Spesifikasi Sistem

| No. | Spesifikasi |
|-----|-------------|
| a   | Pengguna dapat menambahkan, mengubah, dan menghapus data peserta kursus. |
| b   | Sistem mencatat tanggal pendaftaran dan bidang minat yang diambil setiap peserta secara otomatis berdasarkan kursus yang dipilih. |
| c   | Pengguna dapat mencari data peserta berdasarkan **bidang minat** (*Sequential Search*) atau **nama lengkap** (*Binary Search*). |
| d   | Pengguna dapat mengurutkan data peserta berdasarkan **ID pendaftaran** (*Selection Sort*) atau **nama** secara alfabetis (*Insertion Sort*). |
| e   | Sistem menampilkan statistik jumlah pendaftar per bidang minat, per kursus, serta total peserta aktif dan tidak aktif. |

---

## 🚀 Fitur Utama

### 1. Manajemen Data Bidang Minat
- **Tambah Bidang Minat:** Menambahkan kategori bidang pelatihan baru dengan validasi ID unik.
- **Tampil Bidang Minat:** Menampilkan seluruh data bidang minat yang tersedia.

> Bidang minat harus ditambahkan terlebih dahulu sebelum data kursus atau peserta dapat dimasukkan.

### 2. Manajemen Data Kursus
- **Tambah Data Kursus:** Menambahkan program kursus baru yang wajib berelasi dengan ID Bidang Minat yang sudah ada (validasi relasional).
- **Tampil Data Kursus:** Menampilkan seluruh data kursus beserta ID Bidang Minat yang terhubung.

### 3. Manajemen Data Peserta (CRUD Lengkap)
- **Tambah Data Peserta:** Registrasi peserta baru dengan validasi ID unik, validasi format tanggal (`DD-MM-YYYY`), validasi keberadaan ID Kursus, dan pengisian otomatis `idBidang` berdasarkan kursus yang dipilih untuk menjaga integritas relasi data.
- **Ubah Data Peserta:** Mengubah satu atau lebih atribut peserta (Nama Lengkap, Tanggal Daftar, ID Kursus, atau Status Aktif) melalui sub-menu interaktif, tanpa mengubah ID Pendaftaran. Perubahan ID Kursus akan secara otomatis memperbarui `idBidang`.
- **Hapus Data Peserta:** Menghapus data peserta berdasarkan ID Pendaftaran menggunakan mekanisme pergeseran indeks sekuensial pada array.
- **Tampil Data Peserta:** Menampilkan seluruh data peserta yang terdaftar beserta detailnya.

### 4. Pencarian Data Peserta
- **Cari Berdasarkan Bidang Minat:** Menggunakan **Sequential Search** untuk menelusuri dan menampilkan semua peserta yang terdaftar pada bidang minat tertentu (berdasarkan ID Bidang).
- **Cari Berdasarkan Nama Lengkap:** Menggunakan **Binary Search** pada salinan array sementara yang terlebih dahulu diurutkan dengan Insertion Sort. Dilengkapi penelusuran maju-mundur dari titik tengah (`tengah`) untuk menangani kasus nama peserta yang duplikat.

### 5. Pengurutan Data Peserta
- **Urut berdasarkan ID Peserta (Ascending):** Menggunakan **Selection Sort** — mencari nilai terkecil di setiap *pass*.
- **Urut berdasarkan ID Peserta (Descending):** Menggunakan **Selection Sort** — mencari nilai terbesar di setiap *pass*.
- **Urut berdasarkan Nama Lengkap (A–Z):** Menggunakan **Insertion Sort** secara *ascending*.
- **Urut berdasarkan Nama Lengkap (Z–A):** Menggunakan **Insertion Sort** secara *descending*.

> Hasil pengurutan langsung ditampilkan setelah proses selesai.

### 6. Statistik Peserta
Menyajikan ringkasan data analitik kepada koordinator kursus, meliputi:
- Total peserta terdaftar.
- Total peserta **Aktif** beserta persentasenya.
- Total peserta **Tidak Aktif** beserta persentasenya.
- Distribusi jumlah peserta dan persentase per **Bidang Minat**.
- Distribusi jumlah peserta dan persentase per **Kursus**.

---

## 📊 Struktur Data

Program menggunakan array statis dengan ukuran tetap (`NMAX = 1000`) untuk seluruh entitas data.

```go
const NMAX = 1000

// Entitas peserta yang mendaftar kursus
type Peserta struct {
    idPendaftaran int    // ID unik peserta (> 0)
    namaLengkap   string // Nama peserta (harus mengandung huruf)
    tanggalDaftar string // Format: DD-MM-YYYY
    idKursus      int    // Referensi ke Kursus.idKursus
    idBidang      int    // Diisi otomatis dari kursus yang dipilih
    statusAktif   bool   // true = aktif, false = tidak aktif
}

// Entitas program kursus
type Kursus struct {
    idKursus   int    // ID unik kursus (> 0)
    namaKursus string // Nama program kursus
    idBidang   int    // Referensi ke BidangMinat.idBidang
}

// Entitas kategori bidang pelatihan
type BidangMinat struct {
    idBidang   int    // ID unik bidang minat (> 0)
    namaBidang string // Nama bidang minat
}

// Tipe array statis untuk setiap entitas
type tabPeserta [NMAX]Peserta
type tabKursus  [NMAX]Kursus
type tabBidang  [NMAX]BidangMinat
```

### Relasi Antar Entitas

```
BidangMinat (1) ──< Kursus (1) ──< Peserta
  idBidang  ──────── idBidang    
                     idKursus ──── idKursus
```

Setiap `Kursus` harus berelasi dengan satu `BidangMinat` yang valid. Setiap `Peserta` harus berelasi dengan satu `Kursus` yang valid, dan field `idBidang` pada peserta diisi **secara otomatis** dari kursus yang dipilih.

---

## 🔬 Algoritma yang Diimplementasikan

### Selection Sort — Pengurutan ID Peserta

Digunakan untuk mengurutkan data berdasarkan `idPendaftaran` (ascending maupun descending). Pada setiap *pass*, algoritma mencari indeks nilai minimum (atau maksimum) dari sisa elemen yang belum terurut, lalu menukarnya dengan elemen pada posisi *pass* saat ini.

### Insertion Sort — Pengurutan Nama Peserta

Digunakan untuk mengurutkan data berdasarkan `namaLengkap` (ascending A–Z maupun descending Z–A), serta sebagai pra-proses sebelum Binary Search dijalankan. Setiap elemen disisipkan ke posisi yang tepat dalam bagian array yang sudah terurut.

### Sequential Search — Pencarian Berdasarkan Bidang Minat

Menelusuri seluruh elemen array `tabPeserta` dari indeks 0 hingga n-1 dan menampilkan semua peserta yang memiliki `idBidang` yang cocok dengan nilai yang dicari.

### Binary Search — Pencarian Berdasarkan Nama Lengkap

Bekerja pada salinan array sementara (`temp`) yang terlebih dahulu diurutkan menggunakan Insertion Sort. Pencarian dilakukan dengan membandingkan nilai tengah (`tengah = (kiri + kanan) / 2`) secara berulang. Setelah ditemukan, algoritma melakukan penelusuran maju dan mundur dari titik `tengah` untuk menampilkan semua entri dengan nama yang sama (duplikat).

---

## 🛡️ Validasi Input

Setiap input dari pengguna divalidasi sebelum disimpan. Seluruh fungsi validasi menggunakan perulangan (`for !valid`) hingga input yang benar diterima.

| Fungsi | Aturan Validasi |
|--------|-----------------|
| `validasiIDPositif` | Nilai harus berupa bilangan bulat dan lebih dari 0. |
| `validasiNama` | Tidak boleh kosong dan harus mengandung minimal satu karakter huruf (a–z / A–Z). |
| `validasiTanggal` | Panjang tepat 10 karakter, format `DD-MM-YYYY`, pemisah `-` pada posisi ke-3 dan ke-6, serta semua bagian angka hanya berisi digit. |
| `validasiStatusAktif` | Hanya menerima string `"true"` atau `"false"`. |
| `cekIdPeserta` | Memastikan `idPendaftaran` belum digunakan (unik). |
| `cekIdKursus` | Memastikan `idKursus` yang direferensikan benar-benar ada. |
| `cekIdBidang` | Memastikan `idBidang` yang direferensikan benar-benar ada. |

---

## 🗂️ Struktur Menu CLI

```
MENU UTAMA KURSUSIN
├── 1. Data Peserta
│   ├── 1. Kelola Data (Tambah/Ubah/Hapus/Tampil)
│   │   ├── 1. Tambah Data Peserta
│   │   ├── 2. Ubah Data Peserta
│   │   ├── 3. Hapus Data Peserta
│   │   ├── 4. Tampil Data Peserta
│   │   └── 5. Kembali
│   ├── 2. Urutkan Data
│   │   ├── 1. Urut ID Peserta (Ascending)
│   │   ├── 2. Urut ID Peserta (Descending)
│   │   ├── 3. Urut Nama Peserta (A–Z)
│   │   ├── 4. Urut Nama Peserta (Z–A)
│   │   └── 5. Kembali
│   ├── 3. Cari Data
│   │   ├── 1. Cari Berdasarkan Bidang Minat
│   │   ├── 2. Cari Berdasarkan Nama Lengkap
│   │   └── 3. Kembali
│   ├── 4. Statistik Peserta
│   └── 5. Kembali ke Menu Utama
├── 2. Data Kursus
│   ├── 1. Tambah Data Kursus
│   ├── 2. Tampil Data Kursus
│   └── 3. Kembali ke Menu Utama
├── 3. Data Bidang Minat
│   ├── 1. Tambah Bidang Minat
│   ├── 2. Tampil Bidang Minat
│   └── 3. Kembali ke Menu Utama
└── 4. Keluar
```

---

## ▶️ Cara Menjalankan

### Prasyarat

- Go (Golang) versi 1.18 atau lebih baru telah terpasang di sistem.
- Tidak ada dependensi eksternal — hanya menggunakan paket `fmt` dari pustaka standar Go.

### Langkah Menjalankan

1. Pastikan file `main.go` tersedia di direktori kerja.
2. Buka terminal, lalu jalankan perintah berikut:

```bash
go run main.go
```

Atau, kompilasi terlebih dahulu menjadi *executable*:

```bash
go build -o kursusin main.go
./kursusin
```

### Urutan Penggunaan yang Disarankan

Karena adanya ketergantungan relasional antar entitas, ikuti urutan berikut saat pertama kali menggunakan aplikasi:

```
1. Tambah Bidang Minat  →  2. Tambah Kursus  →  3. Tambah Peserta
```

---

## 👥 Identitas Kelompok

| Atribut | Keterangan |
|---------|------------|
| Mata Kuliah | Algoritma dan Pemrograman 2 |
| Tugas | Tugas Besar — Nomor 22 |
| Nama Aplikasi | KursusIn — Sistem Pendaftaran Kursus Online Terpadu |
| Bahasa Pemrograman | Go (Golang) |
| Anggota 1 | *(Muhammad Irgie Dapiandri — 103012500187)* |
| Anggota 2 | *(Angga Yudhistira — 103012500322)* |

---

*Dikembangkan untuk memenuhi Tugas Besar Mata Kuliah Algoritma dan Pemrograman 2.*