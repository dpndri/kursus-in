package main

import (
	"os"
	"testing"
)

func seedSamplePeserta() (tabPeserta, int) {
	var p tabPeserta
	p[0] = Peserta{idPendaftaran: 103, namaLengkap: "Citra", tanggalDaftar: "12-05-2026", idKursus: 10, idBidang: 1, statusAktif: true}
	p[1] = Peserta{idPendaftaran: 101, namaLengkap: "Andi", tanggalDaftar: "10-05-2026", idKursus: 10, idBidang: 1, statusAktif: true}
	p[2] = Peserta{idPendaftaran: 102, namaLengkap: "Budi", tanggalDaftar: "11-05-2026", idKursus: 20, idBidang: 2, statusAktif: false}
	return p, 3
}

func TestUrutIDPesertaAscending(t *testing.T) {
	p, n := seedSamplePeserta()
	urutIDPesertaAscending(&p, n)

	expected := [3]int{101, 102, 103}
	for i := 0; i < n; i++ {
		if p[i].idPendaftaran != expected[i] {
			t.Fatalf("indeks %d: ingin ID %d, dapat %d", i, expected[i], p[i].idPendaftaran)
		}
	}
}

func TestUrutIDPesertaDescending(t *testing.T) {
	p, n := seedSamplePeserta()
	urutIDPesertaDescending(&p, n)

	expected := [3]int{103, 102, 101}
	for i := 0; i < n; i++ {
		if p[i].idPendaftaran != expected[i] {
			t.Fatalf("indeks %d: ingin ID %d, dapat %d", i, expected[i], p[i].idPendaftaran)
		}
	}
}

func TestUrutNamaPesertaAscendingDanDescending(t *testing.T) {
	p, n := seedSamplePeserta()

	urutNamaPesertaAscending(&p, n)
	if p[0].namaLengkap != "Andi" || p[1].namaLengkap != "Budi" || p[2].namaLengkap != "Citra" {
		t.Fatalf("urutan A-Z salah: %+v", p[:n])
	}

	urutNamaPesertaDescending(&p, n)
	if p[0].namaLengkap != "Citra" || p[1].namaLengkap != "Budi" || p[2].namaLengkap != "Andi" {
		t.Fatalf("urutan Z-A salah: %+v", p[:n])
	}
}

func TestRelasiDanPencarianID(t *testing.T) {
	p, n := seedSamplePeserta()

	idx := cariIDPeserta(p, n, 102)
	if idx != 2 {
		t.Fatalf("cariIDPeserta(102): ingin indeks 2, dapat %d", idx)
	}

	if cekIdPeserta(p, n, 999) {
		t.Fatal("cekIdPeserta(999) seharusnya false")
	}

	var k tabKursus
	k[0] = Kursus{idKursus: 10, namaKursus: "Backend_Go", idBidang: 1}
	k[1] = Kursus{idKursus: 20, namaKursus: "UI_Design", idBidang: 2}

	idBidang := ambilIdBidangDariKursus(k, 2, 20)
	if idBidang != 2 {
		t.Fatalf("ambilIdBidangDariKursus(20): ingin 2, dapat %d", idBidang)
	}
}

func TestFullInteractiveCLI(t *testing.T) {
	inputs := "3\n1\n1\nBackend\n1\n2\nUIUX\n2\n3\n" +
		"2\n1\n10\nGolang\n1\n1\n20\nFigma\n2\n2\n3\n" +
		"1\n1\n" +
		"1\n103\nCitra\n12-05-2026\n10\ntrue\n" +
		"1\n101\nAndi\n10-05-2026\n10\ntrue\n" +
		"1\n102\nBudi\n11-05-2026\n20\nfalse\n" +
		"2\n102\n4\ntrue\n5\n" +
		"4\n5\n" +
		"2\n1\n2\n3\n4\n5\n" +
		"3\n1\n1\n2\nAndi\n3\n" +
		"4\n" +
		"1\n3\n103\n5\n" +
		"5\n4\n"

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	_, _ = w.WriteString(inputs)
	_ = w.Close()

	oldStdin := os.Stdin
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }()

	main()
}
