package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	SISTEM PENERIMAAN MAHASISWA BUKIT DURI

# Sistem bedasarkan seleksi nilai UTBK apabila nilai lebih dari 600 maka diterima

======== SPESIFIKASI UMUM
✅ Program harus dibuat secara modular dengan menggunakan subprogram.
✅ Struktur subprogram ditentukan oleh masing-masing kelompok. Setiap subprogram harus dilengkapi dengan parameter dan spesifikasinya.
✅ Program harus mengimplementasikan Array dan Tipe Bentukan. Array statis harus digunakan, bukan array dinamis atau slice.
✅ Program harus mengimplementasikan algoritma Sequential dan Binary Search untuk pencarian data, pengubahan (edit), atau penghapusan data tertentu.
✅ Program harus mengimplementasikan algoritma Selection Sort dan Insertion Sort untuk mengurutkan data dalam kategori yang berbeda saat menampilkan hasil. Setiap kategori harus bisa diurutkan dalam urutan naik (ascending) maupun turun (descending).
✅ Penggunaan statement break (kecuali dalam repeat-until) dan continue tidak diperbolehkan.
✅ Variabel global hanya diperbolehkan untuk menyimpan array utama yang akan diproses.
*/
const max = 10

type dataPendaftar struct {
	nisn              int
	nama              string
	tempatLahir       string
	tanggalLahir      string
	jenisKelamin      string
	agama             string
	email             string
	jurusan           string
	asalSekolah       string
	tahunLulus        int
	jurusanYangDituju string
	nilaiUTBK         float64
	status            string
}

type dataSiswa = [max]dataPendaftar

// Tipe Bentukan untuk login
type login struct {
	nama     string
	password string
}

// Tipe bentukan untuk calon mahasiswa yang diterima
type dataMahasiswa struct {
	nim     int
	nama    string
	jurusan string
}
type mahasiswa = [max]dataMahasiswa

var mahasigma mahasiswa

func main() {
	var calonMahasiswa dataSiswa
	//Ubah nama dan password kalian yeah 💌
	var pengguna = [4]login{
		{"admin", "admin"},
		{"Jinggachan", "jinggachan"},
		{"Nevytachan", "nevytachan"},
		{"Nayyachan", "nayyachan"},
	}

	//Data dummy calon mahasiswa

	calonMahasiswa[0] = dataPendaftar{nisn: 1234567890, nama: "Andi Pratama", tempatLahir: "Surabaya", tanggalLahir: "2000-01-01", jenisKelamin: "Laki-laki", agama: "Islam", email: "andi@email.com", jurusan: "IPA", asalSekolah: "SMA 1 Surabaya", tahunLulus: 2018, jurusanYangDituju: "Teknik Informatika", nilaiUTBK: 0}
	calonMahasiswa[1] = dataPendaftar{nisn: 1234567891, nama: "Siti Nurjanah", tempatLahir: "Malang", tanggalLahir: "2001-02-02", jenisKelamin: "Perempuan", agama: "Islam", email: "siti@email.com", jurusan: "IPS", asalSekolah: "SMA 2 Malang", tahunLulus: 2019, jurusanYangDituju: "Manajemen", nilaiUTBK: 1}
	calonMahasiswa[2] = dataPendaftar{nisn: 1234567892, nama: "Budi Santoso", tempatLahir: "Surabaya", tanggalLahir: "2000-03-03", jenisKelamin: "Laki-laki", agama: "Kristen", email: "budi@email.com", jurusan: "IPA", asalSekolah: "SMA 3 Surabaya", tahunLulus: 2019, jurusanYangDituju: "Teknik Mesin", nilaiUTBK: 710}
	calonMahasiswa[3] = dataPendaftar{nisn: 1234567893, nama: "Dewi Lestari", tempatLahir: "Bandung", tanggalLahir: "1999-04-04", jenisKelamin: "Perempuan", agama: "Hindu", email: "dewi@email.com", jurusan: "IPA", asalSekolah: "SMA 4 Bandung", tahunLulus: 2018, jurusanYangDituju: "Teknik Elektro", nilaiUTBK: 720}
	calonMahasiswa[4] = dataPendaftar{nisn: 1234567894, nama: "Fajar Ramadhan", tempatLahir: "Yogyakarta", tanggalLahir: "2000-05-05", jenisKelamin: "Laki-laki", agama: "Islam", email: "fajar@email.com", jurusan: "IPS", asalSekolah: "SMA 5 Yogyakarta", tahunLulus: 2019, jurusanYangDituju: "Ekonomi", nilaiUTBK: 680}
	calonMahasiswa[5] = dataPendaftar{nisn: 1234567895, nama: "Hanafi Rizky", tempatLahir: "Jakarta", tanggalLahir: "1999-06-06", jenisKelamin: "Laki-laki", agama: "Islam", email: "hanafi@email.com", jurusan: "IPA", asalSekolah: "SMA 6 Jakarta", tahunLulus: 2018, jurusanYangDituju: "Fisika", nilaiUTBK: 1000}
	calonMahasiswa[6] = dataPendaftar{nisn: 1234567896, nama: "Indah Purnama", tempatLahir: "Semarang", tanggalLahir: "2001-07-07", jenisKelamin: "Perempuan", agama: "Islam", email: "indah@email.com", jurusan: "IPS", asalSekolah: "SMA 7 Semarang", tahunLulus: 2019, jurusanYangDituju: "Psikologi", nilaiUTBK: 630}
	calonMahasiswa[7] = dataPendaftar{nisn: 1234567897, nama: "Joko Susanto", tempatLahir: "Surakarta", tanggalLahir: "2000-08-08", jenisKelamin: "Laki-laki", agama: "Kristen", email: "joko@email.com", jurusan: "IPA", asalSekolah: "SMA 8 Surakarta", tahunLulus: 2019, jurusanYangDituju: "Teknik Industri", nilaiUTBK: 715}
	calonMahasiswa[8] = dataPendaftar{nisn: 1234567898, nama: "Kartika Sari", tempatLahir: "Solo", tanggalLahir: "1999-09-09", jenisKelamin: "Perempuan", agama: "Islam", email: "kartika@email.com", jurusan: "IPA", asalSekolah: "SMA 9 Solo", tahunLulus: 2018, jurusanYangDituju: "Teknik Kimia", nilaiUTBK: 705}
	calonMahasiswa[9] = dataPendaftar{nisn: 1234567899, nama: "Lutfi Fauzi", tempatLahir: "Surabaya", tanggalLahir: "2001-10-10", jenisKelamin: "Laki-laki", agama: "Islam", email: "lutfi@email.com", jurusan: "IPS", asalSekolah: "SMA 10 Surabaya", tahunLulus: 2019, jurusanYangDituju: "Ilmu Komunikasi", nilaiUTBK: 690}

	// calonMahasiswa[10] = dataPendaftar{nisn: 1234567900, nama: "Rina Marlina", tempatLahir: "Medan", tanggalLahir: "2001-11-11", jenisKelamin: "Perempuan", agama: "Kristen", email: "rina@email.com", jurusan: "IPA", asalSekolah: "SMA 11 Medan", tahunLulus: 2019, jurusanYangDituju: "Biologi", nilaiUTBK: 680}
	// calonMahasiswa[11] = dataPendaftar{nisn: 1234567901, nama: "Toni Suryadi", tempatLahir: "Jakarta", tanggalLahir: "1999-12-12", jenisKelamin: "Laki-laki", agama: "Islam", email: "toni@email.com", jurusan: "IPS", asalSekolah: "SMA 12 Jakarta", tahunLulus: 2018, jurusanYangDituju: "Sosiologi", nilaiUTBK: 675}
	// calonMahasiswa[12] = dataPendaftar{nisn: 1234567902, nama: "Siti Aisyah", tempatLahir: "Surabaya", tanggalLahir: "2000-01-13", jenisKelamin: "Perempuan", agama: "Islam", email: "aisyah@email.com", jurusan: "IPA", asalSekolah: "SMA 13 Surabaya", tahunLulus: 2019, jurusanYangDituju: "Teknik Lingkungan", nilaiUTBK: 690}
	// calonMahasiswa[13] = dataPendaftar{nisn: 1234567903, nama: "Wawan Prasetyo", tempatLahir: "Semarang", tanggalLahir: "1999-02-14", jenisKelamin: "Laki-laki", agama: "Buddha", email: "wawan@email.com", jurusan: "IPA", asalSekolah: "SMA 14 Semarang", tahunLulus: 2018, jurusanYangDituju: "Arsitektur", nilaiUTBK: 725}
	// calonMahasiswa[14] = dataPendaftar{nisn: 1234567904, nama: "Eva Nabila", tempatLahir: "Bandung", tanggalLahir: "2001-03-15", jenisKelamin: "Perempuan", agama: "Islam", email: "eva@email.com", jurusan: "IPS", asalSekolah: "SMA 15 Bandung", tahunLulus: 2019, jurusanYangDituju: "Hukum", nilaiUTBK: 660}
	// calonMahasiswa[15] = dataPendaftar{nisn: 1234567905, nama: "Aditya Putra", tempatLahir: "Yogyakarta", tanggalLahir: "2000-04-16", jenisKelamin: "Laki-laki", agama: "Islam", email: "aditya@email.com", jurusan: "IPA", asalSekolah: "SMA 16 Yogyakarta", tahunLulus: 2018, jurusanYangDituju: "Informatika", nilaiUTBK: 700}
	// calonMahasiswa[16] = dataPendaftar{nisn: 1234567906, nama: "Nurul Azzahra", tempatLahir: "Surakarta", tanggalLahir: "1999-05-17", jenisKelamin: "Perempuan", agama: "Islam", email: "nurul@email.com", jurusan: "IPS", asalSekolah: "SMA 17 Surakarta", tahunLulus: 2019, jurusanYangDituju: "Akuntansi", nilaiUTBK: 680}
	// calonMahasiswa[17] = dataPendaftar{nisn: 1234567907, nama: "Rudi Kurniawan", tempatLahir: "Surabaya", tanggalLahir: "2001-06-18", jenisKelamin: "Laki-laki", agama: "Kristen", email: "rudi@email.com", jurusan: "IPA", asalSekolah: "SMA 18 Surabaya", tahunLulus: 2019, jurusanYangDituju: "Bioteknologi", nilaiUTBK: 710}
	// calonMahasiswa[18] = dataPendaftar{nisn: 1234567908, nama: "Hendri Susilo", tempatLahir: "Jakarta", tanggalLahir: "2000-07-19", jenisKelamin: "Laki-laki", agama: "Islam", email: "hendri@email.com", jurusan: "IPA", asalSekolah: "SMA 19 Jakarta", tahunLulus: 2018, jurusanYangDituju: "Teknik Komputer", nilaiUTBK: 690}
	// calonMahasiswa[19] = dataPendaftar{nisn: 1234567909, nama: "Siti Mulyani", tempatLahir: "Semarang", tanggalLahir: "2001-08-20", jenisKelamin: "Perempuan", agama: "Hindu", email: "siti@email.com", jurusan: "IPS", asalSekolah: "SMA 20 Semarang", tahunLulus: 2019, jurusanYangDituju: "Manajemen", nilaiUTBK: 660}

	//Jika nilai UTBK >= 600 maka status diterima
	statusPendaftaran(&calonMahasiswa) //Hapus jika menggunakan input manual

	fmt.Println("🏫 SELAMAT DATANG 🏫")
	fmt.Println("Penerimaan Mahasiswa Bukit Duri")
awalProgram:
	for {
		if aksiLogin(pengguna) {
			for {
				menuUtama := menuUtama()
				switch menuUtama {
				case 1:
				labelMenuCalonMahasiswa:
					for {
						input := menuCalonMahasiswa()
						switch input {
						case 1:
							fmt.Println("------- Tampilkan semua data -------")
							tampilkanIsi(calonMahasiswa)
						case 2:
							fmt.Println("------- Urutkan Data -------")
						labelUrutkanData:
							for {
								var metode int
								fmt.Println("Pilh metode pengurutan:")
								fmt.Println("1. Selection Sort (Ascending)")
								fmt.Println("2. Selection Sort (Descending)")
								fmt.Println("3. Insertion Sort (Ascending)")
								fmt.Println("4. Insertion Sort (Descending)")
								fmt.Println("5. Kembali")
								fmt.Print("Pilih metode: ")
								fmt.Scan(&metode)

								switch metode {
								case 1:
									selectionSortAscending(&calonMahasiswa)
									fmt.Println("✅ Pengurutan dengan Selection Sort (Ascending) berhasil ✅")
								case 2:
									selectionSortDescending(&calonMahasiswa)
									fmt.Println("✅ Pengurutan dengan Selection Sort (Descending) berhasil ✅")
								case 3:
									insertionSort_Asc(&calonMahasiswa)
									fmt.Println("✅ Pengurutan dengan Insertion Sort (Ascending) berhasil ✅")
								case 4:
									insertionSort_Desc(&calonMahasiswa)
									fmt.Println("✅ Pengurutan dengan Insertion Sort (Descending) berhasil ✅")
								case 5:
									break labelUrutkanData
								}

							}
						case 3:
							fmt.Println("------- 📖 Tambah data 📖 -------")
							var banyakData int
							fmt.Print("Masukkan banyak data yang akan ditambahkan: ")
							fmt.Scan(&banyakData)
							tambahData(&calonMahasiswa, banyakData)

							for i := 0; i < banyakData; i++ {
								fmt.Println(i+1, ". NISN: ", calonMahasiswa[i].nisn)
								fmt.Println("Nama: ", calonMahasiswa[i].nama)
								fmt.Println("Tempat Lahir: ", calonMahasiswa[i].tempatLahir)
								fmt.Println("Tanggal Lahir: ", calonMahasiswa[i].tanggalLahir)
								fmt.Println("Jenis Kelamin: ", calonMahasiswa[i].jenisKelamin)
								fmt.Println("Agama: ", calonMahasiswa[i].agama)
								fmt.Println("Email: ", calonMahasiswa[i].email)
								fmt.Println("Jurusan: ", calonMahasiswa[i].jurusan)
								fmt.Println("Asal Sekolah: ", calonMahasiswa[i].asalSekolah)
								fmt.Println("Tahun Lulus: ", calonMahasiswa[i].tahunLulus)
								fmt.Println("Jurusan yang Dituju: ", calonMahasiswa[i].jurusanYangDituju)
								fmt.Println("Nilai UTBK: ", calonMahasiswa[i].nilaiUTBK)
								fmt.Println("Status: ", calonMahasiswa[i].status)
								fmt.Println()
							}
						case 4:
							//Melakukan pencarian data dengan sequential search dan binary search
							fmt.Println("------- 🔍 Cari Data 🔍 -------")
							var cariNISN int
						labelMenuCari:
							for {
								var metode int
								fmt.Println("Pilih metode pencarian:")
								fmt.Println("1. Sequential Search")
								fmt.Println("2. Binary Search")
								fmt.Println("3. Kembali")
								fmt.Print("Pilih metode pengurutan: ")
								fmt.Scan(&metode)
								switch metode {
								case 1:
									fmt.Println("🔍 Pencarian dengan Sequential Search 🔍")
									fmt.Print("Masukkan NISN yang dicari: ")
									fmt.Scan(&cariNISN)
									hasilSequen := sequentialSearchNISN(&calonMahasiswa, cariNISN)

									if hasilSequen == -1 {
										fmt.Printf("🧐  Tidak ditemukan data untuk calon mahasiswa dengan NISN %d 🧐\n", cariNISN)
										break labelMenuCari
									} else {
										switch menuEditable() {
										case 1:
											// Ubah
											ubahDataCalon(&calonMahasiswa, cariNISN)
										case 2:
											//Hapus
											calonMahasiswa = hapusDataBsByNISN(&calonMahasiswa, cariNISN)
										case 3:
											break labelMenuCari
										}
									}

								case 2:
									fmt.Println("🔍 Pencarian dengan Binary Search 🔍")
									//Urutkan data
									urutkanNISN(&calonMahasiswa)
									//Input NISN
									fmt.Print("Masukkan NISN yang dicari: ")
									fmt.Scan(&cariNISN)
									hasilBinary := binarySearchNISN(&calonMahasiswa, cariNISN)
									if hasilBinary != -1 {
										tampilkanBedasarkanNISN(calonMahasiswa, cariNISN)
										switch menuEditable() {
										case 1:
											//Ubah
											ubahDataCalon(&calonMahasiswa, cariNISN)
										case 2:
											//Hapus
											calonMahasiswa = hapusDataBsByNISN(&calonMahasiswa, cariNISN)
										case 3:
											break labelMenuCalonMahasiswa
										default:
											break
										}
									} else {
										fmt.Printf("🧐  Tidak ditemukan data untuk calon mahasiswa dengan NISN %d 🧐\n", cariNISN)
										break labelMenuCari
									}

								default:
									break labelMenuCari
								}
								if metode == 3 {
									break
								}
							}
						case 5:
							break labelMenuCalonMahasiswa
						}
					}
				case 2:
				labelMenuMahasiswa:
					for {
						pilMenuMahasiswa := menuMahasiswa()
						switch pilMenuMahasiswa {
						case 1:
							fmt.Println("🥳 Menampilkan data siswa yang diterima 🥳")
							mahasiswaDiterima(mahasigma)
						case 2:
							fmt.Println("😊 Menampilkan data siswa yang ditolak 😊")
							mahasiswaDitolak(&calonMahasiswa)
						case 3:
							break labelMenuMahasiswa
						}
					}
				case 3:
					fmt.Println("🪷  Terima kasih telah menggunakan sistem ini 🪷")
					break awalProgram
				}
			}
		} else {
			fmt.Println("❌ Login Gagal ❌")
			fmt.Println("Silahkan coba lagi")
		}

	}
}

// Fungsi menampilkan mahasiswa yang diterima
func mahasiswaDiterima(mahasigma mahasiswa) {
	fmt.Println("Mahasiswa yang diterima: ")
	for i := 0; i < len(mahasigma); i++ {
		fmt.Printf("%d. Nama: %s \n", i+1, mahasigma[i].nama)
		fmt.Println("NIM: ", mahasigma[i].nim)
		fmt.Println("Jurusan: ", mahasigma[i].jurusan)
		fmt.Println()
	}
}

// Fungsi menampilkan mahasiswa yang ditolak
func mahasiswaDitolak(calonMahasiswa *dataSiswa) {
	fmt.Println("Mahasiswa yang ditolak: ")
	for i := 0; i < len(calonMahasiswa); i++ {
		if calonMahasiswa[i].status == "Ditolak" {
			fmt.Println(i+1, ". NISN: ", calonMahasiswa[i].nisn)
			fmt.Println("Nama: ", calonMahasiswa[i].nama)
			fmt.Println("Tempat Lahir: ", calonMahasiswa[i].tempatLahir)
			fmt.Println("Tanggal Lahir: ", calonMahasiswa[i].tanggalLahir)
			fmt.Println("Jenis Kelamin: ", calonMahasiswa[i].jenisKelamin)
			fmt.Println("Agama: ", calonMahasiswa[i].agama)
			fmt.Println("Email: ", calonMahasiswa[i].email)
			fmt.Println("Jurusan: ", calonMahasiswa[i].jurusan)
			fmt.Println("Asal Sekolah: ", calonMahasiswa[i].asalSekolah)
			fmt.Println("Tahun Lulus: ", calonMahasiswa[i].tahunLulus)
			fmt.Println("Jurusan yang Dituju: ", calonMahasiswa[i].jurusanYangDituju)
			fmt.Println("Nilai UTBK: ", calonMahasiswa[i].nilaiUTBK)
			fmt.Println()
		}
	}
}

// Fungsi mengurutkan ASC bedasarkan NISN
func urutkanNISN(calonMahasiswa *dataSiswa) {
	for i := 1; i < len(calonMahasiswa); i++ {
		temp := calonMahasiswa[i]
		// j penunjuk kedua (berada di posisi belakang i)
		j := i
		for j > 0 && float64(temp.nisn) < float64(calonMahasiswa[j-1].nisn) {
			calonMahasiswa[j] = calonMahasiswa[j-1]
			j--
		}
		calonMahasiswa[j] = temp
	}
}

// Fungsi menampilkan semua isi array
func tampilkanIsi(calonMahasiswa dataSiswa) {
	for i, p := range calonMahasiswa {
		fmt.Println(i+1, ". NISN: ", p.nisn)
		fmt.Println("Nama: ", p.nama)
		fmt.Println("Tempat Lahir: ", p.tempatLahir)
		fmt.Println("Tanggal Lahir: ", p.tanggalLahir)
		fmt.Println("Jenis Kelamin: ", p.jenisKelamin)
		fmt.Println("Agama: ", p.agama)
		fmt.Println("Email: ", p.email)
		fmt.Println("Jurusan: ", p.jurusan)
		fmt.Println("Asal Sekolah: ", p.asalSekolah)
		fmt.Println("Tahun Lulus: ", p.tahunLulus)
		fmt.Println("Jurusan yang Dituju: ", p.jurusanYangDituju)
		fmt.Println("Nilai UTBK: ", p.nilaiUTBK)
		fmt.Println("Status: ", p.status)
		fmt.Println()
	}
}

// Fungsi untuk mengubah data mahasiswa
func ubahDataCalon(calonMahasiswa *dataSiswa, cariNISN int) {
	var gantiData int
	fmt.Println("-------------------------------------")
	fmt.Println("Pilih data apa yang ingin di ubah")
	fmt.Println("1. Nama")
	fmt.Println("2. Tempat Lahir")
	fmt.Println("3. Tanggal Lahir")
	fmt.Println("4. Jenis Kelamin")
	fmt.Println("5. Agama")
	fmt.Println("6. Email")
	fmt.Println("7. Jurusan")
	fmt.Println("8. Asal Sekolah")
	fmt.Println("9. Tahun Lulus")
	fmt.Println("10. Jurusan yang Dituju")
	fmt.Print("Masukkan data yang ingin diganti: ")
	fmt.Scan(&gantiData)

	for i := 0; i < len(calonMahasiswa); i++ {
		if calonMahasiswa[i].nisn == cariNISN {
			switch gantiData {
			case 1:
				calonMahasiswa[i].nama = inputString("Masukkan nama baru: ")
				fmt.Println("🎉 Nama berhasil diganti 🎉")
				fmt.Println("---------------------------------")
			case 2:
				fmt.Print("Masukkan tempat lahir baru: ")
				fmt.Scan(&calonMahasiswa[i].tempatLahir)
				fmt.Println("🎉 Tempat lahir berhasil diganti 🎉")
				fmt.Println("-----------------------------------------")
			case 3:
				fmt.Print("Masukkan tanggal lahir baru: ")
				fmt.Scan(&calonMahasiswa[i].tanggalLahir)
				fmt.Println("🎉 Tanggal lahir berhasil diganti 🎉")
				fmt.Println("------------------------------------------")
			case 4:
				fmt.Print("Masukkan jenis kelamin baru: ")
				fmt.Scan(&calonMahasiswa[i].jenisKelamin)
				fmt.Println("🎉 Jenis Kelamin berhasil diganti 🎉")
				fmt.Println("-------------------------------------------")
			case 5:
				calonMahasiswa[i].agama = inputString("Masukkan agama baru: ")
				fmt.Println("🎉 Agama berhasil diganti 🎉")
				fmt.Println("----------------------------------")
			case 6:
				fmt.Print("Masukkan email baru: ")
				fmt.Scan(&calonMahasiswa[i].email)
				fmt.Println("🎉 Email berhasil diganti 🎉")
				fmt.Println("----------------------------------")
			case 7:
				fmt.Print("Masukkan jurusan baru: ")
				fmt.Scan(&calonMahasiswa[i].jurusan)
				fmt.Println("🎉 Jurusan berhasil diganti 🎉")
				fmt.Println("-----------------------------------")
			case 8:
				calonMahasiswa[i].asalSekolah = inputString("Masukkan asal sekolah baru: ")
				fmt.Println("🎉 Asal Sekolah berhasil diganti 🎉")
				fmt.Println("-----------------------------------------")
			case 9:
				fmt.Print("Masukkan tahun lulus baru: ")
				fmt.Scan(&calonMahasiswa[i].tahunLulus)
				fmt.Println("🎉 Tahun Lulus berhasil diganti 🎉")
				fmt.Println("----------------------------------------")
			case 10:
				calonMahasiswa[i].jurusanYangDituju = inputString("Masukkan jurusan yang dituju baru: ")
				fmt.Println("🎉 Jurusan yang dituju berhasil diganti 🎉")
				fmt.Println("------------------------------------------------")
			}
		}
	}
}

func inputString(judul string) string {
	fmt.Print(judul)
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Fungsi untuk pencarian (sequential search)
func sequentialSearchNISN(calonMahasiswa *dataSiswa, cariNISN int) int {
	hasil := -1
	for i := 0; i < len(calonMahasiswa); i++ {
		if calonMahasiswa[i].nisn == cariNISN {
			fmt.Printf("🕵️  Ditemukan Data Calon Mahasiswa dengan NISN %d 🕵️\n", cariNISN)
			fmt.Println("NISN: ", calonMahasiswa[i].nisn)
			fmt.Println("Nama: ", calonMahasiswa[i].nama)
			fmt.Println("Tempat Lahir: ", calonMahasiswa[i].tempatLahir)
			fmt.Println("Tanggal Lahir: ", calonMahasiswa[i].tanggalLahir)
			fmt.Println("Jenis Kelamin: ", calonMahasiswa[i].jenisKelamin)
			fmt.Println("Agama: ", calonMahasiswa[i].agama)
			fmt.Println("Email: ", calonMahasiswa[i].email)
			fmt.Println("Jurusan: ", calonMahasiswa[i].jurusan)
			fmt.Println("Asal Sekolah: ", calonMahasiswa[i].asalSekolah)
			fmt.Println("Tahun Lulus: ", calonMahasiswa[i].tahunLulus)
			fmt.Println("Jurusan yang Dituju: ", calonMahasiswa[i].jurusanYangDituju)
			fmt.Println("Nilai UTBK: ", calonMahasiswa[i].nilaiUTBK)
			fmt.Println("Status: ", calonMahasiswa[i].status)
			hasil = calonMahasiswa[i].nisn
		}

	}
	return hasil
}

// Fungsi menampilkan data calon mahasiswa bedasarkan NISN
func tampilkanBedasarkanNISN(calonMahasiswa dataSiswa, nisn int) {
	for i := 0; i < len(calonMahasiswa); i++ {
		if calonMahasiswa[i].nisn == nisn {
			fmt.Printf("🕵️  Ditemukan Data Calon Mahasiswa dengan NISN %d 🕵️\n", nisn)
			fmt.Println("NISN: ", calonMahasiswa[i].nisn)
			fmt.Println("Nama: ", calonMahasiswa[i].nama)
			fmt.Println("Tempat Lahir: ", calonMahasiswa[i].tempatLahir)
			fmt.Println("Tanggal Lahir: ", calonMahasiswa[i].tanggalLahir)
			fmt.Println("Jenis Kelamin: ", calonMahasiswa[i].jenisKelamin)
			fmt.Println("Agama: ", calonMahasiswa[i].agama)
			fmt.Println("Email: ", calonMahasiswa[i].email)
			fmt.Println("Jurusan: ", calonMahasiswa[i].jurusan)
			fmt.Println("Asal Sekolah: ", calonMahasiswa[i].asalSekolah)
			fmt.Println("Tahun Lulus: ", calonMahasiswa[i].tahunLulus)
			fmt.Println("Jurusan yang Dituju: ", calonMahasiswa[i].jurusanYangDituju)
			fmt.Println("Nilai UTBK: ", calonMahasiswa[i].nilaiUTBK)
			fmt.Println("Status: ", calonMahasiswa[i].status)
		}

	}

}

// Fungsi untuk penarian (binary search)
func binarySearchNISN(calonMahasiswa *dataSiswa, cariNISN int) int {
	low := 0
	high := len(calonMahasiswa) - 1

	for low <= high {
		mid := (low + high) / 2
		if calonMahasiswa[mid].nisn == cariNISN {
			return mid
		} else if calonMahasiswa[mid].nisn < cariNISN {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Fungsi untuk menu ubah dan hapus
func hapusDataBsByNISN(calonMahasiswa *dataSiswa, cariNisn int) dataSiswa {
	idx := binarySearchNISN(calonMahasiswa, cariNisn)
	if idx == -1 {
		fmt.Println("404 data not found")
		return *calonMahasiswa
	} else {
		fmt.Println("✔️ 🗑️ Data mahasiswa berhasil dihapus")
	}

	var newData [max]dataPendaftar
	j := 0
	for i := 0; i < len(calonMahasiswa); i++ {
		if i != idx {
			newData[j] = calonMahasiswa[i]
			j++
		}
	}

	// Bersihkan sisanya
	for k := j; k < max; k++ {
		newData[k] = dataPendaftar{}
	}

	return newData
}

func menuEditable() int {
	var aksi int
	fmt.Println("-------------")
	fmt.Println("1. Ubah Data")
	fmt.Println("2. Hapus Data")
	fmt.Println("3. Kembali")
	fmt.Print("Aksi: ")
	fmt.Scan(&aksi)
	return aksi
}

// Fungsi Insertion Sort ASC
func insertionSort_Asc(calonMahasiswa *dataSiswa) {
	// i penunjuk pertama
	for i := 1; i < len(calonMahasiswa); i++ {
		temp := calonMahasiswa[i]
		// j penunjuk kedua (berada di posisi belakang i)
		j := i
		for j > 0 && temp.nilaiUTBK < calonMahasiswa[j-1].nilaiUTBK {
			calonMahasiswa[j] = calonMahasiswa[j-1]
			j--
		}
		calonMahasiswa[j] = temp
	}
}

// Fungsi Insertion Sort DESC
func insertionSort_Desc(calonMahasiswa *dataSiswa) {
	// i penunjuk pertama
	for i := 1; i < len(calonMahasiswa); i++ {
		temp := calonMahasiswa[i]
		// j penunjuk kedua (berada di posisi belakang i)
		j := i
		for j > 0 && temp.nilaiUTBK > calonMahasiswa[j-1].nilaiUTBK {
			calonMahasiswa[j] = calonMahasiswa[j-1]
			j--
		}
		calonMahasiswa[j] = temp
	}
}

// Fungsi untuk mengurutkan data menggunakan selection sort ascending
func selectionSortAscending(calonMahasiswa *dataSiswa) {
	for i := 0; i < len(calonMahasiswa)-1; i++ {
		idx_min := i
		for j := i + 1; j < len(calonMahasiswa); j++ {
			if calonMahasiswa[j].nilaiUTBK < calonMahasiswa[idx_min].nilaiUTBK {
				idx_min = j
			}
		}
		calonMahasiswa[i], calonMahasiswa[idx_min] = calonMahasiswa[idx_min], calonMahasiswa[i]
	}
}

// Fungsi untuk mengurutkan data menggunakan selection sort ascending
func selectionSortDescending(calonMahasiswa *dataSiswa) {
	for i := 0; i < len(calonMahasiswa)-1; i++ {
		idx_min := i
		for j := i + 1; j < len(calonMahasiswa); j++ {
			if calonMahasiswa[j].nilaiUTBK > calonMahasiswa[idx_min].nilaiUTBK {
				idx_min = j
			}
		}
		calonMahasiswa[i], calonMahasiswa[idx_min] = calonMahasiswa[idx_min], calonMahasiswa[i]
	}
}

// Fungsi untuk menentukan status penerimaan
func statusPendaftaran(calonMahasiswa *dataSiswa) {
	indeks := 0
	for i := 0; i < len(calonMahasiswa); i++ {
		if calonMahasiswa[i].nilaiUTBK >= 600 {
			calonMahasiswa[i].status = "Diterima"
			mahasigma[indeks].nama = calonMahasiswa[i].nama
			mahasigma[indeks].nim = i + 1
			mahasigma[indeks].jurusan = calonMahasiswa[i].jurusanYangDituju
			indeks++
		} else {
			calonMahasiswa[i].status = "Ditolak"

		}
	}
	// fmt.Println("Banyak :", banyakYangDiterima)
	// for j := 0; j <= banyakYangDiterima; j++ {
	// 	mahasigma[j].nama = calonMahasiswa[j].nama
	// 	mahasigma[j].nim = j + 1
	// 	mahasigma[j].jurusan = calonMahasiswa[j].jurusanYangDituju
	// }
}

// Fungsi untuk menambah data
func tambahData(calonMahasiswa *dataSiswa, banyakData int) {

	for i := 0; i < banyakData; i++ {
		fmt.Print("Masukkan NISN: ")
		fmt.Scan(&calonMahasiswa[i].nisn)
		calonMahasiswa[i].nama = inputString("Masukkan Nama: ")
		fmt.Print("Masukkan Tempat Lahir: ")
		fmt.Scan(&calonMahasiswa[i].tempatLahir)
		fmt.Print("Masukkan Tanggal Lahir: ")
		fmt.Scan(&calonMahasiswa[i].tanggalLahir)
		fmt.Print("Masukkan Jenis Kelamin: ")
		fmt.Scan(&calonMahasiswa[i].jenisKelamin)
		calonMahasiswa[i].agama = inputString("Masukkan Agama: ")
		fmt.Print("Masukkan Email: ")
		fmt.Scan(&calonMahasiswa[i].email)
		fmt.Print("Masukkan Jurusan: ")
		fmt.Scan(&calonMahasiswa[i].jurusan)
		calonMahasiswa[i].asalSekolah = inputString("Masukkan Asal Sekolah: ")
		fmt.Print("Masukkan Tahun Lulus: ")
		fmt.Scan(&calonMahasiswa[i].tahunLulus)
		calonMahasiswa[i].jurusanYangDituju = inputString("Masukkan Jurusan yang Dituju: ")
		fmt.Print("Masukkan Nilai UTBK: ")
		fmt.Scan(&calonMahasiswa[i].nilaiUTBK)
		statusPendaftaran(calonMahasiswa)
		fmt.Println()
	}
	fmt.Println("✅ Data berhasil ditambahkan")
}

// Prosedur untuk menampilkan menu mahasiswa
func menuMahasiswa() int {
	var pilihan int
	fmt.Println("-------- MENU MAHASISWA --------")
	fmt.Println("1. Tampilkan Siswa yang Diterima")
	fmt.Println("2. Tampilkan Siswa yang Ditolak")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)
	return pilihan
}

// Prosedur untuk menampilkan menu calon mahasiswa
func menuCalonMahasiswa() int {
	var pilihan int
	fmt.Println("-------- MENU CALON MAHASISWA --------")
	fmt.Println("1. Tampilkan Semua Data")
	fmt.Println("2. Urutkan Data")
	fmt.Println("3. Tambah Data")
	fmt.Println("4. Cari Data")
	fmt.Println("5. Keluar")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)
	return pilihan

}

// Fungsi untuk menampilkan menu utama
func menuUtama() int {
	var pilihan int
	fmt.Println("-------- MENU UTAMA --------")
	fmt.Println("1. Calon Mahasiswa")
	fmt.Println("2. Mahasiswa")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)
	return pilihan
}

// Fungsi untuk verifikasi login
func aksiLogin(pengguna [4]login) bool {
	var nama, password string
	fmt.Print("Masukkan nama pengguna: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan password: ")
	fmt.Scanln(&password)

	ditemukan := false
	for _, pengguna := range pengguna {
		if pengguna.nama == nama && pengguna.password == password {
			ditemukan = true
		}
	}
	return ditemukan
}
