# KursusIn

Sistem Pendaftaran Kursus Online Terpadu berbasis Command Line Interface (CLI) yang dikembangkan menggunakan bahasa pemrograman Go untuk memenuhi Tugas Besar Mata Kuliah Algoritma Pemrograman 2.

Aplikasi ini digunakan untuk mengelola proses pendaftaran peserta pada berbagai program kursus online, termasuk pengelolaan data peserta, kursus, bidang minat, pencarian data, pengurutan data, dan penyajian statistik peserta.

---

## Deskripsi

KursusIn merupakan sistem manajemen pendaftaran kursus yang dirancang untuk membantu administrator dalam mengelola data peserta secara terstruktur. Sistem memanfaatkan implementasi algoritma searching dan sorting sebagai bagian dari penerapan konsep yang dipelajari pada mata kuliah Algoritma Pemrograman 2.

Data utama yang dikelola meliputi:

* Data Peserta
* Data Kursus
* Data Bidang Minat

---

## Fitur

### Manajemen Data Peserta

* Tambah data peserta
* Ubah data peserta
* Hapus data peserta
* Tampilkan data peserta

### Manajemen Data Kursus

* Tambah data kursus
* Tampilkan data kursus

### Manajemen Data Bidang Minat

* Tambah bidang minat
* Tampilkan bidang minat

### Pencarian Data (Searching)

* Sequential Search berdasarkan bidang minat
* Binary Search berdasarkan nama peserta

### Pengurutan Data (Sorting)

* Selection Sort berdasarkan ID pendaftaran
* Insertion Sort berdasarkan nama peserta

### Statistik

* Menampilkan jumlah peserta pada setiap bidang minat
* Menampilkan total peserta aktif

### Validasi Data

* Mencegah duplikasi ID peserta
* Mencegah duplikasi ID kursus
* Mencegah duplikasi ID bidang minat
* Memastikan ID kursus dan bidang minat tersedia sebelum peserta ditambahkan

---

## Flow Penggunaan Aplikasi

### 1. Menambahkan Data Bidang Minat

Administrator menambahkan bidang minat yang tersedia sebagai kategori kursus.

Contoh:

* Pemrograman
* Data Science
* Desain Grafis
* Digital Marketing

Menu:

```text
7. Tambah Bidang Minat
```

---

### 2. Menambahkan Data Kursus

Setelah bidang minat tersedia, administrator dapat menambahkan kursus baru yang terhubung dengan bidang minat tertentu.

Menu:

```text
5. Tambah Data Kursus
```

---

### 3. Menambahkan Data Peserta

Peserta didaftarkan dengan mengisi:

* ID Pendaftaran
* Nama Lengkap
* Tanggal Daftar
* ID Kursus
* ID Bidang Minat
* Status Aktif

Menu:

```text
1. Tambah Data Peserta
```

---

### 4. Mengelola Data Peserta

Administrator dapat:

* Mengubah data peserta
* Menghapus data peserta
* Menampilkan data peserta

Menu:

```text
2. Ubah Data Peserta
3. Hapus Data Peserta
4. Tampilkan Data Peserta
```

---

### 5. Melakukan Pencarian Data

#### Sequential Search

Digunakan untuk mencari peserta berdasarkan bidang minat.

Menu:

```text
9. Cari Peserta Berdasarkan Bidang
```

#### Binary Search

Digunakan untuk mencari peserta berdasarkan nama.

Menu:

```text
10. Cari Peserta Berdasarkan Nama
```

Catatan:

Data peserta sebaiknya diurutkan terlebih dahulu berdasarkan nama sebelum menggunakan Binary Search.

---

### 6. Melakukan Pengurutan Data

#### Selection Sort

Mengurutkan data peserta berdasarkan ID pendaftaran.

Menu:

```text
11. Urutkan Peserta Berdasarkan ID
```

#### Insertion Sort

Mengurutkan data peserta berdasarkan nama secara alfabetis.

Menu:

```text
12. Urutkan Peserta Berdasarkan Nama
```

---

### 7. Menampilkan Statistik

Administrator dapat melihat:

* Jumlah peserta pada setiap bidang minat
* Total peserta aktif

Menu:

```text
13. Statistik Peserta
```

---

### Diagram Alur

```text
Mulai
  │
  ▼
Tambah Bidang Minat
  │
  ▼
Tambah Kursus
  │
  ▼
Tambah Peserta
  │
  ▼
Kelola Data Peserta
  │
  ├── Ubah Data
  ├── Hapus Data
  └── Tampilkan Data
  │
  ▼
Pencarian Data
  │
  ├── Sequential Search
  └── Binary Search
  │
  ▼
Pengurutan Data
  │
  ├── Selection Sort
  └── Insertion Sort
  │
  ▼
Statistik Peserta
  │
  ▼
Selesai
```

---

## Struktur Proyek

```text
kursus-in/
├── main.go
├── Dokumentasi-KursusIn.jpeg
└── README.md
```

---

## Cara Menjalankan

Pastikan Go telah terinstal pada perangkat Anda.

Jalankan program dengan perintah:

```bash
go run main.go
```

---

## Algoritma yang Digunakan

| Kategori  | Algoritma         |
| --------- | ----------------- |
| Searching | Sequential Search |
| Searching | Binary Search     |
| Sorting   | Selection Sort    |
| Sorting   | Insertion Sort    |

---

## Dokumentasi

Dokumentasi proyek tersedia pada file:

```text
Dokumentasi-KursusIn.jpeg
```

---

## Repository

https://github.com/dpndri/kursus-in

---

## Kelompok

**Kelompok 11**

### Anggota

* Muhammad Irgie Dapiandri (103012500187)
* Angga Yudhistira (103012500322)

### Mata Kuliah

Algoritma Pemrograman 2

### Universitas

Telkom University
