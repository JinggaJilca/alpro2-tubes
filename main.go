package main

import "fmt"

/* SISTEM PENERIMAAN MAHASISWA BUKIT DURI
Sistem bedasarkan seleksi nilai UTBK apabila nilai lebih dari 600 maka diterima

======== SPESIFIKASI UMUM
✅ Program harus dibuat secara modular dengan menggunakan subprogram.
✅ Struktur subprogram ditentukan oleh masing-masing kelompok. Setiap subprogram harus dilengkapi dengan parameter dan spesifikasinya.
✅ Program harus mengimplementasikan Array dan Tipe Bentukan. Array statis harus digunakan, bukan array dinamis atau slice.
✅ Program harus mengimplementasikan algoritma Sequential dan Binary Search untuk pencarian data, pengubahan (edit), atau penghapusan data tertentu.
✅ Program harus mengimplementasikan algoritma Selection Sort dan Insertion Sort untuk mengurutkan data dalam kategori yang berbeda saat menampilkan hasil. Setiap kategori harus bisa diurutkan dalam urutan naik (ascending) maupun turun (descending).
✅ Penggunaan statement break (kecuali dalam repeat-until) dan continue tidak diperbolehkan.
✅ Variabel global hanya diperbolehkan untuk menyimpan array utama yang akan diproses.
*/
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
type dataMahasigma struct {
	nim     int
	nama    string
	jurusan string
}

func main() {
	var calonMahasiswa [20]dataPendaftar
	// var mahasigma [10]dataMahasigma

	//Data dummy calon mahasiswa
	calonMahasiswa[0] = dataPendaftar{nisn: 1234567890, nama: "Andi Pratama", tempatLahir: "Surabaya", tanggalLahir: "2000-01-01", jenisKelamin: "Laki-laki", agama: "Islam", email: "andi@email.com", jurusan: "IPA", asalSekolah: "SMA 1 Surabaya", tahunLulus: 2018, jurusanYangDituju: "Teknik Informatika", nilaiUTBK: 700}
	calonMahasiswa[1] = dataPendaftar{nisn: 1234567891, nama: "Siti Nurjanah", tempatLahir: "Malang", tanggalLahir: "2001-02-02", jenisKelamin: "Perempuan", agama: "Islam", email: "siti@email.com", jurusan: "IPS", asalSekolah: "SMA 2 Malang", tahunLulus: 2019, jurusanYangDituju: "Manajemen", nilaiUTBK: 650}
	calonMahasiswa[2] = dataPendaftar{nisn: 1234567892, nama: "Budi Santoso", tempatLahir: "Surabaya", tanggalLahir: "2000-03-03", jenisKelamin: "Laki-laki", agama: "Kristen", email: "budi@email.com", jurusan: "IPA", asalSekolah: "SMA 3 Surabaya", tahunLulus: 2019, jurusanYangDituju: "Teknik Mesin", nilaiUTBK: 710}
	calonMahasiswa[3] = dataPendaftar{nisn: 1234567893, nama: "Dewi Lestari", tempatLahir: "Bandung", tanggalLahir: "1999-04-04", jenisKelamin: "Perempuan", agama: "Hindu", email: "dewi@email.com", jurusan: "IPA", asalSekolah: "SMA 4 Bandung", tahunLulus: 2018, jurusanYangDituju: "Teknik Elektro", nilaiUTBK: 720}
	calonMahasiswa[4] = dataPendaftar{nisn: 1234567894, nama: "Fajar Ramadhan", tempatLahir: "Yogyakarta", tanggalLahir: "2000-05-05", jenisKelamin: "Laki-laki", agama: "Islam", email: "fajar@email.com", jurusan: "IPS", asalSekolah: "SMA 5 Yogyakarta", tahunLulus: 2019, jurusanYangDituju: "Ekonomi", nilaiUTBK: 680}
	calonMahasiswa[5] = dataPendaftar{nisn: 1234567895, nama: "Hanafi Rizky", tempatLahir: "Jakarta", tanggalLahir: "1999-06-06", jenisKelamin: "Laki-laki", agama: "Islam", email: "hanafi@email.com", jurusan: "IPA", asalSekolah: "SMA 6 Jakarta", tahunLulus: 2018, jurusanYangDituju: "Fisika", nilaiUTBK: 725}
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

	//Menampilkan data calon mahasiswa
	/*
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
	*/
	fmt.Println("======= SELAMAT DATANG ADMIN =======")
	fmt.Println("Penerimaan Mahasiswa Bukit Duri")
	switch mainMenu() {
	case 1:
		fmt.Println("Menampilkan semua data")
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

	case 2:
		fmt.Println("Urutkan data bedasarkan kategori")
		//Penguruan dengan selection sort dan insertion sort harus bisa menaik (ascending) atau menurun (descending)
	case 3:
		fmt.Println("Tambah data")
	case 4:
		//Melakukan pencarian data dengan sequential search dan binary search
		fmt.Println("Ubah data")
	case 5:
		fmt.Println("Hapus data")
	case 6:
		fmt.Println("Terima kasih telah menggunakan sistem ini")
	}

}
func mainMenu() int {
	var pilihan int
	fmt.Println("======= MENU UTAMA =======")
	fmt.Println("1. Tampilkan Semua Data")
	fmt.Println("2. Urutkan Data Sesuai kategori")
	fmt.Println("3. Tambah Data")
	fmt.Println("4. Ubah Data")
	fmt.Println("5. Hapus Data")
	fmt.Println("6. Keluar")
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&pilihan)
	return pilihan

}

// ============ *Kalau ada fitur login beb, jangan panik duluw
/*
func verifikasiLogin() {
	var nama, password string
	fmt.Println("======= SELAMAT DATANG =======")
	fmt.Println("Penerimaan Mahasiswa Bukit Duri")
	//Input nama
	fmt.Print("Masukkan nama : ")
	fmt.Scanln(&nama)
	//Input password
	fmt.Print("Masukkan password : ")
	fmt.Scanln(&password)

	switch {
	case nama == "admin" && password == "admin":
		fmt.Println("Hello admin")
		//Masuk ke menu admin
	default:
	}
}
*/
